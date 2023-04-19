package cli

import (
	"math/rand"
	"time"
)

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
