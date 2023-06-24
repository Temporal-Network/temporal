#!/bin/bash
set -e

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

### Configure the clients and connection
echo "Initiating connection handshake..."
hermes --config $SCRIPTPATH/config.toml create connection --a-chain temporal --b-chain juno

sleep 2
