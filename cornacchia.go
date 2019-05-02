package intalg

import (
	"errors"
	"math"
)

// Sqrt64 returns floor(sqrt(n)).
func Sqrt64(n uint64) uint64 {
	if n < 1<<52 {
		return uint64(math.Sqrt(float64(n)))
	}
	var y = 1 + uint64(math.Sqrt(float64(n)))
	var x = y + 1
	for y < x {
		x = y
		y = (n/x + x) / 2
	}
	return x
}

// Cornacchia64 finds an integer solution for x^2 + d*y^2 = p if one exists.
// p is a prime and 0 < d < p
func Cornacchia64(d, p uint64) (x uint64, y uint64, ok bool) {
	if p == 2 {
		if d == 1 {
			return 1, 1, true
		}
		return 0, 0, false
	}
	x0, ok := ModSqrt64(p-d, p)
	if !ok {
		return 0, 0, false
	}
	if x0 <= p>>1 {
		x0 = p - x0
	}
	a, b := p, x0
	l := Sqrt64(p)
	for b > l {
		a, b = b, a%b
	}
	c := p - b*b
	if c%d != 0 {
		return 0, 0, false
	}
	c /= d
	y = Sqrt64(c)
	if y*y != c {
		return 0, 0, false
	}
	return b, y, true
}

// Cornacchia4p64 finds an integer solution for x^2 + d*y^2 = 4p if one exists.
// p is a prime and 0 < d < 4p and -d = 0,1 mod 4.
func Cornacchia4p64(d, p uint64) (x uint64, y uint64, ok bool) {
	if p&(3<<62) != 0 {
		panic(errors.New("p is too large"))
	}
	p4 := p << 2
	if d == 0 || d >= p4 {
		return 0, 0, false
	}
	if p == 2 {
		switch d {
		case 4:
			return 2, 1, true
		case 7:
			return 1, 1, true
		default:
			return 0, 0, false
		}
	}
	x0, ok := ModSqrt64((p4-d)%p, p)
	if !ok {
		return 0, 0, false
	}
	if x0&1 != d&1 {
		x0 = p - x0
	}
	a := p << 1
	b := x0
	l := Sqrt64(p4)
	for b > l {
		a, b = b, a%b
	}
	c := p4 - b*b
	if c%d != 0 {
		return 0, 0, false
	}
	c /= d
	y = Sqrt64(c)
	if y*y != c {
		return 0, 0, false
	}
	return b, y, true
}
