#!/bin/bash

CMD="temporald"
DENOM="utprl"
CHAIN_ID="temporal"

$CMD tendermint unsafe-reset-all
$CMD tendermint reset-state

rm -rf ~/.temporald/keyring-test/
rm -rf ~/.temporald/config/gentx/
rm  -rf ~/.temporal/

make dev-install

VAL_1_MNEMONIC="ocean toe void inquiry skin exotic artist teach daughter theory melt theme desk more trip person maid agent mesh prefer loop oyster member exact"
VAL_1_ADDRESS="temporal1nhvm23r9pu593sgdu660whdz66pg382czn5quy"

$CMD config keyring-backend test
$CMD config node tcp://127.0.0.1:36657

echo "$VAL_1_MNEMONIC" | $CMD keys add $VAL_1_ADDRESS --recover --keyring-backend=test 2>&1

$CMD init NODE1 --chain-id $CHAIN_ID --overwrite --staking-bond-denom $DENOM

$CMD add-genesis-account $VAL_1_ADDRESS 10000000$DENOM

$CMD gentx $VAL_1_ADDRESS 1000000$DENOM --chain-id $CHAIN_ID --keyring-backend test

$CMD collect-gentxs

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
SCRIPT=$SCRIPTPATH"/no_reset_init_local.sh"
chmod +x $SCRIPT

sh $SCRIPT