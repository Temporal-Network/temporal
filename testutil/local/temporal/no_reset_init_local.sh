#!/bin/bash

CMD="temporald"
DENOM="utprl"

$CMD start --grpc.address 0.0.0.0:10090 --pruning nothing --p2p.laddr 0.0.0.0:36656 --address 0.0.0.0:36658 --rpc.laddr tcp://0.0.0.0:36657 --rpc.pprof_laddr 0.0.0.0:7060 --api.address tcp://0.0.0.0:2317 --api.enable true --grpc-web.address 0.0.0.0:10091