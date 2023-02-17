package replication_test

import (
	"testing"

	"github.com/lemon-mint/super-system/replication"
)

func assertnil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

type Message struct {
}

func TestLogAppend(t *testing.T) {
	ls := replication.NewMemoryLog[*Message](4096)
	msg := &Message{}

	assertnil(t, ls.Append(msg, 0))
	assertnil(t, ls.Append(msg, 1))
	assertnil(t, ls.Append(msg, 2))
}

func TestLogGet(t *testing.T) {
	ls := replication.NewMemoryLog[*Message](4096)
	msg := &Message{}

	assertnil(t, ls.Append(msg, 0))
	assertnil(t, ls.Append(msg, 1))
	assertnil(t, ls.Append(msg, 2))

	_, err := ls.Get(0)
	assertnil(t, err)

	_, err = ls.Get(1)
	assertnil(t, err)

	_, err = ls.Get(2)
	assertnil(t, err)
}

func TestLogLastIndex(t *testing.T) {
	ls := replication.NewMemoryLog[*Message](4096)
	msg := &Message{}

	assertnil(t, ls.Append(msg, 0))
	assertnil(t, ls.Append(msg, 1))
	assertnil(t, ls.Append(msg, 2))

	n, err := ls.LastIndex()
	assertnil(t, err)

	if n != 2 {
		t.Fatal("LastIndex() returned wrong value")
	}
}

func TestLogFirstIndex(t *testing.T) {
	ls := replication.NewMemoryLog[*Message](4096)
	msg := &Message{}

	assertnil(t, ls.Append(msg, 0))
	assertnil(t, ls.Append(msg, 1))
	assertnil(t, ls.Append(msg, 2))
	assertnil(t, ls.Append(msg, 3))
	assertnil(t, ls.Append(msg, 4))

	assertnil(t, ls.Truncate(2))

	n, err := ls.FirstIndex()
	assertnil(t, err)

	if n != 2 {
		t.Fatalf("FirstIndex() returned %d, expected 2", n)
	}
}

func TestLogTruncate(t *testing.T) {
	ls := replication.NewMemoryLog[*Message](4096)
	msg := &Message{}

	assertnil(t, ls.Append(msg, 0))
	assertnil(t, ls.Append(msg, 1))
	assertnil(t, ls.Append(msg, 2))
	assertnil(t, ls.Append(msg, 3))
	assertnil(t, ls.Append(msg, 4))

	assertnil(t, ls.Truncate(2))

	n, err := ls.LastIndex()
	assertnil(t, err)

	if n != 4 {
		t.Fatalf("ls.LastIndex() returned %d, expected 4", n)
	}

	n, err = ls.FirstIndex()
	assertnil(t, err)

	if n != 2 {
		t.Fatalf("FirstIndex() returned %d, expected 2", n)
	}

	_, err = ls.Get(0)
	if err != replication.ErrInvalidIndex {
		t.Fatal("Get(0) should have returned ErrInvalidIndex")
	}

	_, err = ls.Get(1)
	if err != replication.ErrInvalidIndex {
		t.Fatal("Get(1) should have returned ErrInvalidIndex")
	}

	_, err = ls.Get(2)
	assertnil(t, err)

	assertnil(t, ls.Append(msg, 5))
	assertnil(t, ls.Append(msg, 6))

	n, err = ls.LastIndex()
	assertnil(t, err)
	if n != 6 {
		t.Fatalf("ls.LastIndex() returned %d, expected 6", n)
	}

	n, err = ls.FirstIndex()
	assertnil(t, err)
	if n != 2 {
		t.Fatalf("FirstIndex() returned %d, expected 2", n)
	}

	_, err = ls.Get(0)
	if err != replication.ErrInvalidIndex {
		t.Fatal("Get(0) should have returned ErrInvalidIndex")
	}
}
