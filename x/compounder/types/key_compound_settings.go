package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CompoundSettingsKeyPrefix is the prefix to retrieve all CompoundSettings
	CompoundSettingsKeyPrefix = "CompoundSettings/value/"
)

// CompoundSettingsKey returns the store key to retrieve a CompoundSettings from the index fields
func CompoundSettingsKey(
	delegator string,
) []byte {
	var key []byte

	delegatorBytes := []byte(delegator)
	key = append(key, delegatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
