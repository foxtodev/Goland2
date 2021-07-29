package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Worker Pools
	count := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			count++
		}()
	}

	time.Sleep(2 * time.Second)
	if count == 1000 {
		fmt.Println("Success counting", count)
	} else {
		fmt.Println("Calculation is not finished", count)
	}

	//SIGTERM
	// ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancelFunc()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// go func(ctx context.Context) {
	// 	c <- syscall.SIGTERM
	// }(ctx)

	// for {
	// }

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	go func(ctx context.Context) {
		<-makeSignalShutdownChan()
	}(ctx)

}

func makeSignalShutdownChan() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	return c
}
