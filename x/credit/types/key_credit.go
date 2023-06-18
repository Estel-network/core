package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CreditKeyPrefix is the prefix to retrieve all Credit
	CreditKeyPrefix = "Credit/value/"
)

// CreditKey returns the store key to retrieve a Credit from the index fields
func CreditKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
