package mskip

import (
	"encoding/binary"

	"v8.run/go/supersystem/lemondb/uarena"
)

const MAX_HEIGHT = 32
const TIMESTAMP_SIZE = 8

type Node struct {
	Height uint16
	Key    uint64
	Value  uint64

	Next [MAX_HEIGHT]uint32
}

func loadNode(arena *uarena.UArena, addr uint32, node *Node) {
	_ = arena.Data[addr+17] // bounds check hint to compiler; see golang.org/issue/14808
	node.Height = uint16(arena.Data[addr]) | uint16(arena.Data[addr+1])<<8
	node.Key = uint64(arena.Data[addr+2]) | uint64(arena.Data[addr+3])<<8 |
		uint64(arena.Data[addr+4])<<16 | uint64(arena.Data[addr+5])<<24 |
		uint64(arena.Data[addr+6])<<32 | uint64(arena.Data[addr+7])<<40 |
		uint64(arena.Data[addr+8])<<48 | uint64(arena.Data[addr+9])<<56
	node.Value = uint64(arena.Data[addr+10]) | uint64(arena.Data[addr+11])<<8 |
		uint64(arena.Data[addr+12])<<16 | uint64(arena.Data[addr+13])<<24 |
		uint64(arena.Data[addr+14])<<32 | uint64(arena.Data[addr+15])<<40 |
		uint64(arena.Data[addr+16])<<48 | uint64(arena.Data[addr+17])<<56
	for i := uint32(0); i < uint32(node.Height); i++ {
		_ = arena.Data[addr+19+2*i] // bounds check hint to compiler; see golang.org/issue/14808
		node.Next[i] = uint32(arena.Data[addr+18+2*i]) | uint32(arena.Data[addr+19+2*i])<<8
	}
}

func sizeNode(node *Node) uint32 {
	return 18 + 2*uint32(node.Height)
}

func storeNode(arena *uarena.UArena, addr uint32, node *Node) {
	arena.Data[addr] = byte(node.Height)
	arena.Data[addr+1] = byte(node.Height >> 8)
	arena.Data[addr+2] = byte(node.Key)
	arena.Data[addr+3] = byte(node.Key >> 8)
	arena.Data[addr+4] = byte(node.Key >> 16)
	arena.Data[addr+5] = byte(node.Key >> 24)
	arena.Data[addr+6] = byte(node.Key >> 32)
	arena.Data[addr+7] = byte(node.Key >> 40)
	arena.Data[addr+8] = byte(node.Key >> 48)
	arena.Data[addr+9] = byte(node.Key >> 56)
	arena.Data[addr+10] = byte(node.Value)
	arena.Data[addr+11] = byte(node.Value >> 8)
	arena.Data[addr+12] = byte(node.Value >> 16)
	arena.Data[addr+13] = byte(node.Value >> 24)
	arena.Data[addr+14] = byte(node.Value >> 32)
	arena.Data[addr+15] = byte(node.Value >> 40)
	arena.Data[addr+16] = byte(node.Value >> 48)
	arena.Data[addr+17] = byte(node.Value >> 56)
	for i := uint32(0); i < uint32(node.Height); i++ {
		arena.Data[addr+18+2*i] = byte(node.Next[i])
		arena.Data[addr+19+2*i] = byte(node.Next[i] >> 8)
	}
}

type nodeAddr uint32

func (addr nodeAddr) Offset() uint32 {
	return uint32(addr)
}

func (addr nodeAddr) Height(arena *uarena.UArena) uint16 {
	offset := addr.Offset()
	_ = arena.Data[offset+1]
	return uint16(arena.Data[offset]) | uint16(arena.Data[offset+1])<<8
}

func (addr nodeAddr) Key(arena *uarena.UArena) uint64 {
	offset := addr.Offset()
	_ = arena.Data[offset+9]
	return uint64(arena.Data[offset+2]) | uint64(arena.Data[offset+3])<<8 |
		uint64(arena.Data[offset+4])<<16 | uint64(arena.Data[offset+5])<<24 |
		uint64(arena.Data[offset+6])<<32 | uint64(arena.Data[offset+7])<<40 |
		uint64(arena.Data[offset+8])<<48 | uint64(arena.Data[offset+9])<<56
}

