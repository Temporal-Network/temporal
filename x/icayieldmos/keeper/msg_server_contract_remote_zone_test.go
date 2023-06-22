package keeper_test

// TODO These are commented out as the keeper for the account module is not inited in any of these tests and it is required for any of these to even be run.
/*
func TestContractRemoteZoneMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateContractRemoteZone(ctx, &types.MsgCreateContractRemoteZone{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestContractRemoteZoneMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateContractRemoteZone
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateContractRemoteZone{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateContractRemoteZone{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateContractRemoteZone{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateContractRemoteZone(ctx, &types.MsgCreateContractRemoteZone{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateContractRemoteZone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestContractRemoteZoneMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteContractRemoteZone
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteContractRemoteZone{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteContractRemoteZone{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteContractRemoteZone{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateContractRemoteZone(ctx, &types.MsgCreateContractRemoteZone{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteContractRemoteZone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
*/
