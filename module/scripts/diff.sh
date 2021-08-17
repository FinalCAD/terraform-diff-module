#!/bin/bash
set -e
CMD_PATH=$(echo "https://github.com/FinalCAD/terraform-diff-module/releases/download/$1/diff_$1_linux_amd64.tar.gz")
DIFF=$(wget -qO- $CMD_PATH | tar xvzf - -C /tmp &>/dev/null)
/tmp/diff -i "$2" -u "$3" -json
