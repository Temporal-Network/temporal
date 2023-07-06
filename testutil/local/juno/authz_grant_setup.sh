#!/bin/bash

CMD_JUNO="junod"
CMD_TEMPORAL="temporald"

#$CMD_TEMPORAL tx icayieldmos create-connection connection-0 --from aliceT --chain-id temporal --node tcp://127.0.0.1:36657 -y

#sleep 2

$CMD_JUNO tx authz grant $1 generic --msg-type /cosmos.staking.v1beta1.MsgDelegate --from aliceJ --chain-id juno -y

#$CMD_TEMPORAL tx icayieldmos send-test-contract-messages connection-0 $1 $2 $3 --from aliceT --chain-id temporal --node tcp://127.0.0.1:36657