package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PreviousCompoundingKeyPrefix is the prefix to retrieve all PreviousCompounding
	PreviousCompoundingKeyPrefix = "PreviousCompounding/value/"
)

// PreviousCompoundingKey returns the store key to retrieve a PreviousCompounding from the index fields
func PreviousCompoundingKey(
	delegator string,
) []byte {
	var key []byte

	delegatorBytes := []byte(delegator)
	key = append(key, delegatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
