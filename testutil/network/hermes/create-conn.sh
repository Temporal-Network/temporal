#!/bin/bash
set -e

### Configure the clients and connection
echo "Initiating connection handshake..."
hermes --config ./testutil/network/hermes/config.toml create connection --a-chain temporal --b-chain juno

sleep 2
