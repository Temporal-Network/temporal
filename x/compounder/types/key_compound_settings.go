package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CompoundSettingsKeyPrefix is the prefix to retrieve all CompoundSettings
	CompoundSettingsKeyPrefix = "CompoundSettings/value/"
)

// CompoundSettingsKey returns the store key to retrieve a CompoundSettings from the index fields
func CompoundSettingsKey(
	index123 string,
) []byte {
	var key []byte

	index123Bytes := []byte(index123)
	key = append(key, index123Bytes...)
	key = append(key, []byte("/")...)

	return key
}
