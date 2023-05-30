package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"igmf/x/otc/types"
)

var _ = strconv.Itoa(0)

func CmdCreateTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-transaction [buyer] [denom-sell-amount] [denom-buy-amount]",
		Short: "Broadcast message create-transaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBuyer := args[0]
			argDenomSellAmount := args[1]
			argDenomBuyAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTransaction(
				clientCtx.GetFromAddress().String(),
				argBuyer,
				argDenomSellAmount,
				argDenomBuyAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
