package mskip_test

import (
	"bytes"
	"testing"

	"v8.run/go/supersystem/lemondb/mskip"
	"v8.run/go/supersystem/lemondb/uarena"
)

const seed = 0x53514C6974652066

func TestSkipList(t *testing.T) {
	arena := uarena.NewUArena(1 << 16) // 64KB
	skip := mskip.NewSkipList(arena, seed)
	n := skip.PrepareInsert([]byte("key1"), []byte("value1"), 1)
	skip.Insert(n)
	if v, ok := skip.Get([]byte("key1"), 100); !ok || !bytes.Equal(v, []byte("value1")) {
		t.Error("Get failed")
	}
}
