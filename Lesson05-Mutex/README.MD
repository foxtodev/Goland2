#### sync.Mutex and sync.RWMutex benchmark

##### sync.RWMutex 10% write / 90% read
```go
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
```

##### sync.RWMutex 50% write / 50% read
```go
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
```

##### sync.RWMutex 90% write / 10% read
```go
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
```

##### sync.Mutex
```go
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
```

#### Bench
##### bench n = 100
```
go test -bench=.
```
```
goos: linux
goarch: amd64
pkg: Lesson05-Mutex
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRWMutex_W10_R90-2        811257              1446 ns/op
BenchmarkRWMutex_W50_R50-2        554910              1936 ns/op
BenchmarkRWMutex_W90_R10-2        470217              2529 ns/op
BenchmarkMutex-2                  896714              1327 ns/op
PASS
ok      Lesson05-Mutex  7.379s
```
##### bench n = 1000
```
go test -bench=.
```
```
goos: linux
goarch: amd64
pkg: Lesson05-Mutex
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRWMutex_W10_R90-2         80690             12739 ns/op
BenchmarkRWMutex_W50_R50-2         61323             18913 ns/op
BenchmarkRWMutex_W90_R10-2         45489             25108 ns/op
BenchmarkMutex-2                   82431             13391 ns/op
PASS
ok      Lesson05-Mutex  5.237s
```
##### bench n = 100000
```
go test -bench=.
```
```
goos: linux
goarch: amd64
pkg: Lesson05-Mutex
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRWMutex_W10_R90-2           823           1278936 ns/op
BenchmarkRWMutex_W50_R50-2           572           1901530 ns/op
BenchmarkRWMutex_W90_R10-2           452           2511203 ns/op
BenchmarkMutex-2                     837           1326992 ns/op
PASS
ok      Lesson05-Mutex  5.189s
```