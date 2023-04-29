package mskip_test

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/lemon-mint/super-system/lemondb/mskip"
	"github.com/lemon-mint/super-system/lemondb/uarena"
)

const seed = 0x53514C6974652066

func TestSkipList(t *testing.T) {
	arena := uarena.NewUArena(1 << 16) // 64KB
	skip := mskip.NewSkipList(arena, seed)
	n := skip.PrepareInsert([]byte("key1"), []byte("value1"), 100)
	skip.Insert(n)
	if v, ok := skip.Get([]byte("key1"), 200); !ok || !bytes.Equal(v, []byte("value1")) {
		t.Errorf("Get failed, expected %v, got %v", []byte("value1"), v)
	}

	if v, ok := skip.Get([]byte("key1"), 50); ok || v != nil {
		t.Errorf("Get failed, expected %v, got %v", nil, v)
	}

	if v, ok := skip.Get([]byte("key1"), 100); !ok || !bytes.Equal(v, []byte("value1")) {
		t.Errorf("Get failed, expected %v, got %v", []byte("value1"), v)
	}

	n = skip.PrepareInsert([]byte("key2"), []byte("value2"), 200)
	skip.Insert(n)

	if v, ok := skip.Get([]byte("key2"), 300); !ok || !bytes.Equal(v, []byte("value2")) {
		t.Error("Get failed")
	}

	if v, ok := skip.Get([]byte("key2"), 150); ok || v != nil {
		t.Error("Get failed")
	}

	n = skip.PrepareInsert([]byte("key1"), []byte("value3"), 300) // update
	skip.Insert(n)

	if v, ok := skip.Get([]byte("key1"), 400); !ok || !bytes.Equal(v, []byte("value3")) {
		t.Error("Get failed")
	}

	if v, ok := skip.Get([]byte("key1"), 250); !ok || !bytes.Equal(v, []byte("value1")) {
		t.Error("Get failed")
	}

	skip.Reset()
	if v, ok := skip.Get([]byte("key1"), 100); ok || v != nil {
		t.Error("Get failed")
	}

	if v, ok := skip.Get([]byte("key2"), 200); ok || v != nil {
		t.Error("Get failed")
	}

	if v, ok := skip.Get([]byte("key1"), 300); ok || v != nil {
		t.Error("Get failed")
	}
}

func BenchmarkSkipListGet(b *testing.B) {
	arena := uarena.NewUArena(1 << 20) // 1MB
	skip := mskip.NewSkipList(arena, seed)
	for i := 0; i < 1000; i++ {
		n := skip.PrepareInsert([]byte("key"+strconv.Itoa(i)), []byte("value"+strconv.Itoa(i)), uint64(i))
		skip.Insert(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, ok := skip.Get([]byte("key500"), 500)
		if !ok {
			b.Fatalf("Get failed")
			b.FailNow()
		}
	}
}

func TestIterator(t *testing.T) {
	arena := uarena.NewUArena(1 << 20) // 1MB
	skip := mskip.NewSkipList(arena, seed)
	for i := 0; i < 1000; i++ {
		n := skip.PrepareInsert([]byte("key"+strings.Repeat("0", 3-len(strconv.Itoa(i)))+strconv.Itoa(i)), []byte("value"+strconv.Itoa(i)), uint64(i))
		if n == mskip.InvalidNode {
			t.Fatal("PrepareInsert failed")
			t.FailNow()
		}
		skip.Insert(n)
	}

	iter := mskip.NewIterator(skip)
	for i := 0; i < 1000; i++ {
		if !iter.Valid() {
			t.Fatal("Iterator not valid")
			t.FailNow()
		}
		if !bytes.Equal(iter.Key(), []byte("key"+strings.Repeat("0", 3-len(strconv.Itoa(i)))+strconv.Itoa(i))) {
			t.Fatalf("Iterator key not match, expected %s, got %s", []byte("key"+strings.Repeat("0", 3-len(strconv.Itoa(i)))+strconv.Itoa(i)), iter.Key())
			t.FailNow()
		}
		if !bytes.Equal(iter.Value(), []byte("value"+strconv.Itoa(i))) {
			t.Fatalf("Iterator value not match, expected %s, got %s", []byte("value"+strconv.Itoa(i)), iter.Value())
			t.FailNow()
		}
		iter.Next()
	}

	iter.Seek([]byte("key999"), 999)
	if !iter.Valid() {
		t.Fatal("Iterator not valid")
		t.FailNow()
	}

	if !bytes.Equal(iter.Key(), []byte("key999")) {
		t.Fatalf("Iterator key not match, expected %s, got %s", []byte("key999"), iter.Key())
		t.FailNow()
	}

	iter.Prev()
	if !iter.Valid() {
		t.Fatal("Iterator not valid")
		t.FailNow()
	}

	if !bytes.Equal(iter.Key(), []byte("key998")) {
		t.Fatalf("Iterator key not match, expected %s, got %s", []byte("key998"), iter.Key())
		t.FailNow()
	}

	iter.Seek([]byte("key"), 0)
	if !iter.Valid() {
		t.Fatal("Iterator not valid")
		t.FailNow()
	}

	if !bytes.Equal(iter.Key(), []byte("key000")) {
		t.Fatalf("Iterator key not match, expected %s, got %s", []byte("key000"), iter.Key())
		t.FailNow()
	}

	for i := 0; i < 1000; i++ {
		if !iter.Valid() {
			t.Fatal("Iterator not valid")
			t.FailNow()
		}
		if !bytes.Equal(iter.Key(), []byte("key"+strings.Repeat("0", 3-len(strconv.Itoa(i)))+strconv.Itoa(i))) {
			t.Fatalf("Iterator key not match, expected %s, got %s", []byte("key"+strings.Repeat("0", 3-len(strconv.Itoa(i)))+strconv.Itoa(i)), iter.Key())
			t.FailNow()
		}
		if !bytes.Equal(iter.Value(), []byte("value"+strconv.Itoa(i))) {
			t.Fatalf("Iterator value not match, expected %s, got %s", []byte("value"+strconv.Itoa(i)), iter.Value())
			t.FailNow()
		}
		if iter.Timestamp() != uint64(i) {
			t.Fatalf("Iterator timestamp not match, expected %d, got %d", i, iter.Timestamp())
			t.FailNow()
		}
		iter.Next()
	}
}
