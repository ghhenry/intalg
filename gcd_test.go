package intalg

import "testing"

func TestGCD(t *testing.T) {
	var tests = []struct {
		a, b, d uint64
	}{
		{
			a: 0, b: 0,
			d: 0,
		},
		{
			a: 2, b: 0,
			d: 2,
		},
		{
			a: 0, b: 2,
			d: 2,
		},
		{
			a: 6, b: 2,
			d: 2,
		},
		{
			a: 2, b: 6,
			d: 2,
		},
		{
			a: 6, b: 4,
			d: 2,
		},
		{
			a: 4, b: 6,
			d: 2,
		},
		{
			a: 14, b: 6,
			d: 2,
		},
		{
			a: 14, b: 36,
			d: 2,
		},
		{
			a: 18446744073709551615, b: 493884732703523455,
			d: 33502085,
		},
	}
	for _, test := range tests {
		d := GCD(test.a, test.b)
		if d != test.d {
			t.Errorf("gcd(%v,%v) = %v, want %v", test.a, test.b, d, test.d)
		}
	}
}
