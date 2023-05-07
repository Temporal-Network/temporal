# Compounder Module

This module compounds an accounts `utplr` staking rewards every `x` amount of blocks or seconds, minimum of 100.

## Messages

`MsgCreateCompoundSettings` - Used to create new `CompoundSettings`

`MsgUpdateCompoundSettings` - Used to update existing `CompoundSettings`

`MsgDeleteCompoundSettings` - Used to disable compounding by deleting the existing `CompoundSettings`

## CompoundSettings

```json
{
  "compoundSettings": {
    "delegator": "string",
    "validatorSettings": [
      {
        "validatorAddress": "string",
        "percentToCompound": "string"
      }
    ],
    "amountToRemain": {
      "denom": "string",
      "amount": "string"
    },
    "frequency": {
      "onceEvery": "string",
      "secondsOrBlocks": "string"
    }
  }
}
```

## Queries

`CompoundSettingsAll` - Gets all accounts `CompoundSettings`

`CompoundSettings` - Get a single accounts `CompoundSettings`

`Params` - Get all `compounder` module parameters
