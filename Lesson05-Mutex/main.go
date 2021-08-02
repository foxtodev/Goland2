package main

import "sync"

// RWSet structure
type RWSet struct {
	mu sync.RWMutex
}

// Set structure
type Set struct {
	mu sync.Mutex
}

// RWMuWrite sync.RWMutex write
func (s *RWSet) RWMuWrite(x interface{}) {
	s.mu.Lock()
	s.mu.Unlock()
}

// RWMuRead sync.RWMutex read
func (s *RWSet) RWMuRead(x interface{}) {
	s.mu.RLock()
	s.mu.RUnlock()
}

// MuReadWrite  sync.Mutex write and read
func (s *RWSet) MuReadWrite(x interface{}) {
	s.mu.RLock()
	s.mu.RUnlock()
}

func main() {

}
