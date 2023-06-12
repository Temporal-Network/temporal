#!/bin/bash
set -e

### Configure the clients and connection
echo "Initiating connection handshake..."
hermes --config ./testutil/network/hermes/config.toml create connection --a-chain temporal --b-chain uni-6

sleep 2
