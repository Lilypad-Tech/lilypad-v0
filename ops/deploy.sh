#!/bin/bash
set -euo pipefail
IFS=$'\n\t'
set -x

ZONE="us-central1-a"
NAME="lilypad-vm-0"

function gSSH() {
    gcloud compute ssh --quiet --zone=$ZONE $NAME -- sudo $*
}

function gSCP() {
    gcloud compute scp --quiet --zone=$ZONE $1 $NAME:$2
}

gSCP lilypad.service /tmp/lilypad.service
gSCP ./../bin/lilypad-linux-amd64 /tmp/lilypad
gSCP ./install.sh /tmp/install.sh

gSSH mv /tmp/lilypad /usr/bin/lilypad
gSSH chmod 0555 /usr/bin/lilypad

for NETWORK in "$@"; do
    gSCP ../hardhat/$NETWORK.env /tmp/$NETWORK.env
    gSSH bash /tmp/install.sh $NETWORK
done
