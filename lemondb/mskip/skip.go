package mskip

import (
	"encoding/binary"
	"time"

	"v8.run/go/supersystem/lemondb/uarena"
	"v8.run/go/supersystem/lemondb/wyhash/splitmix64"
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
		_ = arena.Data[addr+21+4*i] // bounds check hint to compiler; see golang.org/issue/14808
		node.Next[i] = uint32(arena.Data[addr+18+4*i]) |
			uint32(arena.Data[addr+19+4*i])<<8 |
			uint32(arena.Data[addr+20+4*i])<<16 |
			uint32(arena.Data[addr+21+4*i])<<24
	}
}

func sizeNode(node *Node) uint32 {
	return 18 + 4*uint32(node.Height)
}

func storeNode(arena *uarena.UArena, addr uint32, node *Node) {
	_ = arena.Data[addr+17] // bounds check hint to compiler; see golang.org/issue/14808
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
		_ = arena.Data[addr+21+2*i] // bounds check hint to compiler; see golang.org/issue/14808
		arena.Data[addr+18+2*i] = byte(node.Next[i])
		arena.Data[addr+19+2*i] = byte(node.Next[i] >> 8)
		arena.Data[addr+20+2*i] = byte(node.Next[i] >> 16)
		arena.Data[addr+21+2*i] = byte(node.Next[i] >> 24)
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

func (addr nodeAddr) SetValue(arena *uarena.UArena, value uint64) {
	offset := addr.Offset()
	_ = arena.Data[offset+17]
	arena.Data[offset+10] = byte(value)
	arena.Data[offset+11] = byte(value >> 8)
	arena.Data[offset+12] = byte(value >> 16)
	arena.Data[offset+13] = byte(value >> 24)
	arena.Data[offset+14] = byte(value >> 32)
	arena.Data[offset+15] = byte(value >> 40)
	arena.Data[offset+16] = byte(value >> 48)
	arena.Data[offset+17] = byte(value >> 56)
}

func (addr nodeAddr) ValueBytes(arena *uarena.UArena) []byte {
	value := addr.Value(arena)
	offset := uint32(value >> 32)
	length := uint32(value)
	return arena.Data[offset : offset+length]
}

func (addr nodeAddr) Next(arena *uarena.UArena, level uint32) nodeAddr {
	offset := addr.Offset()
	_ = arena.Data[offset+18+2*level]
	return nodeAddr(uint32(arena.Data[offset+18+2*level]) |
		uint32(arena.Data[offset+19+2*level])<<8 |
		uint32(arena.Data[offset+20+2*level])<<16 |
		uint32(arena.Data[offset+21+2*level])<<24)
}

func (addr nodeAddr) SetNext(arena *uarena.UArena, level uint32, next nodeAddr) {
	offset := addr.Offset()
	_ = arena.Data[offset+18+2*level]
	arena.Data[offset+18+2*level] = byte(next)
	arena.Data[offset+19+2*level] = byte(next >> 8)
	arena.Data[offset+20+2*level] = byte(next >> 16)
	arena.Data[offset+21+2*level] = byte(next >> 24)
}

const InvalidOffset = uint64(1<<64 - 1)

func mkkey(arena *uarena.UArena, key []byte, ts uint64) uint64 {
	size := uint32(len(key) + TIMESTAMP_SIZE)
	addr := arena.Alloc(size)
	if addr == uarena.InvalidOffset {
		return InvalidOffset
	}
	copy(arena.Data[addr:addr+size], key)
	binary.LittleEndian.PutUint64(arena.Data[addr+size-TIMESTAMP_SIZE:], ts)
	return uint64(addr)<<32 | uint64(size)
}

type SkipList struct {
	Arena    *uarena.UArena
	Head     nodeAddr
	RandSeed uint64
}

func NewSkipList(arena *uarena.UArena) *SkipList {
	sk := &SkipList{Arena: arena, Head: 0, RandSeed: uint64(time.Now().UnixNano()) | splitmix64.Next()}
	splitmix64.Splitmix64(&sk.RandSeed)
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

func (sk *SkipList) seekContext(c *ctx, key []byte, ts uint64) (nodeAddr, bool) {
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
		return 0, false
	}

	return x, true
}

func (sk *SkipList) Get(key []byte, ts uint64) []byte {
	x, ok := sk.seek(key, ts)
	if !ok {
		return nil
	}

	keyB := x.KeyBytes(sk.Arena)
	keyData := keyB[:len(keyB)-TIMESTAMP_SIZE]
	tsData := binary.LittleEndian.Uint64(keyB[len(keyB)-TIMESTAMP_SIZE:])
	if string(keyData) == string(key) && tsData <= ts {
		return x.ValueBytes(sk.Arena)
	}

	return nil
}

func randHeight(seed *uint64) uint16 {
	height := uint16(1)
	for height < MAX_HEIGHT && splitmix64.Splitmix64(seed)&1 == 0 {
		height++
	}
	return height
}

func (sk *SkipList) Set(key []byte, ts uint64, value []byte) bool {
	VKey := mkkey(sk.Arena, key, ts)
	if VKey == InvalidOffset {
		return false
	}

	valueAddr := sk.Arena.Alloc(uint32(len(value)))
	if valueAddr == uarena.InvalidOffset {
		return false
	}
	copy(sk.Arena.Data[valueAddr:valueAddr+uint32(len(value))], value)
	VValue := uint64(valueAddr)<<32 | uint64(len(value))

	var c ctx
	addr, ok := sk.seekContext(&c, key, ts)
	if ok {
		addr.SetValue(sk.Arena, VValue)
		return true
	}

	height := randHeight(&sk.RandSeed)
	node := Node{Height: height, Key: VKey, Value: VValue}
	size := sizeNode(&node)
	NAddr := sk.Arena.Alloc(size)
	if NAddr == uarena.InvalidOffset {
		return false
	}

	// link to next node
	for i := uint32(0); i < uint32(height); i++ {
		node.Next[i] = uint32(c[i].Next(sk.Arena, i))
		c[i].SetNext(sk.Arena, i, nodeAddr(NAddr))
	}

	storeNode(sk.Arena, NAddr, &node)
	return true
}
