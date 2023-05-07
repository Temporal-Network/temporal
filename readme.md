# Temporal Network
**Temporal Network** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Custom Modules

[`compounder`](https://github.com/Temporal-Network/temporal/blob/main/x/compounder/README.md) module is used to compound `utprl` staked positions.

[`yieldmos`](https://github.com/Temporal-Network/temporal/blob/main/x/yieldmos/README.md) module is used for saving staking history, registering ICA accounts (in 
progress) and broadcasting ICA messages (in progress).

## Custom Contracts

Osmosis/Mars: https://github.com/kakucodes/delphilabs_hackathon_outpost

## How to locally run the chain

Install [Ignite](https://docs.ignite.com/welcome/install)

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

Various parameters can be configured with `config.yml`. To learn more, see the [Ignite Config docs](https://docs.ignite.com/references/config).

## Attribution

Thank you to the following projects for inspiration and helping us get started:

Osmosis

Stride

Mars Hub

Interchain accounts demo
