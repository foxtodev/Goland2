package main

import (
	"fmt"
	"sync"

	"github.com/pkg/profile"
)

// Structure struct
type Structure struct {
	sync.Mutex
	number int
}

// Func doing something
func (c *Structure) Inc(ch chan int) {
	c.Lock()
	defer c.Unlock()
	c.number++
	ch <- c.number
}

func main() {

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop() // MutexProfile // BlockProfile
	// Eqauls
	// p := profile.Start(profile.TraceProfile, profile.ProfilePath("."))
	// defer p.Stop()

	const n = 100
	var wg sync.WaitGroup
	str := &Structure{}
	ch := make(chan int)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go str.Inc(ch)
		wg.Done()
	}

	wg.Wait()

	for i := 1; i <= n; i++ {
		val := <-ch
		fmt.Println(val)
	}
}
