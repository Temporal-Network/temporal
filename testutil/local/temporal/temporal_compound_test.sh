#!/bin/bash

CMD="temporald"
DENOM="utprl"
CHAIN_ID="temporal"
CONNECTION_ID="connection-0"
LOCAL_NODE="tcp://127.0.0.1:36657"
FLAGS="--from aliceT --chain-id $CHAIN_ID --node $LOCAL_NODE --yes"
RELAYER_SEED="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"

$CMD tx icayieldmos create-connection $CONNECTION_ID $FLAGS

$CMD tx icayieldmos create-contract-remote-zone $CONNECTION_ID juno true juno $FLAGS

echo "$RELAYER_SEED" | $CMD keys add relayer --recover --keyring-backend=test 2>&1

$CMD tx icayieldmos create-remote-contract-compound-settings 0 juno17dtl0mjt3t77kpuhg2edqzjpszulwhgz2qxy4v '{"compound":{"comp_prefs":{"relative":[{"destination":{"juno_staking":{"validator_address":"junovaloper1cvk3scc8akj4psfzre7xxyt9453ug6x7ul5gw3"}},"amount":"1000000000000000000"}]},"delegator_address":"juno17dtl0mjt3t77kpuhg2edqzjpszulwhgz2qxy4v"}}' 100 juno1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5qewxkxd --from relayer --chain-id temporal --node tcp://127.0.0.1:36657