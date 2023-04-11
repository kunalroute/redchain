package keeper

import (
	"redchain/x/redchain/types"
)

var _ types.QueryServer = Keeper{}
