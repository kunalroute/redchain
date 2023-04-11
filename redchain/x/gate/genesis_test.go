package gate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "redchain/testutil/keeper"
	"redchain/testutil/nullify"
	"redchain/x/gate"
	"redchain/x/gate/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GateKeeper(t)
	gate.InitGenesis(ctx, *k, genesisState)
	got := gate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
