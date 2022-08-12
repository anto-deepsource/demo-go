package main

import (
	"fmt"
	"sync"
	"testing"
)

var dataLarge []byte

const size = 64 * 1024 // 65536

func benchmarkLargeSizePool(b *testing.B) {
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, size)
			return b
		},
	}
	for i := 0; i < b.N; i++ {
		dataLarge = bytePool.Get().([]byte)
		_ = dataLarge
		bytePool.Put(dataLarge)
	}
	
	m := new(sync.Map)
	
	actual, ok := m.Load("key")
	if ok {
		m.Delete("key")
		fmt.Println(actual)
	}
}
