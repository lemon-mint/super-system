package mtab

import "v8.run/go/supersystem/lemondb/uarena"

type MemTable struct {
	Arena *uarena.UArena
}
