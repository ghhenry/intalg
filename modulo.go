package intalg

import (
	"fmt"
	"math/bits"
)

// ModMul64 calculates a * b % q.
func ModMul64(a uint64, b uint64, q uint64) uint64 {
	// TODO: remove test output
	if a >= q {
		fmt.Printf("oops, a=%v >= q=%v\n", a, q)
	}
	if b >= q {
		fmt.Printf("oops, b=%v >= q=%v\n", b, q)
	}
	hi, lo := bits.Mul64(a, b)
	_, rem := bits.Div64(hi, lo, q)
	return rem
}

// ModPow32 calculates a^n mod q.
// if n < 0, the function panics if gcd(a,q) != 1
func ModPow32(a uint32, n int64, q uint32) uint32 {
	la := uint64(a)
	lq := uint64(q)
	if n < 0 {
		la = ModInv64(la, lq)
		if -n == n {
			// n is the smallest negative number
			la = la * la % lq
			n = -(n >> 1)
		} else {
			n = -n
		}
	}
	y := uint64(1)
	for n != 0 {
		for n&1 == 0 {
			n >>= 1
			la = la * la % lq
		}
		n--
		y = y * la % lq
	}
	return uint32(y)
}

// ModPow64 calculates a^n mod q.
// if n < 0, the function panics if gcd(a,q) != 1
func ModPow64(a uint64, n int64, q uint64) uint64 {
	if n < 0 {
		a = ModInv64(a, q)
		if -n == n {
			// n is the smallest negative number
			a = ModMul64(a, a, q)
			n = -(n >> 1)
		} else {
			n = -n
		}
	}
	y := uint64(1)
	for n != 0 {
		for n&1 == 0 {
			n >>= 1
			a = ModMul64(a, a, q)
		}
		n--
		y = ModMul64(y, a, q)
	}
	return y
}

// ModInv64 returns the modular inverse of a mod q.
// If there is no inverse, it panics.
func ModInv64(a, q uint64) uint64 {
	d, _, v, neg := EGCD(q, a)
	if d != 1 {
		panic(fmt.Errorf("%v has no inverse mod %v, common divisor %v", a, q, d))
	}
	if neg {
		return v
	}
	return q - v
}

// NonSquare finds a non-square modulo p. p must be an odd prime.
func NonSquare(p uint64) (a uint64) {
	for a = uint64(2); Kronecker64(a, p) != -1; a++ {
	}
	return a
}

// ModSqrt64 calculates a square root of a modulo p if one exists. p must be an odd prime.
func ModSqrt64(a, p uint64) (root uint64, ok bool) {
	if a == 0 {
		return 0, true
	}
	if Kronecker64(a, p) != 1 {
		return 0, false
	}
	if p&3 == 3 {
		return ModPow64(a, int64(p>>2+1), p), true
	}
	if p&7 == 5 {
		x := ModPow64(a, int64(p>>2), p)
		if x == 1 {
			return ModPow64(a, int64(p>>3+1), p), true
		}
		a2 := a << 1
		if a2 < a || a2 > p {
			a2 -= p
		}
		a4 := a2 << 1
		if a4 < a2 || a4 > p {
			a4 -= p
		}
		root = ModPow64(a4, int64(p>>3), p)
		return ModMul64(a2, root, p), true
	}
	// no luck so far, use Tonelli Shanks
	e, q := 1, p>>1
	for q&1 == 0 {
		q >>= 1
		e++
	}
	z := ModPow64(NonSquare(p), int64(q), p)
	y := z
	r := e
	x := ModPow64(a, int64(q>>1), p)
	t := ModMul64(a, x, p)
	b := ModMul64(t, x, p)
	x = t
	for {
		if b == 1 {
			return x, true
		}
		t = b
		var m int
		for m = 1; m < r; m++ {
			t = ModMul64(t, t, p)
			if t == 1 {
				break
			}
		}
		if m == r {
			return 0, false
		}
		t = y
		for i := 0; i < r-m-1; i++ {
			t = ModMul64(t, t, p)
		}
		y = ModMul64(t, t, p)
		r = m
		x = ModMul64(x, t, p)
		b = ModMul64(b, y, p)
	}
}
