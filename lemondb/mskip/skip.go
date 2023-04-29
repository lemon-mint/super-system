package mskip

import (
	"bytes"
	"sync/atomic"
	"unsafe"

	"github.com/lemon-mint/super-system/lemondb/uarena"
	"github.com/lemon-mint/super-system/lemondb/wyhash/splitmix64"
)

const MAX_HEIGHT = 20
const TIMESTAMP_SIZE = 8

// Key Format:
// <key data> <inverted big endian timestamp>
// <32 bit offset> <32 bit size> = <64 bit key>

// Value Format:
// <value data>
// <32 bit offset> <32 bit size> = <64 bit value>

func timestamp(key []byte) uint64 {
	offset := len(key) - TIMESTAMP_SIZE
	_ = key[offset+7]
	ts := uint64(key[offset])<<56 | uint64(key[offset+1])<<48 | uint64(key[offset+2])<<40 | uint64(key[offset+3])<<32 | uint64(key[offset+4])<<24 | uint64(key[offset+5])<<16 | uint64(key[offset+6])<<8 | uint64(key[offset+7])
	return ^ts // inverted
}

func dataKey(key []byte) []byte {
	return key[:len(key)-TIMESTAMP_SIZE]
}

type node struct {
	Key    uint64 // Offset: 0, Size: 8 (immutable)
	Value  uint64 // Offset: 8, Size: 8 (mutable, need Atomic Update)
	Height uint32 // Offset: 16, Size: 4 (immutable)

	Next [MAX_HEIGHT]uint32 // Offset: 20+4*i, Size: 4, (i = 0..MAX_HEIGHT-1) (mutable, need Atomic Update)
}

type SkipList struct {
	arena *uarena.UArena
	head  NodeRef
	seed  uint64
}

type NodeRef uint32

const InvalidNode = NodeRef(1<<32 - 1)

func (n NodeRef) IsValid() bool {
	return n != InvalidNode
}

func (n NodeRef) Height(arena *uarena.UArena) uint32 {
	return *(*uint32)(unsafe.Pointer(&arena.Data[n+16]))
}

func (n NodeRef) SetHeight(arena *uarena.UArena, height uint32) {
	*(*uint32)(unsafe.Pointer(&arena.Data[n+16])) = height
}

func (n NodeRef) KeyUint64(arena *uarena.UArena) uint64 {
	return *(*uint64)(unsafe.Pointer(&arena.Data[n]))
}

func (n NodeRef) Key(arena *uarena.UArena) []byte {
	key := n.KeyUint64(arena)
	offset, size := uint32(key>>32), uint32(key)
	return arena.Data[offset : offset+size]
}

func (n NodeRef) SetKeyUint64(arena *uarena.UArena, key uint64) {
	*(*uint64)(unsafe.Pointer(&arena.Data[n])) = key
}

func (n NodeRef) ValueUint64(arena *uarena.UArena) uint64 {
	return atomic.LoadUint64((*uint64)(unsafe.Pointer(&arena.Data[n+8])))
}

func (n NodeRef) SetValueUint64(arena *uarena.UArena, value uint64) {
	atomic.StoreUint64((*uint64)(unsafe.Pointer(&arena.Data[n+8])), value)
}

func (n NodeRef) Value(arena *uarena.UArena) []byte {
	value := n.ValueUint64(arena)
	offset, size := uint32(value>>32), uint32(value)
	return arena.Data[offset : offset+size]
}

func (n NodeRef) Next(arena *uarena.UArena, i uint32) NodeRef {
	return NodeRef(atomic.LoadUint32((*uint32)(unsafe.Pointer(&arena.Data[uint32(n)+20+4*i]))))
}

func (n NodeRef) SetNext(arena *uarena.UArena, i uint32, next NodeRef) {
	atomic.StoreUint32((*uint32)(unsafe.Pointer(&arena.Data[uint32(n)+20+4*i])), uint32(next))
}

func (n NodeRef) Size(arena *uarena.UArena) uint32 {
	return 8 + 8 + 4 + n.Height(arena)*4
}

func NewSkipList(arena *uarena.UArena, seed uint64) *SkipList {
	sk := &SkipList{
		arena: arena,
		seed:  seed,
	}
	if initskip(sk) != 0 {
		return nil
	}
	return sk
}

