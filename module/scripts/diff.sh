#!/bin/bash
set -e
OS=$(uname | tr '[:upper:]' '[:lower:]')
CMD_PATH=$(echo "https://github.com/FinalCAD/terraform-diff-module/releases/download/$1/diff_$1_${OS}_amd64.tar.gz")
RM=$(rm -f /tmp/diff)
DIFF=$(wget -qO- $CMD_PATH | tar xvzf - -C /tmp &>/dev/null)
/tmp/diff -i "$2" -u "$3" -json
