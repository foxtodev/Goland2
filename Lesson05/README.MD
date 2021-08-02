#### Run n goroutines and wait for their completion
```go
func main() {

	const n = 100
	var wg sync.WaitGroup
	str := &Structure{}
	ch := make(chan int)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go str.Func(ch)
		wg.Done()
	}

	wg.Wait()

	for i := 1; i <= n; i++ {
		val := <-ch
		fmt.Println(val)

	}
}
```
#### defer mutex
```go
// Structure struct
type Structure struct {
	sync.Mutex
	number int
}

// Func doing something
func (c *Structure) Func(ch chan int) {
	c.Lock()
	defer c.Unlock()
	// do something with number
	c.number++
	ch <- c.number
}
```