package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/temporal-zone/temporal/x/record/keeper"
	"github.com/temporal-zone/temporal/x/record/types"
)

func SimulateMsgCreateUserInstruction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateUserInstruction{
			LocalAddress: simAccount.Address.String(),
		}

		// TODO: Handling the CreateUserInstruction simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateUserInstruction simulation not implemented"), nil, nil
	}
}
