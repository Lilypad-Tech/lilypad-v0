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

gSCP ../hardhat/.env /tmp/env
gSSH mv /tmp/env /root/lilypad.env
gSSH chmod 0400 /root/lilypad.env

gSCP lilypad.service /tmp/lilypad.service
gSSH mv /tmp/lilypad.service /etc/systemd/system/lilypad.service
gSSH chmod 0444 /etc/systemd/system/lilypad.service

gSCP ./../bin/lilypad-linux-amd64 /tmp/lilypad
gSSH mv /tmp/lilypad /usr/bin/lilypad
gSSH chmod 0555 /usr/bin/lilypad

gSSH systemctl daemon-reload
gSSH systemctl enable lilypad
gSSH systemctl restart lilypad
