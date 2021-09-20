package fibonacci

import (
	"fmt"
	"testing"
)

var table = []struct {
	arg  int
	want int
}{
	{0, 0},
	{1, 1},
	{10, 55},
	{20, 6765},
}

const N = 20

var sink int

func TestFibCycle(t *testing.T) {
	for _, entry := range table {
		got := FibCycle(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibCycle(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibCycle(N)
	}
	sink = res
}

func ExampleFibCycle() {
	fmt.Println(FibCycle(20))
	// Output: 6765
}

func TestFibRecursion(t *testing.T) {
	for _, entry := range table {
		got := FibRecursion(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibRecursion(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibRecursion(N)
	}
	sink = res
}

func ExampleFibRecursion() {
	fmt.Println(FibRecursion(20))
	// Output: 6765
}

func TestFibSlice(t *testing.T) {
	for _, entry := range table {
		got := FibSlice(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibSlice(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibSlice(N)
	}
	sink = res
}

func ExampleFibSlice() {
	fmt.Println(FibSlice(20))
	// Output: 6765
}

func TestFibMap(t *testing.T) {
	for _, entry := range table {
		got := FibMap(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibMap(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibMap(N)
	}
	sink = res
}

func ExampleFibMap() {
	fmt.Println(FibMap(20))
	// Output: 6765
}

func TestFibMapV2(t *testing.T) {
	for _, entry := range table {
		got := FibMapV2(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibMapV2(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibMapV2(N)
	}
	sink = res
}

func ExampleFibMapV2() {
	fmt.Println(FibMapV2(20))
	// Output: 6765
}

func TestFibCache(t *testing.T) {
	for _, entry := range table {
		got := FibCache(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibCache(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibCache(N)
	}
	sink = res
}

func ExampleFibCache() {
	fmt.Println(FibCache(20))
	// Output: 6765
}
