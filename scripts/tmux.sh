#!/bin/bash
set -euo pipefail
IFS=$'\n\t'
export TMUX_SESSION=${TMUX_SESSION:="bacalhau-lilypad"}
export APP=${APP:=""}
export PREDICTABLE_API_PORT=1
export BACALHAU_API_HOST=localhost
export BACALHAU_API_PORT=20000
export BACALHAU_DASHBOARD_POSTGRES_HOST=127.0.0.1
export BACALHAU_DASHBOARD_POSTGRES_DATABASE=postgres
export BACALHAU_DASHBOARD_POSTGRES_USER=postgres
export BACALHAU_DASHBOARD_POSTGRES_PASSWORD=postgres
export BACALHAU_DASHBOARD_JWT_SECRET=apples
export LOG_LEVEL=DEBUG

function start() {
  if tmux has-session -t "$TMUX_SESSION" 2>/dev/null; then
    echo "Session $TMUX_SESSION already exists. Attaching..."
    sleep 1
    tmux -2 attach -t $TMUX_SESSION
    exit 0;
  fi

  echo "Creating tmux session $TMUX_SESSION..."

  source ./hardhat/.env

  if [[ -z "$WALLET_PRIVATE_KEY_DEV" ]]; then
    ## print error to stderr and exit
    echo "WALLET_PRIVATE_KEY_DEV not set" >&2
    exit 1
  fi

  export WALLET_PRIVATE_KEY=$WALLET_PRIVATE_KEY_DEV
  export CONTRACT_ADDRESS=$CONTRACT_ADDRESS_DEV

  # get the size of the window and create a session at that size
  local screensize=$(stty size)
  local width=$(echo -n "$screensize" | awk '{print $2}')
  local height=$(echo -n "$screensize" | awk '{print $1}')
  tmux -2 new-session -d -s $TMUX_SESSION -x "$width" -y "$(($height - 1))"

  # the right hand col with a 50% vertical split
  tmux split-window -h -d
  tmux select-pane -t 1
  tmux split-window -v -d
  tmux select-pane -t 0
  tmux split-window -v -d

  tmux send-keys -t 0 'cd hardhat && sleep 5' C-m
  tmux send-keys -t 0 'npx hardhat run --network localhost scripts/deployUpgradeable.ts' C-m
  tmux send-keys -t 1 'cd hardhat && npm run node' C-m
  
  tmux -2 attach-session -t $TMUX_SESSION
}

function stop() {
  echo "Stopping tmux session $TMUX_SESSION..."
  tmux kill-session -t $TMUX_SESSION
}

function help() {
  echo "bash scripts/tmux.sh start|stop"
}

command="$@"

if [ -z "$command" ]; then
  command="help"
fi

eval "$command"