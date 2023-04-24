package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgSendDustpacket = "send_dustpacket"

var _ sdk.Msg = &MsgSendDustpacket{}

func NewMsgSendDustpacket(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	routerAmount github_com_cosmos_cosmos_sdk_types.Int,
	routerRecepient []byte,
	destChainId string,
	requestMetadata string,
	requestPacket string,
	senderAddress []byte,
) *MsgSendDustpacket {
	return &MsgSendDustpacket{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		RouteAmount:   routerAmount,
		RouteRecipient: routerRecepient,
		DestChainId:    destChainId,
		RequestMetadata: requestMetadata, 
		RequestPacket:  requestPacket,
		SenderAddress:  senderAddress,
	}
}

func (msg *MsgSendDustpacket) Route() string {
	return RouterKey
}

func (msg *MsgSendDustpacket) Type() string {
	return TypeMsgSendDustpacket
}

func (msg *MsgSendDustpacket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendDustpacket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendDustpacket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
