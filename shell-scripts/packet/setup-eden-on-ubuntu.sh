#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
EDEN_DIR=$SCRIPT_DIR/../../
EDEN_COMMIT_SHA=$1

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get install -y git make jq qemu binfmt-support qemu-user-static qemu-utils qemu-system-x86 qemu-system-aarch64
git clone https://github.com/lf-edge/eden.git -b "$EDEN_COMMIT_SHA"
