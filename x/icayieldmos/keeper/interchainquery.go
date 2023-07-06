package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cast"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
	icqtypes "github.com/temporal-zone/temporal/x/interchainquery/types"
)

const (
	ICQCallbackID_WithdrawalBalance = "withdrawalbalance"
	ICQCallbackID_FeeBalance        = "feebalance"
	ICQCallbackID_DelegationKV      = "delegationKV"
	ICQCallbackID_DelegationTypeUrl = "delegationTypeUrl"
	ICQCallbackID_Validator         = "validator"
)

// ICQCallbacks wrapper struct for icayieldmos keeper
type ICQCallback func(Keeper, sdk.Context, []byte, icqtypes.Query) error

type ICQCallbacks struct {
	k         Keeper
	callbacks map[string]ICQCallback
}

var _ icqtypes.QueryCallbacks = ICQCallbacks{}

func (k Keeper) ICQCallbackHandler() ICQCallbacks {
	return ICQCallbacks{k, make(map[string]ICQCallback)}
}

func (c ICQCallbacks) CallICQCallback(ctx sdk.Context, id string, args []byte, query icqtypes.Query) error {
	return c.callbacks[id](c.k, ctx, args, query)
}

func (c ICQCallbacks) HasICQCallback(id string) bool {
	_, found := c.callbacks[id]
	return found
}

func (c ICQCallbacks) AddICQCallback(id string, fn interface{}) icqtypes.QueryCallbacks {
	c.callbacks[id] = fn.(ICQCallback)
	return c
}

func (c ICQCallbacks) RegisterICQCallbacks() icqtypes.QueryCallbacks {
	return c.AddICQCallback(ICQCallbackID_DelegationKV, ICQCallback(TestCallbackKV)).
		AddICQCallback(ICQCallbackID_DelegationTypeUrl, ICQCallback(TestCallbackFullType))
	//AddICQCallback(ICQCallbackID_WithdrawalBalance, ICQCallback(WithdrawalBalanceCallback)).
	//AddICQCallback(ICQCallbackID_FeeBalance, ICQCallback(FeeBalanceCallback)).
	//AddICQCallback(ICQCallbackID_Delegation, ICQCallback(DelegatorSharesCallback)).
	//AddICQCallback(ICQCallbackID_Validator, ICQCallback(ValidatorExchangeRateCallback))
}

func TestCallbackKV(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	fmt.Println("KV : *****************************************************************************************************************************************************************************************************************************************************")
	fmt.Println(query.String())

	return nil
}

func TestCallbackFullType(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	fmt.Println("Full Type **************************************************************************************************************************************************************************************************************************************")
	fmt.Println(query.String())

	return nil
}

// QueryAllDelegatorDelegations submits an ICQ to get all delegator delegations
func (k Keeper) QueryDelegations(ctx sdk.Context) error {
	remoteDelegatorAddress := "juno1scund0fnwr7xkvteysyysajdk532vlxge3q68d"
	remoteValidatorAddress := "junovaloper1scund0fnwr7xkvteysyysajdk532vlxgxvk4u5"

	if remoteDelegatorAddress == "" || len(remoteDelegatorAddress) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "delegator address invalid: %s", remoteDelegatorAddress)
	}

	if remoteValidatorAddress == "" || len(remoteValidatorAddress) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "validator address invalid: %s", remoteValidatorAddress)
	}

	currBlockTime, err := cast.ToUint64E(ctx.BlockTime().UnixNano())
	if err != nil {
		return err
	}

	k.KVQueryDelegatorValidator(ctx, remoteDelegatorAddress, remoteValidatorAddress, currBlockTime)

	//k.KVQueryDelegatorAll(ctx, remoteDelegatorAddress, currBlockTime)

	//k.FullTypeQuery(ctx, remoteDelegatorAddress, currBlockTime)

	return nil
}

func (k Keeper) KVQueryDelegatorValidator(ctx sdk.Context, remoteDelegatorAddress string, remoteValidatorAddress string, currBlockTime uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("SENDING KV QUERY DEL VAL"))

	_, delegatorAddressBz, err := bech32.DecodeAndConvert(remoteDelegatorAddress)
	if err != nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid delegator address, could not decode (%s)", err.Error()))
	}

	_, validatorAddressBz, err := bech32.DecodeAndConvert(remoteValidatorAddress)
	if err != nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid validator address, could not decode (%s)", err.Error()))
	}

	queryData := stakingtypes.GetDelegationKey(delegatorAddressBz, validatorAddressBz)

	if err := k.interchainquery.MakeRequest(
		ctx,
		types.ModuleName,
		ICQCallbackID_DelegationKV,
		"juno",
		"connection-0",
		icqtypes.STAKING_STORE_QUERY_WITH_PROOF,
		queryData,
		currBlockTime+ICA_TIMEOUT,
	); err != nil {
		k.Logger(ctx).Info(fmt.Sprintf("Error submitting KV ICQ for delegation, error : %s", err.Error()))
	}
}

func (k Keeper) KVQueryDelegatorAll(ctx sdk.Context, remoteDelegatorAddress string, currBlockTime uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("SENDING KV QUERY DEL ALL"))

	_, delegatorAddressBz, err := bech32.DecodeAndConvert(remoteDelegatorAddress)
	if err != nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid delegator address, could not decode (%s)", err.Error()))
	}

	queryData := stakingtypes.GetDelegationsKey(delegatorAddressBz)

	if err := k.interchainquery.MakeRequest(
		ctx,
		types.ModuleName,
		ICQCallbackID_DelegationKV,
		"juno",
		"connection-0",
		icqtypes.STAKING_STORE_QUERY_WITH_PROOF,
		queryData,
		currBlockTime+ICA_TIMEOUT,
	); err != nil {
		k.Logger(ctx).Info(fmt.Sprintf("Error submitting KV ICQ for delegation, error : %s", err.Error()))
	}
}

func (k Keeper) FullTypeQuery(ctx sdk.Context, remoteDelegatorAddress string, currBlockTime uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("SENDING TYPE URL QUERY"))
	rewardsQuery := distrtypes.QueryDelegationTotalRewardsRequest{DelegatorAddress: remoteDelegatorAddress}
	bz := k.cdc.MustMarshal(&rewardsQuery)

	if err := k.interchainquery.MakeRequest(
		ctx,
		types.ModuleName,
		ICQCallbackID_DelegationTypeUrl,
		"juno",
		"connection-0",
		"/cosmos.distribution.v1beta1.Query/DelegationTotalRewards",
		bz,
		currBlockTime+ICA_TIMEOUT,
	); err != nil {
		k.Logger(ctx).Info(fmt.Sprintf("Error submitting Full Type ICQ for delegation, error : %s", err.Error()))
	}
}
