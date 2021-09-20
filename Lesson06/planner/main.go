package main

import (
	"log"
	"runtime"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop() // go tool trace ./trace.out

	go log.Println("WORK!!!")
	for i := 0; ; i++ {
		if i%1000000 == 0 {
			runtime.Gosched() // GOMAXPROCS=1 go run .
		}
	}

}
