package uwalproto

type UWALVersion uint64

const (
	UWAL_MAGIC_V01   uint64      = 0x8c335e3aa738302b
	UWAL_VERSION_V01 UWALVersion = 0xdfd54931a9209394
)

type UWALHeader struct {
	WALMagic uint64      // Magic number to identify a valid UWAL file
	Version  UWALVersion // Version of the UWAL file format
	FileID   uint64      // File ID of the UWAL file
	UWALTS   uint64      // Timestamp of the UWAL file
}
