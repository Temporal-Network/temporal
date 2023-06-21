#!/bin/bash

VAL_1_MNEMONIC="ocean toe void inquiry skin exotic artist teach daughter theory melt theme desk more trip person maid agent mesh prefer loop oyster member exact"
VAL_1_ADDRESS="temporal1nhvm23r9pu593sgdu660whdz66pg382czn5quy"

CMD="temporald"
DENOM="utprl"

$CMD start --grpc.address 0.0.0.0:10090 --pruning nothing --p2p.laddr 0.0.0.0:36656 --address 0.0.0.0:36658 --rpc.laddr tcp://0.0.0.0:36657 --rpc.pprof_laddr 0.0.0.0:7060 --api.address tcp://0.0.0.0:2317 --grpc-web.address 0.0.0.0:10091