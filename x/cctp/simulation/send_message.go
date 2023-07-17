package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/strangelove-ventures/noble/x/cctp/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func SimulateMsgSendMessage(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	fk types.FiatTokenfactoryKeeper,
	k *keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSendMessage{
			From: simAccount.Address.String(),
		}

		// TODO: Handling the SendMessage simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendMessage simulation not implemented"), nil, nil
	}
}
