// Code generated by "gobe ."; DO NOT EDIT.
// versions:
//     gobe: v0.2.1

package errno

func (ns25519 *Errno) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (..Errno)(*ns25519)

	// ZZ: (uint16)(*ns25519)
	ns25520 += 2

	return ns25520
}

func (ns25521 *Errno) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (..Errno)(*ns25521)

	// ZZ: (uint16)(*ns25521)
	_ = dst[ns25522+1]
	dst[ns25522+0] = byte(*ns25521 >> 0)
	dst[ns25522+1] = byte(*ns25521 >> 8)
	ns25522 += 2

	return ns25522
}

func (ns25523 *Errno) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..Errno)(*ns25523)

	// ZZ: (uint16)(*ns25523)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	*ns25523 = Errno(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
