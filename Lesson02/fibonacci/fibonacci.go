package fibonacci

// FibCycle a function that uses a loop. Returns Nth fibonacci number.
func FibCycle(number int) int {
	fib1, fib2 := 0, 1
	for ; number > 0; number-- {
		fib1, fib2 = fib2, fib1+fib2
	}
	return fib1
}

// FibRecursion a function that uses recursion. Returns Nth fibonacci number.
func FibRecursion(number int) int {
	if number < 2 {
		return number
	}
	return FibRecursion(number-1) + FibRecursion(number-2)
}

// FibSlice a function that uses slice. Returns Nth fibonacci number.
func FibSlice(number int) int {
	if number < 2 {
		return number
	}
	fib := []int{0, 1}
	for i := 2; i <= number; i++ {
		fib = []int{fib[1], fib[0] + fib[1]}
	}
	return fib[1]
}

// FibFunc a function that uses a closure. Returns Nth fibonacci number.
func FibFunc() func() int {
	fib1, fib2 := 1, 0
	return func() int {
		fib1, fib2 = fib2, fib1+fib2
		return fib1
	}
}

// FibMap a function that uses map. Returns Nth fibonacci number.
func FibMap(number int) int {
	fmap := make(map[int]int)
	fmap[0], fmap[1] = 0, 1
	for i := 2; i <= number; i++ {
		fmap[i] = fmap[i-1] + fmap[i-2]
	}
	return fmap[number]
}

// FibMapV2 a second version of function that uses map. Returns Nth fibonacci number.
func FibMapV2(number int) int {
	fib := map[int]int{
		0: 0,
		1: 1,
	}
	return fibonacciMapV2Rec(number, fib)
}

func fibonacciMapV2Rec(number int, fibc map[int]int) int {
	val, exists := fibc[number]
	if exists {
		return val
	}
	fibc[number] = fibonacciMapV2Rec(number-1, fibc) + fibonacciMapV2Rec(number-2, fibc)
	return fibc[number]
}

var fibcache = map[int]int{} // Cache of fibonacci numbers

//FibCache a function that uses cache(map). Returns Nth fibonacci number.
func FibCache(number int) int {
	val, exists := fibcache[number]
	if !exists {
		val = fibonacciCacheRec(number)
		fibcache[number] = val
	}
	return val
}

func fibonacciCacheRec(number int) int {
	if number < 2 {
		return number
	}
	return fibonacciCacheRec(number-1) + fibonacciCacheRec(number-2)
}

// FibMapFunc a second version of function that uses a closure and cache. Returns Nth fibonacci number.
func FibMapFunc() func() int {
	fib := make(map[int]int)
	fib[0], fib[1] = 0, 1
	number := -1
	return func() int {
		number++
		if val, exists := fib[number]; exists {
			return val
		}
		fib[number] = fib[number-1] + fib[number-2]
		return fib[number]
	}
}
