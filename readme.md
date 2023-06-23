# Temporal Zone
**Temporal Zone** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Custom Modules

[`compounder`](https://github.com/temporal-zone/temporal/blob/main/x/compounder/README.md) module is used to compound `utprl` staked positions.

[`yieldmos`](https://github.com/temporal-zone/temporal/blob/main/x/yieldmos/README.md) module is used for saving staking history, registering ICA accounts (in 
progress) and broadcasting ICA messages (in progress).

## Custom Contracts

Osmosis/Mars: https://github.com/yieldmos/ac-outpost

## How to locally run the chain

Install gcc

Install make

Install [go](https://go.dev/doc/install)

Optional Install: docker if you plan on adding or updating protos.

On first run you might need to add a go path:
```
export PATH=$PATH:~/go/bin/
```


## Using Make

Build the binary:
```
make install
```

Start the chain:
```
temporald start
```

## Using Ignite

Install [Ignite](https://docs.ignite.com/welcome/install)

From the root of the temporal directory run:
```
ignite chain serve -v
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

Various parameters can be configured with `config.yml`. To learn more, see the [Ignite Config docs](https://docs.ignite.com/references/config).

## Attribution

Thank you to the following projects for inspiration, code and helping us get started:

Osmosis

Stride

Mars Hub

Interchain Queries by Ingenuity Build

Interchain Accounts Demo
