package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "redchain/testutil/keeper"
	"redchain/x/redchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RedchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
