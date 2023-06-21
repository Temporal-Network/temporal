#!/bin/bash

CMD="junod"
DENOM="ujunox"

mkdir ../juno
cd ../juno

cp ../temporal/testutil/local/juno/config.yml config.yml

ignite chain init

ignite chain serve -v --reset-once