package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"

	"temporal/x/icayieldmos/types"

	controllertypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		scopedKeeper  exported.ScopedKeeper

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper

		msgServiceRouter    *baseapp.MsgServiceRouter
		iCAControllerKeeper icacontrollerkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper exported.ScopedKeeper,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,

	msgServiceRouter *baseapp.MsgServiceRouter,
	iCAControllerKeeper icacontrollerkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,

		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		msgServiceRouter:    msgServiceRouter,
		iCAControllerKeeper: iCAControllerKeeper,
	}
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO Create unit test
// RegisterInterchainAccounts attempts to register an ICA account on a host chain.
func (k Keeper) RegisterInterchainAccount(ctx sdk.Context) error {
	owner := k.accountKeeper.GetModuleAddress(types.ModuleName)
	connectionId := "connection-0"

	if k.DoesInterchainAccountExist(ctx, connectionId, owner.String()) && k.IsChannelOpen(ctx, connectionId, owner.String()) {
		return nil
	}

	msg := controllertypes.NewMsgRegisterInterchainAccount(
		connectionId,
		owner.String(),
		"",
	)
	handler := k.msgServiceRouter.Handler(msg)
	res, err := handler(ctx, msg)
	if err != nil {
		return err
	}

	// IMPORTANT: emit the events!
	// the IBC relayer listens to these events
	ctx.EventManager().EmitEvents(res.GetEvents())

	// TODO: currently this gets printed to CLI even during tx simulations.
	// how can we make it only appear during actual deliverTx?
	k.Logger(ctx).Info(
		"initiated interchain account channel handshake",
		"connectionID", connectionId,
	)

	fmt.Println(res)

	return nil
}

// TODO Create unit test
func (k Keeper) GetPortId(owner string) string {
	portId := icatypes.ControllerPortPrefix + owner

	return portId
}

// TODO Create unit test
func (k Keeper) DoesInterchainAccountExist(ctx sdk.Context, connectionId string, owner string) bool {
	portId := k.GetPortId(owner)
	_, found := k.iCAControllerKeeper.GetActiveChannelID(ctx, connectionId, portId)

	return found
}

// TODO Create unit test
func (k Keeper) IsChannelOpen(ctx sdk.Context, connectionId string, owner string) bool {
	portId := k.GetPortId(owner)
	_, found := k.iCAControllerKeeper.GetOpenActiveChannel(ctx, connectionId, portId)

	return found
}

// TODO Create unit test
func (k Keeper) SendMsgSend(ctx sdk.Context) error {
	amount := sdk.NewCoin("ujunox", sdk.NewInt(10000))

	var msgs []proto.Message
	msgs = append(msgs, &banktypes.MsgSend{
		FromAddress: "juno1rlyxhrlyj7g0hwcxgkf4eqtq26l7ypdgn4lxp8whralfzp0wc9ts609at8",
		ToAddress:   "juno1qk56sklsv5xx4eh29su5fdw08gcyl9h4kl9sepl9nm0slm2lffhq6vj3pd",
		Amount:      sdk.NewCoins(amount),
	})

	msgs = append(msgs, &stakingtypes.MsgDelegate{
		DelegatorAddress: "juno1rlyxhrlyj7g0hwcxgkf4eqtq26l7ypdgn4lxp8whralfzp0wc9ts609at8",
		ValidatorAddress: "junovaloper14qekdkj2nmmwea4ufg9n002a3pud23y8vd4ld4",
		Amount:           amount,
	})

	err := k.SubmitTxs(ctx, msgs)
	if err != nil {
		return err
	}

	return nil
}

// TODO Create unit test
// SubmitTxs submits an ICA transaction containing multiple messages
func (k Keeper) SubmitTxs(ctx sdk.Context, msgs []proto.Message) error {
	owner := k.accountKeeper.GetModuleAddress(types.ModuleName).String()

	var protoMsgs []proto.Message
	for _, msg := range msgs {
		protoMsgs = append(protoMsgs, msg)
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, protoMsgs)
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	msg := icacontrollertypes.NewMsgSendTx(owner, "connection-0", 3_600_000_000_000, packetData)

	handler := k.msgServiceRouter.Handler(msg)
	res, err := handler(ctx, msg)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(res.GetEvents())

	return nil
}
