package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DelegationHistoryKeyPrefix is the prefix to retrieve all DelegationHistory
	DelegationHistoryKeyPrefix = "DelegationHistory/value/"
)

// DelegationHistoryKey returns the store key to retrieve a DelegationHistory from the index fields
func DelegationHistoryKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
