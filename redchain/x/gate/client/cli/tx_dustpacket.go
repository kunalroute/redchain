package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v6/modules/core/04-channel/client/utils"
	"github.com/spf13/cobra"
	"redchain/x/gate/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdSendDustpacket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-dustpacket [src-port] [src-channel] [routeAmount] [routeRecipient] [destChainId] [requestMetadata] [requestPacket] [senderAddress]",
		Short: "Send a dustpacket over IBC",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			routeAmount := args[2]
			routeRecipient := args[3]
			destChainId := args[4]
			requestMetadata := args[5]
			requestPacket := args[6]
			senderAddress := args[7]

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			// TODO: check routeAmount size
			routeAmountInt, _ := github_com_cosmos_cosmos_sdk_types.NewIntFromString(routeAmount)

			msg := types.NewMsgSendDustpacket(creator, srcPort, srcChannel, timeoutTimestamp, routeAmountInt, []byte(routeRecipient), destChainId, requestMetadata, requestPacket, []byte(senderAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
