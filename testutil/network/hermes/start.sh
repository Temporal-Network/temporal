#!/bin/bash

# Start the hermes relayer in multi-paths mode
echo "Starting hermes relayer..."
hermes --config ./testutil/network/hermes/config.toml start
