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

	// Random Read Benchmark
	for n := 0; n < 10; n++ {
		t := time.Now()
		for i := 0; i < 1<<10; i++ {
			// Copy 1KB of data from the file into the buffer
			copy(buffer, mem[i<<20:(i<<20)+1<<10])
		}
		fmt.Println("Random Read MMap:", time.Since(t), "n=", n)
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
// Random Read MMap: 76.726667ms n= 0
// Random Read MMap: 303.5µs n= 1
// Random Read MMap: 158.417µs n= 2
// Random Read MMap: 143.667µs n= 3
// Random Read MMap: 142µs n= 4
// Random Read MMap: 140.375µs n= 5
// Random Read MMap: 141.667µs n= 6
// Random Read MMap: 152.584µs n= 7
// Random Read MMap: 146.958µs n= 8
// Random Read MMap: 137.417µs n= 9
// Random Read Heap: 522.25µs n= 0
// Random Read Heap: 108.541µs n= 1
// Random Read Heap: 110.042µs n= 2
// Random Read Heap: 108.042µs n= 3
// Random Read Heap: 108.166µs n= 4
// Random Read Heap: 114.208µs n= 5
// Random Read Heap: 112.875µs n= 6
// Random Read Heap: 110.458µs n= 7
// Random Read Heap: 107.708µs n= 8
// Random Read Heap: 135.458µs n= 9
// Random Read MMap: 302.125µs n= 0
// Random Read MMap: 106.833µs n= 1
// Random Read MMap: 108.959µs n= 2
// Random Read MMap: 110.791µs n= 3
// Random Read MMap: 110.125µs n= 4
// Random Read MMap: 107.625µs n= 5
// Random Read MMap: 107.625µs n= 6
// Random Read MMap: 107.125µs n= 7
// Random Read MMap: 107.042µs n= 8
// Random Read MMap: 105.5µs n= 9
// Random Read Heap: 153.75µs n= 0
// Random Read Heap: 104.917µs n= 1
// Random Read Heap: 106.709µs n= 2
// Random Read Heap: 106.416µs n= 3
// Random Read Heap: 106.166µs n= 4
// Random Read Heap: 106.083µs n= 5
// Random Read Heap: 106.666µs n= 6
// Random Read Heap: 106.75µs n= 7
// Random Read Heap: 106.375µs n= 8
// Random Read Heap: 113.833µs n= 9
