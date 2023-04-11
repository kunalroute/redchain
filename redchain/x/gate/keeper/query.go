package keeper

import (
	"redchain/x/gate/types"
)

var _ types.QueryServer = Keeper{}
