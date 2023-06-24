#!/bin/bash
set -e

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

### Delete keys if they exist
###hermes --config ./testutil/network/hermes/config.toml keys delete --chain temporal --key-name testkey
###hermes --config ./testutil/network/hermes/config.toml keys delete --chain uni-6 --key-name testkey

### Restore Keys
hermes --config $SCRIPTPATH/config.toml keys add --chain temporal --mnemonic-file $SCRIPTPATH/temporal.txt
sleep 5

hermes --config $SCRIPTPATH/config.toml keys add --chain juno --mnemonic-file $SCRIPTPATH/localjuno.txt
sleep 5
