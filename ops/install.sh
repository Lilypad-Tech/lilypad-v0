#!/bin/bash
set -euo pipefail
IFS=$'\n\t'
set -x

NETWORK="$1"

cp /tmp/$NETWORK.env /root/lilypad-$NETWORK.env
chmod 0400 /root/lilypad-$NETWORK.env
echo "SQLITE_FILE_LOCATION=/lilypad-$NETWORK.sqlite" >> /root/lilypad-$NETWORK.env

cp /tmp/lilypad.service /etc/systemd/system/lilypad-$NETWORK.service
chmod 0444 /etc/systemd/system/lilypad-$NETWORK.service
echo "EnvironmentFile=/root/lilypad-$NETWORK.env" >> /etc/systemd/system/lilypad-$NETWORK.service

systemctl daemon-reload
systemctl enable lilypad-$NETWORK
systemctl restart lilypad-$NETWORK
