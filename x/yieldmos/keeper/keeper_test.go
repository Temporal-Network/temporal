package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"math"
	"math/rand"
	"temporal/app"
	"temporal/app/apptesting"
	"temporal/cmd/temporald/cmd"
	"temporal/x/yieldmos/types"
	"testing"
	"time"
)

func setupValidators(t *testing.T, ctx sdk.Context, stakingKeeper stakingkeeper.Keeper, testAmount int, validatorAddresses []sdk.ValAddress) []stakingTypes.Validator {
	PKs := simapp.CreateTestPubKeys(testAmount)
	validators := make([]stakingTypes.Validator, testAmount)
	for i := 0; i < testAmount; i++ {
		validators[i] = teststaking.NewValidator(t, validatorAddresses[i], PKs[i])
		validators[i], _ = validators[i].AddTokensFromDel(sdk.NewInt(10))
		validators[i] = updateValidator(stakingKeeper, ctx, validators[i], true)
	}

	return validators
}

// update validator for testing
func updateValidator(keeper stakingkeeper.Keeper, ctx sdk.Context, validator stakingTypes.Validator, apply bool) stakingTypes.Validator {
	keeper.SetValidator(ctx, validator)

	keeper.SetNewValidatorByPowerIndex(ctx, validator)

	if !apply {
		ctx, _ = ctx.CacheContext()
	}
	_, err := keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		panic(err)
	}

	validator, found := keeper.GetValidator(ctx, validator.GetOperator())
	if !found {
		panic("validator expected but not found")
	}

	return validator
}

func setupAccounts(app *app.App, ctx sdk.Context, addressAmount int) []sdk.AccAddress {
	var addresses []sdk.AccAddress

	for i := 0; i < addressAmount; i++ {
		addresses = append(addresses, getRandomAddress())
	}

	addressAmountInt := sdk.NewInt(int64(addressAmount))

	coinAmount := sdk.NewInt(1000000).Mul(addressAmountInt)
	coins := sdk.NewCoins(sdk.NewCoin(cmd.BaseCoinUnit, coinAmount))
	err := app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	if err != nil {
		panic(err)
	}

	coins = sdk.NewCoins(sdk.NewCoin(cmd.BaseCoinUnit, coinAmount.Quo(addressAmountInt)))
	for _, address := range addresses {
		err := app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, address, coins)
		if err != nil {
			panic(err)
		}
	}

	return addresses
}

func getRandomAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()

	return sdk.AccAddress(pk.Address())
}

func TestNewDelegationTimestamp(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	for i := 0; i <= 1000; i++ {
		randomAmount := sdk.NewInt(rand.Int63n(math.MaxInt))
		testDelegationTimestamp := s.App.YieldmosKeeper.NewDelegationTimestamp(s.Ctx, randomAmount)
		newDelegationTimestamp := types.DelegationTimestamp{
			Timestamp: time.Unix(s.Ctx.BlockTime().Unix(), 0),
			Balance: sdk.NewCoin(
				s.App.StakingKeeper.BondDenom(s.Ctx),
				randomAmount,
			),
		}

		require.Equal(t, newDelegationTimestamp, testDelegationTimestamp)
	}
}

func TestNewDelegationHistory(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	for i := 0; i <= 1000; i++ {
		delegatorAddress := getRandomAddress()
		delegatedAmount := sdk.NewInt(rand.Int63n(math.MaxInt))
		delegationTimestamp := s.App.YieldmosKeeper.NewDelegationTimestamp(s.Ctx, delegatedAmount)
		testDelegationHistory := s.App.YieldmosKeeper.NewDelegationHistory(s.Ctx, delegatorAddress, delegatedAmount)

		newDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{&delegationTimestamp},
		}

		require.Equal(t, newDelegationHistory, testDelegationHistory)
	}
}

func TestAddDelegationTimestamp(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	for i := 0; i <= 1000; i++ {
		delegatorAddress := getRandomAddress()
		delegatedAmount := sdk.NewInt(rand.Int63n(math.MaxInt))
		testDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{},
		}

		testDelegationHistory = s.App.YieldmosKeeper.AddDelegationTimestamp(s.Ctx, delegatedAmount, testDelegationHistory)

		newDelegationTimestamp := s.App.YieldmosKeeper.NewDelegationTimestamp(s.Ctx, delegatedAmount)
		newDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{&newDelegationTimestamp},
		}

		require.Equal(t, newDelegationHistory, testDelegationHistory)
	}
}

