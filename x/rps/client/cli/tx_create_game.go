package cli

import (
	"encoding/hex"
	"errors"
	"strconv"

	"rps/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/crypto/sha3"
)

var _ = strconv.Itoa(0)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [bet-amount] [turn]",
		Short: "Broadcast message create-game with turn",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBetAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argTurn := args[1]
			if !ValidateTurn(argTurn) {
				return errors.New("invalid turn. Turn must be rock, scissors or paper")
			}
			salt := GenerateSalt()
			// print salt
			cmd.Println("Remember your Salt: ", salt)
			turnHashBytes := sha3.Sum224([]byte(argTurn + salt))
			turnHash := hex.EncodeToString(turnHashBytes[:])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(
				clientCtx.GetFromAddress().String(),
				argBetAmount,
				turnHash,
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
