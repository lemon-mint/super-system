package uwalproto

func (ns25519 *UWALHeader) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{WALMagic uint64; Version ./lemondb/uwal/uwalproto.UWALVersion; FileID uint64; UWALTS uint64})(ns25519)

	// ZZ: (uint64)(ns25519.WALMagic)
	ns25520 += 8

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(ns25519.Version)

	// ZZ: (uint64)(ns25519.Version)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.FileID)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.UWALTS)
	ns25520 += 8

	return ns25520
}

func (ns25521 *UWALHeader) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{WALMagic uint64; Version ./lemondb/uwal/uwalproto.UWALVersion; FileID uint64; UWALTS uint64})(ns25521)

	// ZZ: (uint64)(ns25521.WALMagic)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.WALMagic >> 0)
	dst[ns25522+1] = byte(ns25521.WALMagic >> 8)
	dst[ns25522+2] = byte(ns25521.WALMagic >> 16)
	dst[ns25522+3] = byte(ns25521.WALMagic >> 24)
	dst[ns25522+4] = byte(ns25521.WALMagic >> 32)
	dst[ns25522+5] = byte(ns25521.WALMagic >> 40)
	dst[ns25522+6] = byte(ns25521.WALMagic >> 48)
	dst[ns25522+7] = byte(ns25521.WALMagic >> 56)
	ns25522 += 8

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(ns25521.Version)

	// ZZ: (uint64)(ns25521.Version)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.Version >> 0)
	dst[ns25522+1] = byte(ns25521.Version >> 8)
	dst[ns25522+2] = byte(ns25521.Version >> 16)
	dst[ns25522+3] = byte(ns25521.Version >> 24)
	dst[ns25522+4] = byte(ns25521.Version >> 32)
	dst[ns25522+5] = byte(ns25521.Version >> 40)
	dst[ns25522+6] = byte(ns25521.Version >> 48)
	dst[ns25522+7] = byte(ns25521.Version >> 56)
	ns25522 += 8

	// ZZ: (uint64)(ns25521.FileID)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.FileID >> 0)
	dst[ns25522+1] = byte(ns25521.FileID >> 8)
	dst[ns25522+2] = byte(ns25521.FileID >> 16)
	dst[ns25522+3] = byte(ns25521.FileID >> 24)
	dst[ns25522+4] = byte(ns25521.FileID >> 32)
	dst[ns25522+5] = byte(ns25521.FileID >> 40)
	dst[ns25522+6] = byte(ns25521.FileID >> 48)
	dst[ns25522+7] = byte(ns25521.FileID >> 56)
	ns25522 += 8

	// ZZ: (uint64)(ns25521.UWALTS)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.UWALTS >> 0)
	dst[ns25522+1] = byte(ns25521.UWALTS >> 8)
	dst[ns25522+2] = byte(ns25521.UWALTS >> 16)
	dst[ns25522+3] = byte(ns25521.UWALTS >> 24)
	dst[ns25522+4] = byte(ns25521.UWALTS >> 32)
	dst[ns25522+5] = byte(ns25521.UWALTS >> 40)
	dst[ns25522+6] = byte(ns25521.UWALTS >> 48)
	dst[ns25522+7] = byte(ns25521.UWALTS >> 56)
	ns25522 += 8

	return ns25522
}

func (ns25523 *UWALHeader) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{WALMagic uint64; Version ./lemondb/uwal/uwalproto.UWALVersion; FileID uint64; UWALTS uint64})(ns25523)

	// ZZ: (uint64)(ns25523.WALMagic)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.WALMagic = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(ns25523.Version)

	// ZZ: (uint64)(ns25523.Version)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.Version = UWALVersion(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.FileID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.FileID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.UWALTS)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.UWALTS = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25524 *UWALVersion) SizeGOBE() uint64 {
	var ns25525 uint64

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(*ns25524)

	// ZZ: (uint64)(*ns25524)
	ns25525 += 8

	return ns25525
}

func (ns25526 *UWALVersion) MarshalGOBE(dst []byte) uint64 {
	var ns25527 uint64

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(*ns25526)

	// ZZ: (uint64)(*ns25526)
	_ = dst[ns25527+7]
	dst[ns25527+0] = byte(*ns25526 >> 0)
	dst[ns25527+1] = byte(*ns25526 >> 8)
	dst[ns25527+2] = byte(*ns25526 >> 16)
	dst[ns25527+3] = byte(*ns25526 >> 24)
	dst[ns25527+4] = byte(*ns25526 >> 32)
	dst[ns25527+5] = byte(*ns25526 >> 40)
	dst[ns25527+6] = byte(*ns25526 >> 48)
	dst[ns25527+7] = byte(*ns25526 >> 56)
	ns25527 += 8

	return ns25527
}

func (ns25528 *UWALVersion) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (./lemondb/uwal/uwalproto.UWALVersion)(*ns25528)

	// ZZ: (uint64)(*ns25528)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	*ns25528 = UWALVersion(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}
