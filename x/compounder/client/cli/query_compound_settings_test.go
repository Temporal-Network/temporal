package cli_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/app/apptesting"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/temporal-zone/temporal/testutil/network"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/compounder/client/cli"
	"github.com/temporal-zone/temporal/x/compounder/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithCompoundSettingsObjects(t *testing.T, n int) (*network.Network, []types.CompoundSettings) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		compoundSettings := types.CompoundSettings{
			Delegator: apptesting.GetRandomAddress().String(),
			Frequency: &types.Frequency{
				OnceEvery:       sdk.NewInt(100),
				SecondsOrBlocks: "block",
			},
		}
		nullify.Fill(&compoundSettings)
		state.CompoundSettingsList = append(state.CompoundSettingsList, compoundSettings)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.CompoundSettingsList
}

func TestShowCompoundSettings(t *testing.T) {
	net, objs := networkWithCompoundSettingsObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		delegator string

		args []string
		err  error
		obj  types.CompoundSettings
	}{
		{
			desc:      "found",
			delegator: objs[0].Delegator,

			args: common,
			obj:  objs[0],
		},
		{
			desc:      "not found",
			delegator: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delegator,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCompoundSettings(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCompoundSettingsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.CompoundSettings)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.CompoundSettings),
				)
			}
		})
	}
}

func TestListCompoundSettings(t *testing.T) {
	net, objs := networkWithCompoundSettingsObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCompoundSettings(), args)
			require.NoError(t, err)
			var resp types.QueryAllCompoundSettingsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.CompoundSettings), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.CompoundSettings),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCompoundSettings(), args)
			require.NoError(t, err)
			var resp types.QueryAllCompoundSettingsResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.CompoundSettings), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.CompoundSettings),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCompoundSettings(), args)
		require.NoError(t, err)
		var resp types.QueryAllCompoundSettingsResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.CompoundSettings),
		)
	})
}
