package keeper

import (
	"testing"

	"redchain/x/gate/keeper"
	"redchain/x/gate/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// gateChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type gateChannelKeeper struct{}

func (gateChannelKeeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	return channeltypes.Channel{}, false
}

func (gateChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}

func (gateChannelKeeper) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	return 0, nil
}

func (gateChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// gateportKeeper is a stub of cosmosibckeeper.PortKeeper
type gatePortKeeper struct{}

func (gatePortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

func GateKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"GateParams",
	)
	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		gateChannelKeeper{},
		gatePortKeeper{},
		capabilityKeeper.ScopeToModule("GateScopedKeeper"),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
