package cli

import (
	"errors"
	"strconv"

	"rps/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRevealGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-game [game-id] [turn] [salt]",
		Short: "Broadcast message reveal-game",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTurn := args[1]
			argSalt := args[2]

			if !ValidateTurn(argTurn) {
				return errors.New("invalid turn. Turn must be rock, scissors or paper")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevealGame(
				clientCtx.GetFromAddress().String(),
				argGameId,
				argTurn,
				argSalt,
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
