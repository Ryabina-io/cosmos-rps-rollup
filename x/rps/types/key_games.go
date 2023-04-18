package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GamesKeyPrefix is the prefix to retrieve all Games
	GamesKeyPrefix = "Games/value/"
)

// GamesKey returns the store key to retrieve a Games from the index fields
func GamesKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