func initskip(sk *SkipList) int {
	sk.arena.Alloc(8 + 8 + 4 + MAX_HEIGHT*4) // Allocate space for tail node (not used)
	sk.seed = splitmix64.Splitmix64(&sk.seed)
	head := sk.arena.Alloc(8 + 8 + 4 + MAX_HEIGHT*4) // Allocate space for head node
	if head == uarena.OutOfMemory {
		return -1
	}
	sk.head = NodeRef(head)
	sk.head.SetHeight(sk.arena, MAX_HEIGHT)
	for i := uint32(0); i < MAX_HEIGHT; i++ {
		sk.head.SetNext(sk.arena, i, 0)
	}
	sk.head.SetKeyUint64(sk.arena, 0)
	sk.head.SetValueUint64(sk.arena, 0)
	return 0
}

func calcSize(height uint32, keySize, valueSize uint32) uint32 {
	return 8 + // Key
		8 + // Value
		4 + // Height
		4*height + // Next
		keySize + TIMESTAMP_SIZE + // Key Data + Timestamp
		valueSize // Value Data
}

func randHeight(seed *uint64) uint16 {
	height := uint16(1)
	for height < MAX_HEIGHT && splitmix64.Splitmix64(seed)&1 == 0 {
		height++
	}
	return height
}

// PrepareInsert returns a new node to be inserted into the skip list.
func (sk *SkipList) PrepareInsert(key, value []byte, ts uint64) NodeRef {
	height := randHeight(&sk.seed)
	size := calcSize(uint32(height), uint32(len(key)), uint32(len(value)))
	buffer := sk.arena.Alloc(size)
	if buffer == uarena.OutOfMemory {
		return InvalidNode
	}
	newNode := NodeRef(buffer)
	newNode.SetHeight(sk.arena, uint32(height))
	buffer += 8 + 8 + 4 + 4*uint32(height)

	KeyV64 := uint64(buffer)<<32 | uint64(len(key)+TIMESTAMP_SIZE)
	newNode.SetKeyUint64(sk.arena, KeyV64)

	copy(sk.arena.Data[buffer:], key)
	buffer += uint32(len(key))
	ts = ^ts
	_ = sk.arena.Data[buffer+7]
	sk.arena.Data[buffer] = byte(ts >> 56)
	sk.arena.Data[buffer+1] = byte(ts >> 48)
	sk.arena.Data[buffer+2] = byte(ts >> 40)
	sk.arena.Data[buffer+3] = byte(ts >> 32)
	sk.arena.Data[buffer+4] = byte(ts >> 24)
	sk.arena.Data[buffer+5] = byte(ts >> 16)
	sk.arena.Data[buffer+6] = byte(ts >> 8)
	sk.arena.Data[buffer+7] = byte(ts)
	buffer += TIMESTAMP_SIZE

	copy(sk.arena.Data[buffer:], value)
	ValueV64 := uint64(buffer)<<32 | uint64(len(value))
	newNode.SetValueUint64(sk.arena, ValueV64)

	return newNode
}

// Insert inserts a new node into the skip list.
func (sk *SkipList) Insert(newNode NodeRef) {
	key := newNode.Key(sk.arena)
	height := newNode.Height(sk.arena)
	keyData := dataKey(key)
	KeyTS := timestamp(key)
	var prev [MAX_HEIGHT]NodeRef
	for i := uint32(0); i < MAX_HEIGHT; i++ {
		prev[i] = sk.head
	}

	x := sk.head
	for i := uint32(MAX_HEIGHT - 1); i < MAX_HEIGHT; i-- {
	L:
		for {
			next := x.Next(sk.arena, i)
			if next == 0 { // Reached tail
				break
			}
			nextKey := next.Key(sk.arena)
			nextData := dataKey(nextKey)
			switch bytes.Compare(keyData, nextData) {
			case -1: // (key < nextKey)
				break L
			case 0: // (key == nextKey)
				nextTS := timestamp(nextKey)
				if KeyTS > nextTS {
					break L
				}

				if KeyTS == nextTS {
					// Same key and timestamp. Overwrite the value.
					next.SetValueUint64(sk.arena, newNode.ValueUint64(sk.arena))
					return
				}
			}
			x = next
		}
		prev[i] = x
	}

	// Prepare the new node.
	for i := uint32(0); i < height; i++ {
		newNode.SetNext(sk.arena, i, prev[i].Next(sk.arena, i))
	}

	// Atomically insert the new node into the skip list.
	// Update the next pointers of the previous nodes.
	for i := uint32(0); i < height; i++ {
		prev[i].SetNext(sk.arena, i, newNode)
	}
}

