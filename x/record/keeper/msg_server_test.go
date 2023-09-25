package keeper_test

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/temporal-zone/temporal/testutil/sample"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/record/keeper"
	"github.com/temporal-zone/temporal/x/record/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RecordKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestCreateUserInstruction(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)

	localAddress := sample.AccAddress()
	remoteAddress, err := types.DeriveRemoteAddress(localAddress, "juno")
	require.NoError(t, err)

	tests := []struct {
		name    string
		userIns types.MsgCreateUserInstruction
		err     error
	}{
		{
			name: "invalid expiration",
			userIns: types.MsgCreateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Date(0000, 01, 01, 0, 0, 0, 0, time.UTC),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			userIns: types.MsgCreateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
		}, {
			name: "already exists",
			userIns: types.MsgCreateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().UTC().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ms.CreateUserInstruction(ctx, &tt.userIns)

			if tt.err == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}

func TestDeleteUserInstruction(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	ms := keeper.NewMsgServerImpl(*k)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotNil(t, k)

	localAddress := sample.AccAddress()
	junoRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "juno")
	require.NoError(t, err)
	osmosisRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "osmosis")
	require.NoError(t, err)
	migalooRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "migaloo")
	require.NoError(t, err)

	junoCon1 := "junoContractAddress1"
	junoCon2 := "junoContractAddress2"

	UerInsJuno1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   junoRemoteAddress,
		ChainId:         "juno-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: junoCon1,
	}

	userInsJuno2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   junoRemoteAddress,
		ChainId:         "juno-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: junoCon2,
	}

	userInsOsmo1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   osmosisRemoteAddress,
		ChainId:         "osmosis-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "osmosisContractAddress1",
	}

	userInsOsmo2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   osmosisRemoteAddress,
		ChainId:         "osmosis-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "osmosisContractAddress2",
	}

	userInsMig1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   migalooRemoteAddress,
		ChainId:         "migaloo-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "migalooContractAddress1",
	}

	userInsMig2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   migalooRemoteAddress,
		ChainId:         "migaloo-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "migalooContractAddress2",
	}

	userInss := []types.MsgCreateUserInstruction{UerInsJuno1, userInsJuno2, userInsOsmo1, userInsOsmo2, userInsMig1, userInsMig2}

	deleteAll := types.MsgDeleteUserInstruction{
		LocalAddress: localAddress,
	}

	tests := []struct {
		name        string
		userIns     types.MsgDeleteUserInstruction
		userInsLeft int
	}{
		{
			name: "all",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress: localAddress,
			},
			userInsLeft: 0,
		}, {
			name: "all on juno remote address",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:  localAddress,
				RemoteAddress: junoRemoteAddress,
			},
			userInsLeft: 4,
		}, {
			name: "all on juno chain id",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress: localAddress,
				ChainId:      "juno-1",
			},
			userInsLeft: 4,
		}, {
			name: "all on juno contract 1 address",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:    localAddress,
				ContractAddress: junoCon1,
			},
			userInsLeft: 5,
		}, {
			name: "all on juno remote address and juno chain id",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:  localAddress,
				RemoteAddress: junoRemoteAddress,
				ChainId:       "juno-1",
			},
			userInsLeft: 4,
		}, {
			name: "all on juno remote address and juno contract address",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   junoRemoteAddress,
				ContractAddress: junoCon1,
			},
			userInsLeft: 5,
		}, {
			name: "all on juno remote address, juno contract address and juno chain id",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   junoRemoteAddress,
				ChainId:         "juno-1",
				ContractAddress: junoCon1,
			},
			userInsLeft: 5,
		},
		{
			name: "all on migaloo remote address",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress:  localAddress,
				RemoteAddress: migalooRemoteAddress,
			},
			userInsLeft: 4,
		}, {
			name: "all on migaloo chain id",
			userIns: types.MsgDeleteUserInstruction{
				LocalAddress: localAddress,
				ChainId:      "migaloo-1",
			},
			userInsLeft: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, userIns := range userInss {
				_, err := ms.CreateUserInstruction(ctx, &userIns)
				require.NoError(t, err)
			}

			_, err = ms.DeleteUserInstruction(ctx, &tt.userIns)
			require.NoError(t, err)

			useInsGet, found := k.GetUserInstructions(ctx, localAddress)

			if tt.name != "all" {
				require.True(t, found)
				require.Equal(t, tt.userInsLeft, len(useInsGet.GetUserInstruction()))

				_, err = ms.DeleteUserInstruction(ctx, &deleteAll)
				require.NoError(t, err)

				_, found := k.GetUserInstructions(ctx, localAddress)
				require.False(t, found)
			} else {
				require.False(t, found)
			}
		})
	}
}

func TestUpdateUserInstruction(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	ms := keeper.NewMsgServerImpl(*k)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotNil(t, k)

	localAddress := sample.AccAddress()
	junoRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "juno")
	require.NoError(t, err)
	osmosisRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "osmosis")
	require.NoError(t, err)
	migalooRemoteAddress, err := types.DeriveRemoteAddress(localAddress, "migaloo")
	require.NoError(t, err)

	junoCon1 := "junoContractAddress1"
	junoCon2 := "junoContractAddress2"

	UerInsJuno1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   junoRemoteAddress,
		ChainId:         "juno-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: junoCon1,
	}

	userInsJuno2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   junoRemoteAddress,
		ChainId:         "juno-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: junoCon2,
	}

	userInsOsmo1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   osmosisRemoteAddress,
		ChainId:         "osmosis-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "osmosisContractAddress1",
	}

	userInsOsmo2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   osmosisRemoteAddress,
		ChainId:         "osmosis-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "osmosisContractAddress2",
	}

	userInsMig1 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   migalooRemoteAddress,
		ChainId:         "migaloo-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "migalooContractAddress1",
	}

	userInsMig2 := types.MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   migalooRemoteAddress,
		ChainId:         "migaloo-1",
		Frequency:       3600,
		Expires:         time.Now().Add(time.Hour * 24),
		Instruction:     "{}",
		StrategyId:      1,
		ContractAddress: "migalooContractAddress2",
	}

	userInss := []types.MsgCreateUserInstruction{UerInsJuno1, userInsJuno2, userInsOsmo1, userInsOsmo2, userInsMig1, userInsMig2}

	for _, userIns := range userInss {
		_, err := ms.CreateUserInstruction(ctx, &userIns)
		require.NoError(t, err)
	}

	tests := []struct {
		name    string
		userIns types.MsgUpdateUserInstruction
		err     error
	}{
		{
			name: "invalid expiration",
			userIns: types.MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   junoRemoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Date(0000, 01, 01, 0, 0, 0, 0, time.UTC),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			userIns: types.MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   junoRemoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress3",
			},
		}, {
			name: "not found",
			userIns: types.MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   junoRemoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().UTC().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress4",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = ms.UpdateUserInstruction(ctx, &tt.userIns)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}
