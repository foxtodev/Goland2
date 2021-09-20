package main

import (
	"fmt"
	"github.com/qencept/fibonacci/cache"
	"github.com/qencept/fibonacci/closure"
	"github.com/qencept/fibonacci/recursion"
	"github.com/qencept/fibonacci/web"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	fiber := fibImplementations[rand.Intn(len(fibImplementations))]
	fmt.Printf("Starting %s\n", fiber.Name())
	web.Serve(fiber)
}
