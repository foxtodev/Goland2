package recursion

type ctx struct {
}

func NewRecursion() *ctx {
	return &ctx{}
}

func (c *ctx) Fib(n int) int {
	if n <= 1 {
		return n
	}
	return c.Fib(n-1) + c.Fib(n-2)
}

func (c *ctx) Name() string {
	return "Fibonacci (Recursion)"
}
