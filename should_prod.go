//go:build !assert

package should

func Be(cond bool)    { /* no-op */ }
func NotBe(cond bool) { /* no-op */ }
