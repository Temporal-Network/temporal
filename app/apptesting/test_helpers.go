package apptesting

import (
	"github.com/Temporal-Network/temporal/app"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	ChainID = "temporal-test"
)

type SuitelessAppTestHelper struct {
	App *app.App
	Ctx sdk.Context
}

// Instantiates an TestHelper without the test suite
// This is for testing scenarios where we simply need the setup function to run,
// and need access to the TestHelper attributes and keepers (e.g. genesis tests)
func SetupSuitelessTestHelper() SuitelessAppTestHelper {
	s := SuitelessAppTestHelper{}
	s.App = app.InitTestApp(true)
	s.Ctx = s.App.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: ChainID})
	return s
}

func GetRandomAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()

	return sdk.AccAddress(pk.Address())
}
