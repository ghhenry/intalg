package intalg

var kronTab1 = []int{0, 1, 0, -1, 0, -1, 0, 1}

// Kronecker returns the kronecker symbol (a/b).
func Kronecker(a, b int64) int {
	if b == 0 {
		if a == 1 || a == -1 {
			return 1
		}
		return 0
	}
	if a&1 == 0 && b&1 == 0 {
		return 0
	}
	var v int
	for b&1 == 0 {
		b >>= 1
		v++
	}
	var k int
	if v&1 == 0 {
		k = 1
	} else {
		k = kronTab1[a&7]
	}
	if b < 0 {
		b = -b
		if a < 0 {
			k = -k
		}
	}
	for {
		if a == 0 {
			if b == 1 {
				return k
			}
			return 0
		}
		v = 0
		for a&1 == 0 {
			a >>= 1
			v++
		}
		if v&1 != 0 && kronTab1[b&7] == -1 {
			k = -k
		}
		if a&b&2 != 0 {
			k = -k
		}
		r := a
		if a < 0 {
			r = -r
		}
		a, b = b%r, r
	}
}

// Kronecker64 calculates the kronecker symbol for unsigned 64 bit integers.
func Kronecker64(a, b uint64) int {
	if b == 0 {
		if a == 1 {
			return 1
		}
		return 0
	}
	if a&1 == 0 && b&1 == 0 {
		return 0
	}
	var v int
	for b&1 == 0 {
		b >>= 1
		v++
	}
	var k int
	if v&1 == 0 {
		k = 1
	} else {
		k = kronTab1[a&7]
	}
	for {
		if a == 0 {
			if b == 1 {
				return k
			}
			return 0
		}
		v = 0
		for a&1 == 0 {
			a >>= 1
			v++
		}
		if v&1 != 0 && kronTab1[b&7] == -1 {
			k = -k
		}
		if a&b&2 != 0 {
			k = -k
		}
		a, b = b%a, a
	}
}
