#!/bin/bash


if [[ -L /etc/init.d/opsAgent ]]; then
    rm -f /etc/init.d/opsAgent
fi

if [[ -L /etc/systemd/system/opsAgent.service ]]; then
    rm -f /etc/systemd/system/opsAgent.service
fi

if [[ ! -f /etc/opsAgent/opsAgent.conf ]]; then
    cp -fv /opt/opsAgent/conf/opsAgent.conf.sample /etc/opsAgent/opsAgent.conf.sample
fi

if [[ ! -f /etc/opsAgent/opsAgent.conf ]] && [[ -f /etc/opsAgent/opsAgent.conf.sample ]]; then
   cp /etc/opsAgent/opsAgent.conf.sample /etc/opsAgent/opsAgent.conf
fi

cp -af /opt/opsAgent/opsAgent /usr/bin/opsAgent

LOG_DIR=/var/log/opsAgent
test -d $LOG_DIR || mkdir -p $LOG_DIR
chown -R -L opsAgent:opsAgent $LOG_DIR
chmod 755 $LOG_DIR

chown -R -L opsAgent:opsAgent /opt/opsAgent

if [[ -d /run/systemd/system ]]; then
    cp -f /opt/opsAgent/lib/scripts/opsAgent.service /usr/lib/systemd/system/opsAgent.service
    systemctl enable opsAgent
    systemctl daemon-reload
fi