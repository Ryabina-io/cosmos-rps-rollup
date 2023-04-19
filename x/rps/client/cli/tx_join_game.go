package cli

import (
	"encoding/hex"
	"errors"
	"strconv"

	"rps/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/crypto/sha3"
)

var _ = strconv.Itoa(0)

func CmdJoinGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-game [game-id] [turn]",
		Short: "Broadcast message join-game for a game with turn",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameId, err := cast.ToUint64E(args[0])
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
			// TODO add validation for game before send transaction

			msg := types.NewMsgJoinGame(
				clientCtx.GetFromAddress().String(),
				argGameId,
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
