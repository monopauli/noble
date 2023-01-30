package ibctest_test

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/ibctest/v3"
	"github.com/strangelove-ventures/ibctest/v3/chain/cosmos"
	"github.com/strangelove-ventures/ibctest/v3/ibc"
	"github.com/strangelove-ventures/ibctest/v3/testreporter"
	integration "github.com/strangelove-ventures/noble/ibctest"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// From test_genesis.json file:
const (
	blacklisterMnemonic  = "machine cross board brass farm hamster betray airport need strong patient awkward resemble rent script online reform have home army wood melt cash cage"
	masterMinterMnemonic = "assist quiz car guide dog cash pill service tell shiver nation dune echo inspire fox corn photo six between sudden yellow identify citizen exercise"
	ownerMnemonic        = "goat amateur absorb raven ankle will basic mean edit garbage analyst bacon dial avoid belt aspect puzzle brave umbrella ketchup material region virtual exchange"
	pauserMnemonic       = "donate upon pioneer script achieve carry confirm slogan blame glance buddy mango borrow cash brass average two never tide bonus december member remember violin"

	authorityMnemonic = "fashion race whisper upgrade primary eternal put february sport social before panel sign attack tilt milk now diamond honey unit deny cabin kind budget"
)

func TestNobleGenStart(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctx := context.Background()

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	client, network := ibctest.DockerSetup(t)

	repo, version := integration.GetDockerImageInfo()

	var noble *cosmos.CosmosChain
	var roles NobleRoles

	chainCfg := ibc.ChainConfig{
		Type:           "cosmos",
		Name:           "noble",
		ChainID:        "noble-1",
		Bin:            "nobled",
		Denom:          "token",
		Bech32Prefix:   "noble",
		CoinType:       "118",
		GasPrices:      "0.0token",
		GasAdjustment:  1.1,
		TrustingPeriod: "504h",
		NoHostMount:    false,
		Images: []ibc.DockerImage{
			{
				Repository: repo,
				Version:    version,
				UidGid:     "1025:1025",
			},
		},
		EncodingConfig:  NobleEncoding(),
		GenesisFilePath: "/Users/dkwork/go/src/github.com/strangelove-ventures/noble/ibctest/test_genesis.json",
	}

	nv := 1
	nf := 0

	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{
			ChainConfig:   chainCfg,
			NumValidators: &nv,
			NumFullNodes:  &nf,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	noble = chains[0].(*cosmos.CosmosChain)

	ic := ibctest.NewInterchain().
		AddChain(noble)

	require.NoError(t, ic.Build(ctx, eRep, ibctest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: false,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	nobleValidator := noble.Validators[0]

	err = nobleValidator.RecoverKey(ctx, ownerKeyName, ownerMnemonic)
	require.NoError(t, err, "failed to recover owner key")

	err = nobleValidator.RecoverKey(ctx, masterMinterKeyName, masterMinterMnemonic)
	require.NoError(t, err, "failed to recover master minter key")

	_, err = nobleValidator.ExecTx(ctx, masterMinterKeyName,
		"tokenfactory", "configure-minter-controller", roles.MinterController.Address, roles.Minter.Address,
	)
	require.NoError(t, err, "failed to execute configure minter controller tx")

	_, err = nobleValidator.ExecTx(ctx, minterControllerKeyName,
		"tokenfactory", "configure-minter", roles.Minter.Address, "1000uusdc",
	)
	require.NoError(t, err, "failed to execute configure minter tx")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.User.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	userBalance, err := noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(100), userBalance, "failed to mint uusdc to user")

	_, err = nobleValidator.ExecTx(ctx, ownerKeyName,
		"tokenfactory", "update-blacklister", roles.Blacklister.Address,
	)
	require.NoError(t, err, "failed to set blacklister")

	_, err = nobleValidator.ExecTx(ctx, blacklisterKeyName,
		"tokenfactory", "blacklist", roles.User.Address,
	)
	require.NoError(t, err, "failed to blacklist user address")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.User.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(100), userBalance, "user balance should not have incremented while blacklisted")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.User2.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user2 tx")

	err = nobleValidator.SendFunds(ctx, user2KeyName, ibc.WalletAmount{
		Address: roles.User.Address,
		Denom:   "uusdc",
		Amount:  50,
	})
	require.Error(t, err, "The tx to a blacklisted user should not have been successful")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(100), userBalance, "user balance should not have incremented while blacklisted")

	err = nobleValidator.SendFunds(ctx, user2KeyName, ibc.WalletAmount{
		Address: roles.User.Address,
		Denom:   "token",
		Amount:  100,
	})
	require.NoError(t, err, "The tx should have been successfull as that is no the minting denom")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "token")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(10_100), userBalance, "user balance should have incremented")

	_, err = nobleValidator.ExecTx(ctx, blacklisterKeyName,
		"tokenfactory", "unblacklist", roles.User.Address,
	)
	require.NoError(t, err, "failed to unblacklist user address")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.User.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(200), userBalance, "user balance should have increased now that they are no longer blacklisted")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.Minter.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	minterBalance, err := noble.GetBalance(ctx, roles.Minter.Address, "uusdc")
	require.NoError(t, err, "failed to get minter balance")
	require.Equal(t, int64(100), minterBalance, "minter balance should have increased")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName, "tokenfactory", "burn", "10uusdc")
	require.NoError(t, err, "failed to execute burn tx")

	minterBalance, err = noble.GetBalance(ctx, roles.Minter.Address, "uusdc")
	require.NoError(t, err, "failed to get minter balance")
	require.Equal(t, int64(90), minterBalance, "minter balance should have decreased because tokens were burned")

	_, err = nobleValidator.ExecTx(ctx, ownerKeyName,
		"tokenfactory", "update-pauser", roles.Pauser.Address,
	)
	require.NoError(t, err, "failed to update pauser")

	_, err = nobleValidator.ExecTx(ctx, pauserKeyName,
		"tokenfactory", "pause",
	)
	require.NoError(t, err, "failed to pause mints")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName,
		"tokenfactory", "mint", roles.User.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")

	require.Equal(t, int64(200), userBalance, "user balance should not have increased while chain is paused")

	_, err = nobleValidator.ExecTx(ctx, userKeyName,
		"bank", "send", roles.User.Address, roles.Alice.Address, "100uusdc",
	)
	require.Error(t, err, "transaction was successful while chain was paused")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")

	require.Equal(t, int64(200), userBalance, "user balance should not have changed while chain is paused")

	aliceBalance, err := noble.GetBalance(ctx, roles.Alice.Address, "uusdc")
	require.NoError(t, err, "failed to get alice balance")

	require.Equal(t, int64(0), aliceBalance, "alice balance should not have increased while chain is paused")

	_, err = nobleValidator.ExecTx(ctx, minterKeyName, "tokenfactory", "burn", "10uusdc")
	require.NoError(t, err, "failed to execute burn tx")
	require.Equal(t, int64(90), minterBalance, "this burn should not have been successful because the chain is paused")

	_, err = nobleValidator.ExecTx(ctx, masterMinterKeyName,
		"tokenfactory", "configure-minter-controller", roles.MinterController.Address, roles.User.Address,
	)
	require.NoError(t, err, "failed to execute configure minter controller tx")

	_, _, err = nobleValidator.ExecQuery(ctx, "tokenfactory", "show-minters", roles.User.Address)
	require.Error(t, err, "'user' should not have been able to become a minter while chain is paused")

	_, err = nobleValidator.ExecTx(ctx, minterControllerKeyName, "tokenfactory", "remove-minter", roles.Minter.Address)
	require.NoError(t, err, "failed to send remove minter tx")

	_, _, err = nobleValidator.ExecQuery(ctx, "tokenfactory", "show-minters", roles.Minter.Address)
	require.Error(t, err, "minter should have been removed, even while chain is puased")

	_, err = nobleValidator.ExecTx(ctx, pauserKeyName,
		"tokenfactory", "unpause",
	)
	require.NoError(t, err, "failed to unpause mints")

	_, err = nobleValidator.ExecTx(ctx, userKeyName,
		"bank", "send", roles.User.Address, roles.Alice.Address, "100uusdc",
	)
	require.NoError(t, err, "failed to send tx bank from user to alice")

	userBalance, err = noble.GetBalance(ctx, roles.User.Address, "uusdc")
	require.NoError(t, err, "failed to get user balance")
	require.Equal(t, int64(100), userBalance, "user balance should not have changed while chain is paused")

	aliceBalance, err = noble.GetBalance(ctx, roles.Alice.Address, "uusdc")
	require.NoError(t, err, "failed to get alice balance")
	require.Equal(t, int64(100), aliceBalance, "alice balance should not have increased while chain is paused")

}
