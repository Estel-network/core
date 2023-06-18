package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CollateralKeyPrefix is the prefix to retrieve all Collateral
	CollateralKeyPrefix = "Collateral/value/"
)

// CollateralKey returns the store key to retrieve a Collateral from the index fields
func CollateralKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
