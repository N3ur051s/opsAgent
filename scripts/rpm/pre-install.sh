#!/bin/bash

if ! grep "^opsAgent" /etc/group &>/dev/null; then
    groupadd -r opsAgent
fi

if ! id opsAgent &>/dev/null; then
    useradd -r -M opsAgent -s /bin/false -d /home/opsAgent -g opsAgent
    mkdir -p /home/opsAgent/.opsAgent/
    chown -R -L opsAgent:opsAgent /home/opsAgent/.opsAgent/
fi

if id opsAgent &>/dev/null; then
    HOMEDIR=`cat /etc/passwd | grep 'opsAgent' | cut -d ':' -f 6`
    mkdir -p $HOMEDIR/.opsAgent/
    chown -R -L opsAgent:opsAgent $HOMEDIR/.opsAgent/
fi

if [[ ! -d /etc/opsAgent ]]; then
    echo -e "Please note, opsAgent's configuration is now located at '/etc/opsAgent'."
    mkdir -p /etc/opsAgent
    chown -R -L opsAgent:opsAgent /etc/opsAgent

    if [[ -f /etc/opsAgent/opsAgent.conf ]]; then
        backup_name="opsAgent.conf.$(date +%s).backup"
        echo "A backup of your current configuration can be found at: /etc/opsAgent/${backup_name}"
        cp -a "/etc/opsAgent/opsAgent.conf" "/etc/opsAgent/${backup_name}"
    fi
fi
