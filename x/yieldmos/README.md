# Yieldmos Module

This module saves and prunes staking history of the `utprl` staking token, registers interchain accounts on other networks (in progress) and broadcasts ICA messages to other networks (in progress).

## Staking History

Stores in state when a staking or unstaking action happens by registering a hook with the staking module.

## Registers interchain account

Attempt to register an interchain account on a remote network

## Broadcast interchain accounts messages

Attempt to broadcast an ICA message to a remote network

## Queries

`DelegationHistoryAll` - Gets all accounts `DelegationHistory`

`DelegationHistory` - Get a single addressess `DelegationHistory`

`Params` - Get all `yieldmos` module parameters
