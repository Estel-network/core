package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PositionsKeyPrefix is the prefix to retrieve all Positions
	PositionsKeyPrefix = "Positions/value/"
)

// PositionsKey returns the store key to retrieve a Positions from the index fields
func PositionsKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