func (sk *SkipList) Delete(key []byte, ts uint64) bool {
	height := randHeight(&sk.seed)
	size := calcSize(uint32(height), uint32(len(key)), 0)
	buffer := sk.arena.Alloc(size)
	if buffer == uarena.OutOfMemory {
		return false
	}
	newNode := NodeRef(buffer)
	newNode.SetHeight(sk.arena, uint32(height))
	buffer += 8 + 8 + 4 + 4*uint32(height)

	KeyV64 := uint64(buffer)<<32 | uint64(len(key)+TIMESTAMP_SIZE)
	newNode.SetKeyUint64(sk.arena, KeyV64)

	copy(sk.arena.Data[buffer:], key)
	buffer += uint32(len(key))
	ts = ^ts
	_ = sk.arena.Data[buffer+7]
	sk.arena.Data[buffer] = byte(ts >> 56)
	sk.arena.Data[buffer+1] = byte(ts >> 48)
	sk.arena.Data[buffer+2] = byte(ts >> 40)
	sk.arena.Data[buffer+3] = byte(ts >> 32)
	sk.arena.Data[buffer+4] = byte(ts >> 24)
	sk.arena.Data[buffer+5] = byte(ts >> 16)
	sk.arena.Data[buffer+6] = byte(ts >> 8)
	sk.arena.Data[buffer+7] = byte(ts)

	// Insert Delete Marker
	newNode.SetValueUint64(sk.arena, 0)
	sk.Insert(newNode)

	return true
}

// seekLT returns the last node with a key < the given key and timestamp < the given timestamp.
func (sk *SkipList) seekLT(key []byte, ts uint64) NodeRef {
	x := sk.head
	for i := uint32(MAX_HEIGHT - 1); i < MAX_HEIGHT; i-- {
	L:
		for {
			next := x.Next(sk.arena, i)
			if next == 0 { // Reached tail
				break
			}
			nextKey := next.Key(sk.arena)
			nextData := dataKey(nextKey)
			nextTS := timestamp(nextKey)
			switch bytes.Compare(key, nextData) {
			case -1: // (key < nextKey)
				break L
			case 0: // (key == nextKey)
				if ts >= nextTS { // (ts >= nextTS)
					break L
				}
			}
			x = next
		}
	}
	// x.key < key && x.ts < ts
	return x
}

// seekGE returns the first node with a key >= the given key and timestamp >= the given timestamp.
func (sk *SkipList) seekGE(key []byte, ts uint64) NodeRef {
	x := sk.seekLT(key, ts)
	// x.key < key && x.ts < ts

	next := x.Next(sk.arena, 0)
	if next == 0 { // Reached tail
		return InvalidNode
	}

	return next
}

// Get returns the value for the given key at the given timestamp.
func (sk *SkipList) Get(key []byte, ts uint64) ([]byte, bool) {
	x := sk.seekGE(key, ts)
	if x == InvalidNode {
		return nil, false
	}

	nextKey := x.Key(sk.arena)
	nextData := dataKey(nextKey)
	nextTS := timestamp(nextKey)
	if bytes.Equal(key, nextData) && ts >= nextTS && x.ValueUint64(sk.arena) != 0 {
		return x.Value(sk.arena), true
	}

	return nil, false
}

func (sk *SkipList) Reset() {
	sk.arena.Reset()
	sk.head = 0
	initskip(sk)
}

type Iterator struct {
	sk *SkipList
	x  NodeRef
}

func NewIterator(sk *SkipList) *Iterator {
	return &Iterator{
		sk: sk,
		x:  sk.head.Next(sk.arena, 0),
	}
}

func (it *Iterator) Valid() bool {
	return it.x.IsValid()
}

func (it *Iterator) Key() []byte {
	k := it.x.Key(it.sk.arena)
	return dataKey(k)
}

func (it *Iterator) Value() []byte {
	return it.x.Value(it.sk.arena)
}

func (it *Iterator) Timestamp() uint64 {
	k := it.x.Key(it.sk.arena)
	return timestamp(k)
}

func (it *Iterator) Seek(key []byte, ts uint64) {
	it.x = it.sk.seekGE(key, ts)
}

func (it *Iterator) Next() {
	it.x = it.x.Next(it.sk.arena, 0)
}

func (it *Iterator) Prev() {
	if it.x == it.sk.head {
		return
	}

	if !it.x.IsValid() {
		panic("mskip: prev on invalid iterator")
	}

	key := it.x.Key(it.sk.arena)
	keyData := dataKey(key)
	KeyTS := timestamp(key)

	it.x = it.sk.seekLT(keyData, KeyTS)
}
