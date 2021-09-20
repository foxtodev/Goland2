package cache

type ctx struct {
	cache map[int]int
}

func NewCache() *ctx {
	return &ctx{cache: map[int]int{}}
}

func (c *ctx) cacheFib(n int) int {
	v, ok := c.cache[n]
	if !ok {
		v = c.Fib(n)
		c.cache[n] = v
	}
	return v
}

func (c *ctx) Fib(n int) int {
	if n <= 1 {
		return n
	}
	return c.cacheFib(n-1) + c.cacheFib(n-2)
}

func (c *ctx) Name() string {
	return "Fibonacci (Cache)"
}
