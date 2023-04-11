package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "redchain/testutil/keeper"
	"redchain/x/gate/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
