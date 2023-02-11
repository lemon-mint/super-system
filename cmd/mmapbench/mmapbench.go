package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/edsrzf/mmap-go"
)

var buffer []byte = make([]byte, 1<<20) // 1MB

func main() {
	f, err := os.Create("test.bin")
	if err != nil {
		panic(err)
	}
	defer os.Remove("test.bin")
	defer f.Close()

	// Write 1GB of data to the file
	for i := 0; i < 1<<10; i++ {
		_, err = rand.Read(buffer)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(buffer)
		if err != nil {
			panic(err)
		}
	}
	// Sync the file to disk
	err = f.Sync()
	if err != nil {
		panic(err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	mem, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer mem.Unmap()

	// Random Read Benchmark
	for n := 0; n < 10; n++ {
		t := time.Now()
		for i := 0; i < 1<<10; i++ {
			// Copy 1KB of data from the file into the buffer
			copy(buffer, mem[i<<20:(i<<20)+1<<10])
		}
		fmt.Println("Random Read MMap:", time.Since(t), "n=", n)
	}

	// Do it again to heap memory
	newmem := make([]byte, 1<<30)
	for i := 0; i < 1<<10; i++ {
		_, err = rand.Read(newmem[i<<20 : (i<<20)+1<<10])
		if err != nil {
			panic(err)
		}
	}

	// Random Read Benchmark
	for n := 0; n < 10; n++ {
		t := time.Now()
		for i := 0; i < 1<<10; i++ {
			// Copy 1KB of data from the file into the buffer
			copy(buffer, newmem[i<<20:(i<<20)+1<<10])
		}
		fmt.Println("Random Read Heap:", time.Since(t), "n=", n)
	}
}

// Results on Apple Silicon M1
//
// Random Read MMap: 84.789291ms n= 0
// Random Read MMap: 234.667µs n= 1
// Random Read MMap: 182.042µs n= 2
// Random Read MMap: 156.791µs n= 3
// Random Read MMap: 173.167µs n= 4
// Random Read MMap: 174.333µs n= 5
// Random Read MMap: 164.5µs n= 6
// Random Read MMap: 162.292µs n= 7
// Random Read MMap: 157µs n= 8
// Random Read MMap: 154.292µs n= 9
// Random Read Heap: 648.459µs n= 0
// Random Read Heap: 117.458µs n= 1
// Random Read Heap: 127.333µs n= 2
// Random Read Heap: 115µs n= 3
// Random Read Heap: 143.125µs n= 4
// Random Read Heap: 121.625µs n= 5
// Random Read Heap: 119.875µs n= 6
// Random Read Heap: 119.417µs n= 7
// Random Read Heap: 158.291µs n= 8
// Random Read Heap: 114.209µs n= 9
