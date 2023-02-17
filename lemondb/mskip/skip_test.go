package mskip_test

import (
	"bytes"
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
		t.Error("Get failed")
	}

	if v, ok := skip.Get([]byte("key1"), 50); ok || v != nil {
		t.Error("Get failed")
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
}
