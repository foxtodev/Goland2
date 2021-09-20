package main

import (
	"log"
)

// Counter struct
type Counter struct {
	Cnt int64
}

// Inc func
func (c *Counter) Inc(ch chan int64) {
	c.Cnt++
	ch <- c.Cnt
}

func main() {

	c := &Counter{}
	ch := make(chan int64, 10)

	for i := 0; i < 100; i++ {
		go c.Inc(ch)
	}

	for i := 0; i < 100; i++ {
		v := <-ch
		if v != int64(i+1) {
			log.Println(i+1, " != ", v)
		}
	}
}
