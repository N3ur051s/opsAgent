#!/bin/bash

HOMEDIR=`cat /etc/passwd | grep 'opsAgent' | cut -d ':' -f 6`

rm -rf /etc/opsAgent
rm -rf $HOMEDIR/.opsAgent
rm -rf /var/log/opsAgent
rm -rf /opt/opsAgent
rm -f /usr/bin/opsAgent

if [[ -d /run/systemd/system ]]; then
    systemctl disable opsAgent
    rm -f /usr/lib/systemd/system/opsAgent.service
    systemctl daemon-reload
fi