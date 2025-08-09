//go:build assert

package should

func Be(cond bool) {
	if cond {
		return
	}

	panic("ASSERTION")
}

func NotBe(cond bool) {
	if !cond {
		return
	}

	panic("ASSERTION")
}