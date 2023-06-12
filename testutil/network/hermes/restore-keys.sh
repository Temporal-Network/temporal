#!/bin/bash
set -e

### Delete keys if they exist
###hermes --config ./testutil/network/hermes/config.toml keys delete --chain temporal --key-name testkey
###hermes --config ./testutil/network/hermes/config.toml keys delete --chain uni-6 --key-name testkey

### Restore Keys
hermes --config ./testutil/network/hermes/config.toml keys add --chain temporal --mnemonic-file ./testutil/network/hermes/temporal.txt
sleep 5

hermes --config ./testutil/network/hermes/config.toml keys add --chain uni-6 --mnemonic-file ./testutil/network/hermes/uni-6.txt
sleep 5
