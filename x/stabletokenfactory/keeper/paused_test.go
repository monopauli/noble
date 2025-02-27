package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/strangelove-ventures/noble/v5/testutil/keeper"
	"github.com/strangelove-ventures/noble/v5/testutil/nullify"
	"github.com/strangelove-ventures/noble/v5/x/stabletokenfactory/keeper"
	"github.com/strangelove-ventures/noble/v5/x/stabletokenfactory/types"
)

func createTestPaused(keeper *keeper.Keeper, ctx sdk.Context) types.Paused {
	item := types.Paused{}
	keeper.SetPaused(ctx, item)
	return item
}

func TestPausedGet(t *testing.T) {
	keeper, ctx := keepertest.StableTokenFactoryKeeper(t)
	item := createTestPaused(keeper, ctx)
	rst := keeper.GetPaused(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
