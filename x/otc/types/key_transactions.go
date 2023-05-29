package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TransactionsKeyPrefix is the prefix to retrieve all Transactions
	TransactionsKeyPrefix = "Transactions/value/"
)

// TransactionsKey returns the store key to retrieve a Transactions from the index fields
func TransactionsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
