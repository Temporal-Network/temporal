#!/bin/bash

rm -r ~/.temporald/keyring-test/
rm -r ~/.temporald/config/gentx/
rm  ~/.temporal/config/gentx/gentx-30c040a4b9ca0d9bc3e62efd2ded9c860a8df799.json

VAL_1_MNEMONIC="ocean toe void inquiry skin exotic artist teach daughter theory melt theme desk more trip person maid agent mesh prefer loop oyster member exact"
VAL_1_ADDRESS="temporal1nhvm23r9pu593sgdu660whdz66pg382czn5quy"

CMD="temporald"
DENOM="utprl"
CHAIN_ID="temporal"

$CMD tendermint unsafe-reset-all
$CMD tendermint reset-state

echo "$VAL_1_MNEMONIC" | $CMD keys add $VAL_1_ADDRESS --recover --keyring-backend=test 2>&1

$CMD init NODE1 --chain-id $CHAIN_ID --overwrite --staking-bond-denom $DENOM

$CMD add-genesis-account $VAL_1_ADDRESS 10000000$DENOM

$CMD gentx $VAL_1_ADDRESS 1000000$DENOM --chain-id $CHAIN_ID --keyring-backend test

$CMD collect-gentxs

$CMD start