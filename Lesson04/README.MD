#### 1. // WORKER POOLS

```go

type work struct {
    id     int
    result int
}

count := 0
ch := make(chan work)

for i := 1; i <= 1000; i++ {
	id := i
	go func() {
		count++
		ch <- work{id, count}
	}()
}

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

```
output
```
WorkID 4 Received 1
WorkID 1 Received 2
WorkID 2 Received 3
WorkID 3 Received 4
WorkID 15 Received 5
WorkID 5 Received 6
WorkID 6 Received 7
...
WorkID 761 Received 118
WorkID 60 Received 119
WorkID 762 Received 120
WorkID 763 Received 121
...
WorkID 553 Received 657
WorkID 554 Received 658
WorkID 555 Received 659
WorkID 556 Received 660
WorkID 557 Received 661
...
WorkID 997 Received 997
WorkID 998 Received 998
WorkID 999 Received 999
WorkID 1000 Received 1000
Success counting 1000
```
#### 2. // SIGTERM

```go
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

```
output by timeout
```
 > Cancelled by timeout
```
output by CTRL+C
```
^C > Cancelled by SIGTERM
```