func (addr nodeAddr) KeyBytes(arena *uarena.UArena) []byte {
	key := addr.Key(arena)
	offset := uint32(key >> 32)
	length := uint32(key)
	return arena.Data[offset : offset+length]
}

func (addr nodeAddr) Value(arena *uarena.UArena) uint64 {
	offset := addr.Offset()
	_ = arena.Data[offset+17]
	return uint64(arena.Data[offset+10]) | uint64(arena.Data[offset+11])<<8 |
		uint64(arena.Data[offset+12])<<16 | uint64(arena.Data[offset+13])<<24 |
		uint64(arena.Data[offset+14])<<32 | uint64(arena.Data[offset+15])<<40 |
		uint64(arena.Data[offset+16])<<48 | uint64(arena.Data[offset+17])<<56
}

func (addr nodeAddr) ValueBytes(arena *uarena.UArena) []byte {
	value := addr.Value(arena)
	offset := uint32(value >> 32)
	length := uint32(value)
	return arena.Data[offset : offset+length]
}

func (addr nodeAddr) Next(arena *uarena.UArena, i uint32) nodeAddr {
	offset := addr.Offset()
	_ = arena.Data[offset+19+2*i]
	return nodeAddr(uint32(arena.Data[offset+18+2*i]) | uint32(arena.Data[offset+19+2*i])<<8)
}

func mkkey(arena *uarena.UArena, key []byte, ts uint64) uint32 {
	size := uint32(len(key) + TIMESTAMP_SIZE)
	addr := arena.Alloc(size)
	if addr == uarena.InvalidOffset {
		return uarena.InvalidOffset
	}
	copy(arena.Data[addr:addr+size], key)
	binary.LittleEndian.PutUint64(arena.Data[addr+size-TIMESTAMP_SIZE:], ts)
	return addr
}

type SkipList struct {
	Arena *uarena.UArena
	Head  nodeAddr
}

func NewSkipList(arena *uarena.UArena) *SkipList {
	sk := &SkipList{Arena: arena, Head: 0}

	head := Node{Height: MAX_HEIGHT}
	size := sizeNode(&head)
	addr := arena.Alloc(size)
	if addr == uarena.InvalidOffset {
		return nil
	}
	storeNode(arena, addr, &head)
	sk.Head = nodeAddr(addr)

	return sk
}

func (sk *SkipList) seek(key []byte, ts uint64) (nodeAddr, bool) {
	x := sk.Head
	for i := uint32(MAX_HEIGHT - 1); i >= 0; i-- {
		for {
			next := x.Next(sk.Arena, i)
			if next == 0 {
				// Tail
				break
			}

			keyB := next.KeyBytes(sk.Arena)
			keyData := keyB[:len(keyB)-TIMESTAMP_SIZE]
			tsData := binary.LittleEndian.Uint64(keyB[len(keyB)-TIMESTAMP_SIZE:])
			if string(keyData) > string(key) {
				break
			}
			if string(keyData) == string(key) && tsData >= ts {
				break
			}

			// Skip to next node
			x = next
		}
	}
	x = x.Next(sk.Arena, 0)
	if x == 0 {
		return 0, false
	}

	return x, true
}

type ctx [MAX_HEIGHT]nodeAddr

func (sk *SkipList) seekContext(c *ctx, key []byte, ts uint64) bool {
	x := sk.Head
	for i := uint32(MAX_HEIGHT - 1); i >= 0; i-- {
		for {
			next := x.Next(sk.Arena, i)
			if next == 0 {
				// Tail
				break
			}

			keyB := next.KeyBytes(sk.Arena)
			keyData := keyB[:len(keyB)-TIMESTAMP_SIZE]
			tsData := binary.LittleEndian.Uint64(keyB[len(keyB)-TIMESTAMP_SIZE:])
			if string(keyData) > string(key) {
				break
			}
			if string(keyData) == string(key) && tsData >= ts {
				break
			}

			// Skip to next node
			x = next
		}
		c[i] = x
	}
	x = x.Next(sk.Arena, 0)
	if x == 0 {
		return false
	}

	return true
}