func TestPruneDelegationHistory(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	for i := 0; i <= 1000; i++ {
		delegatorAddress := getRandomAddress()

		testDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{},
		}

		for c := 0; c <= 100; c++ {
			delegatedAmount := sdk.NewInt(rand.Int63n(math.MaxInt))
			testDelegationHistory = s.App.YieldmosKeeper.AddDelegationTimestamp(s.Ctx, delegatedAmount, testDelegationHistory)
			testDelegationHistory = s.App.YieldmosKeeper.AddDelegationTimestamp(s.Ctx, sdk.NewInt(0), testDelegationHistory)
		}

		testDelegationHistory = s.App.YieldmosKeeper.PruneDelegationHistory(testDelegationHistory)

		for _, delegationTimestamp := range testDelegationHistory.GetHistory() {
			require.NotEqual(t, delegationTimestamp.GetBalance().Amount, sdk.NewInt(0))
		}
	}
}

func TestAdjustDelegationTimestamps(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	var err error

	for i := 0; i <= 1000; i++ {
		delegatorAddress := getRandomAddress()

		testDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{},
		}

		for c := 0; c <= 100; c++ {
			delegatedAmount := sdk.NewInt(int64(rand.Intn(math.MaxInt)))
			testDelegationHistory = s.App.YieldmosKeeper.AddDelegationTimestamp(s.Ctx, delegatedAmount, testDelegationHistory)
		}

		newDelegatedAmount := sdk.NewInt(int64(rand.Intn(math.MaxInt)))
		difference := s.App.YieldmosKeeper.CalcDelegationHistoryDifference(newDelegatedAmount, testDelegationHistory)
		require.Equal(t, difference.IsNegative(), true)

		testDelegationHistory, err = s.App.YieldmosKeeper.AdjustDelegationTimestamps(testDelegationHistory, difference)
		require.NoError(t, err)

		newDifference := s.App.YieldmosKeeper.CalcDelegationHistoryDifference(newDelegatedAmount, testDelegationHistory)
		require.Equal(t, newDifference.IsZero(), true)
	}
}

func TestCalcDelegationHistoryDifference(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	for i := 0; i <= 1000; i++ {
		delegatorAddress := getRandomAddress()
		totalDelegatedAmount := sdk.NewInt(0)

		testDelegationHistory := types.DelegationHistory{
			Address: delegatorAddress.String(),
			History: []*types.DelegationTimestamp{},
		}

		for c := 0; c <= 100; c++ {
			delegatedAmount := sdk.NewInt(int64(rand.Intn(math.MaxInt)))
			totalDelegatedAmount = totalDelegatedAmount.Add(delegatedAmount)
			testDelegationHistory = s.App.YieldmosKeeper.AddDelegationTimestamp(s.Ctx, delegatedAmount, testDelegationHistory)
		}

		newDelegatedAmount := sdk.NewInt(int64(rand.Intn(math.MaxInt)))
		difference := s.App.YieldmosKeeper.CalcDelegationHistoryDifference(newDelegatedAmount, testDelegationHistory)

		require.Equal(t, newDelegatedAmount.Sub(totalDelegatedAmount), difference)
	}
}

func TestCalcTotalDelegatedAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	//TODO: Greater than 900, breaks something???
	testAmount := 900

	delegatorAddresses := setupAccounts(s.App, s.Ctx, testAmount)
	validatorAddresses := simapp.ConvertAddrsToValAddrs(delegatorAddresses)

	validators := setupValidators(t, s.Ctx, s.App.StakingKeeper, testAmount, validatorAddresses)
	require.Len(t, validators, testAmount)

	delegations := make([]sdk.Dec, testAmount)
	for i := 0; i < testAmount; i++ {
		maxDelegations := rand.Intn(100)

		if delegations[i].IsNil() {
			delegations[i] = sdk.NewDec(0)
		}

		for c := 0; c < maxDelegations; c++ {
			bonding := sdk.NewDec(int64(rand.Intn(math.MaxInt)))
			require.True(t, bonding.IsPositive())

			delegations[i] = delegations[i].Add(bonding)

			testBond := stakingTypes.NewDelegation(delegatorAddresses[i], validatorAddresses[i], bonding)
			currentDelegations, found := s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
			if found {
				testBond.Shares = currentDelegations.Shares.Add(bonding)
			}

			s.App.StakingKeeper.SetDelegation(s.Ctx, testBond)

			bond, found := s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
			require.True(t, found)
			require.Equal(t, delegations[i], bond.GetShares())
		}

		bondAmount := s.App.YieldmosKeeper.CalcTotalDelegatedAmount(s.Ctx, delegatorAddresses[i])
		require.Equal(t, delegations[i].TruncateInt().String(), bondAmount.String())
	}
}

