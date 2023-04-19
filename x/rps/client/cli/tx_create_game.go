package cli

import (
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"rps/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/crypto/sha3"
)

var _ = strconv.Itoa(0)

func init() {
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// function generates random salt string
func GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	return RandStringRunes(10)
}

// function validates rock scissors paper turn string
func ValidateTurn(turn string) bool {
	if turn == "rock" || turn == "scissors" || turn == "paper" {
		return true
	}
	return false
}

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
			if ValidateTurn(argTurn) == false {
				return errors.New("Invalid turn. Turn must be rock, scissors or paper")
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
