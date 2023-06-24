#!/bin/bash

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
SCRIPT=$SCRIPTPATH"/helper_ignite_local_chain_setup.sh"
chmod +x $SCRIPT

#Change the last parameter below to change which version of which local net to launch
$SCRIPT "juno" "https://github.com/CosmosContracts/juno.git" "v15.0.0"
$SCRIPT "temporal" "https://github.com/temporal-zone/temporal.git" "main"

echo "####################################"
echo "Run the following three commands in new console windows:"
echo "####################################"
echo "temporald config node tcp://127.0.0.1:36657 && temporald start --grpc.address 0.0.0.0:10090 --pruning nothing --p2p.laddr 0.0.0.0:36656 --address 0.0.0.0:36658 --rpc.laddr tcp://0.0.0.0:36657 --rpc.pprof_laddr 0.0.0.0:7060 --api.address tcp://0.0.0.0:2317 --grpc-web.address 0.0.0.0:10091"
echo "####################################"
echo "junod start --pruning nothing"
echo "#################################### On first run:"
echo "bash $SCRIPTPATH/hermes/restore-keys.sh && bash $SCRIPTPATH/hermes/create-conn.sh && bash $SCRIPTPATH/hermes/start.sh"
echo "#################################### On subsequent runs after a chain reset:"
echo "bash $SCRIPTPATH/hermes/create-conn.sh && bash $SCRIPTPATH/hermes/start.sh"
echo "#################################### On subsequent runs with out a chain reset:"
echo "bash $SCRIPTPATH/hermes/start.sh"
echo "####################################"
echo "Note if either of the above commands throw that it can't find temporald or junod:"
echo "Run: export PATH=$PATH:$HOME/go/bin/"