// TODO: This bypasses the hook setup in hooks.go, add a test that validates the "AfterDelegationModified" hook
// TODO: Test the MsgDelegate Hook.
func TestCheckDelegationHistoryRecords(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	//TODO: Greater than 900, breaks something???
	testAmount := 900

	delegatorAddresses := setupAccounts(s.App, s.Ctx, testAmount)
	validatorAddresses := simapp.ConvertAddrsToValAddrs(delegatorAddresses)

	validators := setupValidators(t, s.Ctx, s.App.StakingKeeper, testAmount, validatorAddresses)
	require.Len(t, validators, testAmount)

	delegations := make([]sdk.Dec, testAmount)
	for i := 0; i < testAmount; i++ {
		maxDelegations := rand.Intn(100)

		if delegations[i].IsNil() {
			delegations[i] = sdk.NewDec(0)
		}

		for c := 0; c < maxDelegations; c++ {
			bonding := sdk.NewDec(int64(rand.Intn(math.MaxInt)))
			require.True(t, bonding.IsPositive())

			delegations[i] = delegations[i].Add(bonding)

			testBond := stakingTypes.NewDelegation(delegatorAddresses[i], validatorAddresses[i], bonding)
			currentDelegation, found := s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
			if found {
				testBond.Shares = currentDelegation.Shares.Add(bonding)
			}

			s.App.StakingKeeper.SetDelegation(s.Ctx, testBond)

			bond, found := s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
			require.True(t, found)
			require.Equal(t, delegations[i], bond.GetShares())

			err := s.App.YieldmosKeeper.CheckDelegationHistoryRecords(s.Ctx, delegatorAddresses[i])
			require.NoError(t, err)

			delegationHistory, found := s.App.YieldmosKeeper.GetDelegationHistory(s.Ctx, delegatorAddresses[i].String())
			require.True(t, found)

			totalDelegationHistory := s.App.YieldmosKeeper.CalcDelegationHistoryDifference(sdk.NewInt(0), delegationHistory)
			require.Equal(t, delegations[i].TruncateInt().String(), totalDelegationHistory.Abs().String())
		}

		currentDelegation, found := s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
		if found && currentDelegation.Shares.GTE(sdk.NewDec(500)) {
			for c := 0; c < 5; c++ {
				decreaseBonded := sdk.NewDec(int64(rand.Intn(100)))
				delegations[i] = delegations[i].Sub(decreaseBonded)
				testBond := stakingTypes.NewDelegation(delegatorAddresses[i], validatorAddresses[i], sdk.NewDec(0))
				testBond.Shares = currentDelegation.Shares.Sub(decreaseBonded)

				s.App.StakingKeeper.SetDelegation(s.Ctx, testBond)

				err := s.App.YieldmosKeeper.CheckDelegationHistoryRecords(s.Ctx, delegatorAddresses[i])
				require.NoError(t, err)

				delegationHistory, found := s.App.YieldmosKeeper.GetDelegationHistory(s.Ctx, delegatorAddresses[i].String())
				require.True(t, found)

				totalDelegationHistory := s.App.YieldmosKeeper.CalcDelegationHistoryDifference(delegations[i].TruncateInt(), delegationHistory)
				require.True(t, totalDelegationHistory.IsZero())

				currentDelegation, _ = s.App.StakingKeeper.GetDelegation(s.Ctx, delegatorAddresses[i], validatorAddresses[i])
			}
		}
	}
}
