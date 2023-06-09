package types

const (
	// ModuleName defines the module name
	ModuleName = "icayieldmos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_icayieldmos"

	// Version defines the current version the IBC module supports
	Version = "icayieldmos-1"

	// PortID is the default port id that module binds to
	PortID = "icayieldmos"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("icayieldmos-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ContractRemoteZoneKey      = "ContractRemoteZone/value/"
	ContractRemoteZoneCountKey = "ContractRemoteZone/count/"
)

const (
	RemoteContractCompoundSettingsKey= "RemoteContractCompoundSettings/value/"
	RemoteContractCompoundSettingsCountKey= "RemoteContractCompoundSettings/count/"
)
