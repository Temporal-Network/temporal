#!/bin/bash

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# Start the hermes relayer in multi-paths mode
echo "Starting hermes relayer..."
hermes --config $SCRIPTPATH/config.toml start
