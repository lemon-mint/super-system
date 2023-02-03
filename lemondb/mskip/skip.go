package mskip

import (
	"v8.run/go/supersystem/lemondb/uarena"
)

const MAX_HEIGHT = 32

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

type SkipList struct {
	Arena *uarena.UArena
	Head  uint64
}
