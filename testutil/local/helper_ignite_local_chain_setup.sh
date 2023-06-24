#!/bin/bash

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

cd $HOME

git clone $2 $1

cd $HOME/$1

git checkout $3

cp $SCRIPTPATH/$1/config.yml $HOME/$1/config.yml

ignite chain init
