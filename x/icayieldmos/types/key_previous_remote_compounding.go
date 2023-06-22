package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PreviousRemoteCompoundingKeyPrefix is the prefix to retrieve all PreviousRemoteCompounding
	PreviousRemoteCompoundingKeyPrefix = "PreviousRemoteCompounding/value/"
)

// PreviousRemoteCompoundingKey returns the store key to retrieve a PreviousRemoteCompounding from the index fields
func PreviousRemoteCompoundingKey(
	remoteContractCompoundSetting uint64,
) []byte {
	var key []byte

	remoteContractCompoundSettingBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(remoteContractCompoundSettingBytes, remoteContractCompoundSetting)
	key = append(key, remoteContractCompoundSettingBytes...)
	key = append(key, []byte("/")...)

	return key
}
