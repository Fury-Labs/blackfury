// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fury-labs/blackfury/v13/x/recovery/types"
)

var _ types.QueryServer = Keeper{}

// Params returns the module parameters
func (k Keeper) Params(
	c context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
