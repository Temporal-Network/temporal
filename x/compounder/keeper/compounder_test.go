package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/Temporal-Network/temporal/app/apptesting"
	compTypes "github.com/Temporal-Network/temporal/x/compounder/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

//TODO: refactor all unit test to use a cases slice and t.Run() like TestShouldCompoundingHappen and TestHandleLeftOverAmount

func TestShouldCompoundingHappen(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cases := []struct {
		name        string
		cs          compTypes.CompoundSettings
		blockHeight int64
		blockTime   time.Time
		expected    bool
		expectedErr error
		prevComp    compTypes.PreviousCompounding
	}{
		{
			name: "every 5 blocks",
			cs: compTypes.CompoundSettings{
				Delegator: "delegator1",
				Frequency: &compTypes.Frequency{
					OnceEvery:       sdk.NewInt(5),
					SecondsOrBlocks: "block",
				},
			},
			blockHeight: 10,
			blockTime:   time.Now(),
			expected:    false,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompounding{
				Delegator:   "delegator1",
				Timestamp:   time.Now().Add(-time.Minute),
				BlockHeight: sdk.NewInt(100),
			},
		},
		{
			name: "every 10 blocks",
			cs: compTypes.CompoundSettings{
				Delegator: "delegator2",
				Frequency: &compTypes.Frequency{
					OnceEvery:       sdk.NewInt(10),
					SecondsOrBlocks: "block",
				},
			},
			blockHeight: 20,
			blockTime:   time.Now(),
			expected:    true,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompounding{
				Delegator:   "delegator2",
				Timestamp:   time.Now().Add(-time.Minute),
				BlockHeight: sdk.NewInt(10),
			},
		},
		{
			name: "every 3 seconds",
			cs: compTypes.CompoundSettings{
				Delegator: "delegator3",
				Frequency: &compTypes.Frequency{
					OnceEvery:       sdk.NewInt(3),
					SecondsOrBlocks: "second",
				},
			},
			blockHeight: 30,
			blockTime:   time.Now(),
			expected:    true,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompounding{
				Delegator:   "delegator3",
				Timestamp:   time.Now().Add(-time.Minute),
				BlockHeight: sdk.NewInt(100),
			},
		},
		{
			name: "every 90 seconds",
			cs: compTypes.CompoundSettings{
				Delegator: "delegator4",
				Frequency: &compTypes.Frequency{
					OnceEvery:       sdk.NewInt(90),
					SecondsOrBlocks: "second",
				},
			},
			blockHeight: 40,
			blockTime:   time.Now(),
			expected:    false,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompounding{
				Delegator:   "delegator4",
				Timestamp:   time.Now().Add(-time.Minute),
				BlockHeight: sdk.NewInt(100),
			},
		},
		{
			name: "nil Frequency",
			cs: compTypes.CompoundSettings{
				Delegator: "delegator5",
				Frequency: &compTypes.Frequency{},
			},
			blockHeight: 50,
			blockTime:   time.Now(),
			expected:    true,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompounding{
				Delegator:   "delegator5",
				Timestamp:   time.Now().Add(-time.Minute),
				BlockHeight: sdk.NewInt(100),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s.App.CompounderKeeper.SetPreviousCompounding(s.Ctx, tc.prevComp)

			actual, err := s.App.CompounderKeeper.ShouldCompoundingHappen(s.Ctx, tc.cs, tc.blockHeight, tc.blockTime)
			require.Equal(t, tc.expectedErr, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestHandleLeftOverAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cases := []struct {
		name                 string
		compoundActions      []compTypes.StakingCompoundAction
		totalCompoundPercent math.Int
		amountToCompound     sdk.Coin
		expectedResult       []compTypes.StakingCompoundAction
	}{
		{
			name: "total compound percent is 100% 1",
			compoundActions: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
			},
			totalCompoundPercent: sdk.NewInt(100),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)),
			expectedResult: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
			},
		},
		{
			name: "total compound percent is 100% 2",
			compoundActions: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(24))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(80))},
			},
			totalCompoundPercent: sdk.NewInt(100),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(105)),
			expectedResult: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(25))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(80))},
			},
		},

		{
			name: "total compound percent is 100% 2",
			compoundActions: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20))},
			},
			totalCompoundPercent: sdk.NewInt(30),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)),
			expectedResult: []compTypes.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20))},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := s.App.CompounderKeeper.HandleLeftOverAmount(tc.compoundActions, tc.totalCompoundPercent, tc.amountToCompound)
			require.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestStakingCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	delegations := distrTypes.QueryDelegationTotalRewardsResponse{
		Rewards: []distrTypes.DelegationDelegatorReward{
			{
				ValidatorAddress: "cosmosvalopr11",
				Reward: sdk.DecCoins{
					sdk.NewDecCoin("uatom", sdk.NewInt(100)),
				},
			},
			{
				ValidatorAddress: "cosmosvaloper12",
				Reward: sdk.DecCoins{
					sdk.NewDecCoin("uatom", sdk.NewInt(200)),
				},
			},
		},
	}
	walletBalance := sdk.NewCoin("mycoin", sdk.NewInt(1000))

	outstandingRewards := s.App.CompounderKeeper.StakingCompoundAmount(&delegations, walletBalance)

	require.Equal(t, sdk.NewInt(300), outstandingRewards.Amount)
	require.Equal(t, "mycoin", outstandingRewards.Denom)
}

func TestExtraCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cs := compTypes.CompoundSettings{
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}
	walletBalance := sdk.NewCoin("uatom", sdk.NewInt(200000))

	extraCompoundAmount := s.App.CompounderKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount := sdk.NewCoin("uatom", sdk.NewInt(100000))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)

	walletBalance = sdk.NewCoin("uatom", sdk.NewInt(90000))
	extraCompoundAmount = s.App.CompounderKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount = sdk.NewCoin("uatom", sdk.NewInt(0))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)

	cs.AmountToRemain = sdk.Coin{}
	extraCompoundAmount = s.App.CompounderKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount = sdk.NewCoin("uatom", sdk.NewInt(0))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)
}

// TestBuildCompoundActions tests the BuildCompoundActions function
func TestBuildCompoundActions(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cs := compTypes.CompoundSettings{
		Delegator: "cosmos1delegator",
		ValidatorSettings: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: sdk.NewInt(10),
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: sdk.NewInt(20),
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}
	amountToCompound := sdk.NewCoin("uatom", sdk.NewInt(1000000))

	totalCompoundPercent, compoundActions := s.App.CompounderKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(30), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(100000)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(200000)), compoundActions[1].Balance)

	cs = compTypes.CompoundSettings{
		Delegator: "cosmos1delegator",
		ValidatorSettings: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: sdk.NewInt(80),
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: sdk.NewInt(20),
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}

	amountToCompound = sdk.NewCoin("uatom", sdk.NewInt(10_000_000))

	totalCompoundPercent, compoundActions = s.App.CompounderKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(100), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(8_000_000)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(2_000_000)), compoundActions[1].Balance)

	cs = compTypes.CompoundSettings{
		Delegator: "cosmos1delegator",
		ValidatorSettings: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: sdk.NewInt(80),
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: sdk.NewInt(20),
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}

	amountToCompound = sdk.NewCoin("uatom", sdk.NewInt(10_000_081))

	totalCompoundPercent, compoundActions = s.App.CompounderKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(100), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(8_000_064)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(2_000_016)), compoundActions[1].Balance)
}

func TestTotalCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	delegations := &distrTypes.QueryDelegationTotalRewardsResponse{
		Rewards: []distrTypes.DelegationDelegatorReward{
			{
				ValidatorAddress: "cosmosvalopr11",
				Reward: sdk.DecCoins{
					sdk.NewDecCoin("uatom", sdk.NewInt(100)),
				},
			},
			{
				ValidatorAddress: "cosmosvaloper12",
				Reward: sdk.DecCoins{
					sdk.NewDecCoin("uatom", sdk.NewInt(200)),
				},
			},
		},
	}

	walletBalance := sdk.NewCoin("uatom", sdk.NewInt(1000))
	cs := compTypes.CompoundSettings{
		Delegator:         "cosmos11",
		ValidatorSettings: []*compTypes.ValidatorSetting{},
		AmountToRemain:    sdk.NewCoin("uatom", sdk.NewInt(500)),
	}

	total := s.App.CompounderKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 800 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 800)
	}

	walletBalance = sdk.NewCoin("uatom", sdk.NewInt(1500))
	cs = compTypes.CompoundSettings{
		Delegator:         "cosmos12",
		ValidatorSettings: []*compTypes.ValidatorSetting{},
		AmountToRemain:    sdk.NewCoin("uatom", sdk.NewInt(1500)),
	}

	total = s.App.CompounderKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 300 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 300)
	}

	delegations = &distrTypes.QueryDelegationTotalRewardsResponse{}
	walletBalance = sdk.NewCoin("uatom", sdk.NewInt(1500))
	cs = compTypes.CompoundSettings{
		Delegator:         "cosmos13",
		ValidatorSettings: []*compTypes.ValidatorSetting{},
		AmountToRemain:    sdk.NewCoin("uatom", sdk.NewInt(1500)),
	}

	total = s.App.CompounderKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 0 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 0)
	}

	walletBalance = sdk.NewCoin("uatom", sdk.NewInt(0))
	cs = compTypes.CompoundSettings{
		Delegator:         "cosmos13",
		ValidatorSettings: []*compTypes.ValidatorSetting{},
		AmountToRemain:    sdk.NewCoin("uatom", sdk.NewInt(1500)),
	}

	total = s.App.CompounderKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 0 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 0)
	}

	walletBalance = sdk.NewCoin("uatom", sdk.NewInt(1500))
	cs = compTypes.CompoundSettings{
		Delegator:         "cosmos13",
		ValidatorSettings: []*compTypes.ValidatorSetting{},
		AmountToRemain:    sdk.NewCoin("uatom", sdk.NewInt(1000)),
	}

	total = s.App.CompounderKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 500 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 500)
	}
}

func TestRecordCompounding(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	// Test case 1: New delegator record
	address1 := "address1"
	blockTime1 := time.Now().UTC()
	blockHeight := sdk.NewInt(2)
	s.App.CompounderKeeper.RecordCompounding(s.Ctx, address1, blockTime1, blockHeight)

	value, found := s.App.CompounderKeeper.GetPreviousCompounding(s.Ctx, address1)
	if !found {
		t.Fatalf("Expected to find the address in the store but not found")
	}
	if value.Delegator != address1 {
		t.Fatalf("Expected delegator address to be %s but got %s", address1, value.Delegator)
	}
	if value.Timestamp != blockTime1 {
		t.Fatalf("Expected block time to be %s but got %s", blockTime1, value.Timestamp)
	}

	// Test case 2: Update existing record
	address2 := "address2"
	blockTime2 := time.Now().Add(time.Duration(24) * time.Hour).UTC()

	s.App.CompounderKeeper.RecordCompounding(s.Ctx, address2, blockTime2, blockHeight)

	value, found = s.App.CompounderKeeper.GetPreviousCompounding(s.Ctx, address2)
	if !found {
		t.Fatalf("Expected to find the address in the store but not found")
	}
	if value.Delegator != address2 {
		t.Fatalf("Expected delegator address to be %s but got %s", address2, value.Delegator)
	}
	if value.Timestamp != blockTime2 {
		t.Fatalf("Expected block time to be %s but got %s", blockTime2, value.Timestamp)
	}
}

func TestCalculateCompoundingAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	var tests = []struct {
		name              string
		rewardAmount      math.Int
		percentToCompound math.Int
		expectedAmount    math.Int
	}{
		{
			name:              "Test case 1",
			rewardAmount:      sdk.NewInt(100),
			percentToCompound: sdk.NewInt(10),
			expectedAmount:    sdk.NewInt(10),
		},
		{
			name:              "Test case 2",
			rewardAmount:      sdk.NewInt(100),
			percentToCompound: sdk.NewInt(30),
			expectedAmount:    sdk.NewInt(30),
		},
		{
			name:              "Test case 3",
			rewardAmount:      sdk.NewInt(200),
			percentToCompound: sdk.NewInt(20),
			expectedAmount:    sdk.NewInt(40),
		},
		{
			name:              "Test case 4",
			rewardAmount:      sdk.NewInt(200),
			percentToCompound: sdk.NewInt(0),
			expectedAmount:    sdk.NewInt(0),
		},
		{
			name:              "Zero case",
			rewardAmount:      sdk.NewInt(0),
			percentToCompound: math.NewInt(50),
			expectedAmount:    math.NewInt(0),
		},
	}

	for _, test := range tests {
		result := s.App.CompounderKeeper.CalculateCompoundAmount(sdk.NewCoin("test", test.rewardAmount), test.percentToCompound)
		if !result.Sub(test.expectedAmount).IsZero() {
			t.Errorf("Test case %s failed: expected %s but got %s", test.name, test.expectedAmount, result)
		}
	}
}
