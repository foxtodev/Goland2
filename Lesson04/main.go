package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type work struct {
	id     int
	result int
}

func main() {

	// WORKER POOLS

	count := 0
	ch := make(chan work)
	//defer close(ch)

	for i := 1; i <= 1000; i++ {
		id := i
		go func() {
			count++
			ch <- work{id, count}
		}()
	}

	// time.Sleep(1 * time.Second) // wait start gourines
	// for val := range ch {
	// 	fmt.Println("WorkID", val.id, "Received", val.result)
	// 	if val.result == 1000 {
	// 		break
	// 	}
	// }

	for {
		val, _ := <-ch
		fmt.Println("WorkID", val.id, "Received", val.result)
		if val.result == 1000 {
			break
		}
	}
	if count == 1000 {
		fmt.Println("Success counting", count)
	} else {
		fmt.Println("Calculation is not finished", count)
	}

	// SIGTERM

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)
	go func(ctx context.Context) {
		<-chSig
		cancelFunc()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println(" > Cancelled by timeout")
		case context.Canceled:
			fmt.Println(" > Cancelled by SIGTERM")
		}
	}

	// checking whether the channel is closed

	// chSig := make(chan os.Signal, 1)
	// signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-chSig
	// 	close(chSig)
	// }()
	// for val := range chSig {
	// 	fmt.Println(" > Cancelled by SIGTERM", val)
	// 	time.Sleep(1 * time.Second)
	// 	os.Exit(0)
	// }

	// with sync

	// var wg sync.WaitGroup
	// wg.Add(1)
	// chSig := make(chan os.Signal, 1)
	// signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-chSig
	// 	time.Sleep(1 * time.Second)
	// 	wg.Done()
	// 	os.Exit(0)
	// }()
	// wg.Wait()

}
