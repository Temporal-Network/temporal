# Temporal Network
**Temporal Network** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Custom Modules

`compounder` module is used to compound `utprl` staked positions. More information can be found [here.](https://github.com/Temporal-Network/temporal/blob/main/x/compounder/README.md)

`yieldmos` module is used for saving staking history and registering ICA accounts (in progress) and broadcasting ICA messages (in progress). More information can be found [here.](https://github.com/Temporal-Network/temporal/blob/main/x/yieldmos/README.md)


## How to locally run the chain

Install [Ignite](https://docs.ignite.com/welcome/install)

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

Various parameters can be configured with `config.yml`. To learn more, see the [Ignite Config docs](https://docs.ignite.com/references/config).


