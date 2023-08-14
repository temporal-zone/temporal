package compound

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/compound/keeper"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the compoundSetting
	for _, elem := range genState.CompoundSettingList {
		k.SetCompoundSetting(ctx, elem)
	}
	// Set all the previousCompound
	for _, elem := range genState.PreviousCompoundList {
		k.SetPreviousCompound(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CompoundSettingList = k.GetAllCompoundSetting(ctx)
	genesis.PreviousCompoundList = k.GetAllPreviousCompound(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
