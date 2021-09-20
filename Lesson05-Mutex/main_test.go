package main

import (
	"sync"
	"testing"
)

const n = 1000

func BenchmarkRWMutex_W10_R90(b *testing.B) {
	var mtx sync.RWMutex
	for i := 0; i < b.N; i++ {
		for w := 0; w < n/100*10; w++ {
			mtx.Lock()
			// Write
			mtx.Unlock()
		}
		for r := 0; r < n/100*90; r++ {
			mtx.RLock()
			// Read
			mtx.RUnlock()
		}
	}
}

func BenchmarkRWMutex_W50_R50(b *testing.B) {
	var mtx sync.RWMutex
	for i := 0; i < b.N; i++ {
		for w := 0; w < n/100*50; w++ {
			mtx.Lock()
			// Write
			mtx.Unlock()
		}
		for r := 0; r < n/100*50; r++ {
			mtx.RLock()
			// Read
			mtx.RUnlock()
		}
	}
}

func BenchmarkRWMutex_W90_R10(b *testing.B) {
	var mtx sync.RWMutex
	for i := 0; i < b.N; i++ {
		for w := 0; w < n/100*90; w++ {
			mtx.Lock()
			// Write
			mtx.Unlock()
		}
		for r := 0; r < n/100*10; r++ {
			mtx.RLock()
			// Read
			mtx.RUnlock()
		}
	}
}

func BenchmarkMutex(b *testing.B) {
	var mtx sync.Mutex
	for i := 0; i < b.N; i++ {
		for w := 0; w < n/2; w++ {
			mtx.Lock()
			// Write or Read
			mtx.Unlock()
		}
		for w := 0; w < n/2; w++ {
			mtx.Lock()
			// Write or Read
			mtx.Unlock()
		}
	}
}
