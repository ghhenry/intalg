package intalg

// GCD calculates the greatest common divisor of a and b using the binary GCD method.
func GCD(a, b uint64) uint64 {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	a, b = b, a%b
	if b == 0 {
		return a
	}
	p2 := uint(0)
	for a&1 == 0 && b&1 == 0 {
		a >>= 1
		b >>= 1
		p2++
	}
	for a&1 == 0 {
		a >>= 1
	}
	for b&1 == 0 {
		b >>= 1
	}
	for {
		var t uint64
		if a < b {
			a, b = b, a
		}
		t = (a - b) >> 1
		if t == 0 {
			return a << p2
		}
		for t&1 == 0 {
			t >>= 1
		}
		a = t
	}
}
