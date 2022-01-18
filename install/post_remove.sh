#!/bin/sh

# Generated by: tyk-ci/wf-gen
# Generated on: Fri Jan 14 17:45:41 UTC 2022

# Generation commands:
# ./pr.zsh -base v1.2.2-rc5 -branch v1.2.2-rc5-m4-sync -title sync m4 templates -repos tyk-identity-broker
# m4 -E -DxREPO=tyk-identity-broker


cleanRemove() {
    printf "\033[32m Post remove for plain removal\033[0m\n"
    if command -V systemctl >/dev/null 2>&1; then
        systemctl stop tyk-identity-broker ||:
        systemctl daemon-reload ||:
    fi
    service stop tyk-identity-broker ||:
    if command -V chkconfig >/dev/null 2>&1; then
        chkconfig --del tyk-identity-broker ||:
    fi
    if command -V update-rc.d >/dev/null 2>&1; then
        update-rc.d tyk-identity-broker remove
    fi
}

upgrade() {
    printf "\033[32m Post remove for upgrade, nothing to do.\033[0m\n"
}

action="$1"
if  [ "$1" = "configure" ] && [ -z "$2" ]; then
    # Alpine linux does not pass args, and deb passes $1=configure
    action="install"
elif [ "$1" = "configure" ] && [ -n "$2" ]; then
    # deb passes $1=configure $2=<current version>
    action="upgrade"
fi

case "$action" in
    "1" | "install")
        printf "\033[32m Post Install of a clean install\033[0m\n"
        cleanRemove
        ;;
    "2" | "upgrade")
        printf "\033[32m Post Install of an upgrade\033[0m\n"
        upgrade
        ;;
    *)
        # $1 == version being installed
        printf "\033[32m Alpine\033[0m"
        cleanRemove
        ;;
esac
