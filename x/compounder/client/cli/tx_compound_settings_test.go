package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/Temporal-Network/temporal/testutil/network"
	"github.com/Temporal-Network/temporal/x/compounder/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateCompoundSettings(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"null", "10token", "null"}
	for _, tc := range []struct {
		desc      string
		delegator string

		args []string
		err  error
		code uint32
	}{
		{
			delegator: strconv.Itoa(0),

			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delegator,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSettings(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateCompoundSettings(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"null", "10token", "null"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}

	for _, tc := range []struct {
		desc      string
		delegator string

		args []string
		code uint32
		err  error
	}{
		{
			desc:      "valid",
			delegator: strconv.Itoa(0),

			args: common,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delegator,
			}
			args = append(args, fields...)
			args = append(args, common...)
			_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSettings(), args)
			require.NoError(t, err)

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateCompoundSettings(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateCompoundSettingsNoCreate(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"null", "10token", "null"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}

	for _, tc := range []struct {
		desc      string
		delegator string

		args []string
		code uint32
		err  error
	}{
		{
			desc:      "key not found",
			delegator: strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delegator,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateCompoundSettings(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestDeleteCompoundSettings(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"null", "10token", "null"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSettings(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc      string
		delegator string

		args []string
		code uint32
		err  error
	}{
		{
			desc:      "valid",
			delegator: strconv.Itoa(0),

			args: common,
		},
		{
			desc:      "key not found",
			delegator: strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delegator,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteCompoundSettings(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
