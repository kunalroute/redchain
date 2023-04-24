package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"redchain/x/gate/types"
)

func (k msgServer) SendDustpacket(goCtx context.Context, msg *types.MsgSendDustpacket) (*types.MsgSendDustpacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.DustpacketPacketData

	packet.RouteAmount = msg.RouteAmount
	packet.RouteRecipient = msg.RouteRecipient
	packet.DestChainId = msg.DestChainId
	packet.RequestMetadata = msg.RequestMetadata
	packet.RequestPacket = msg.RequestPacket

	// Transmit the packet
	_, err := k.TransmitDustpacketPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendDustpacketResponse{}, nil
}
