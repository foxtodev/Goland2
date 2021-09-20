package web

type Fiber interface {
	Fib(int) int
	Name() string
}
