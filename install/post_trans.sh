#!/bin/sh

# Generated by: tyk-ci/wf-gen
# Generated on: Fri Jan 14 17:45:41 UTC 2022

# Generation commands:
# ./pr.zsh -base v1.2.2-rc5 -branch v1.2.2-rc5-m4-sync -title sync m4 templates -repos tyk-identity-broker
# m4 -E -DxREPO=tyk-identity-broker


if command -V systemctl >/dev/null 2>&1; then
    if [ ! -f /lib/systemd/system/tyk-identity-broker.service ]; then
        cp /opt/tyk-identity-broker/install/inits/systemd/system/tyk-identity-broker.service /lib/systemd/system/tyk-identity-broker.service
    fi
else
    if [ ! -f /etc/init.d/tyk-identity-broker ]; then
        cp /opt/tyk-identity-broker/install/inits/sysv/init.d/tyk-identity-broker /etc/init.d/tyk-identity-broker
    fi
fi
