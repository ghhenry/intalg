package intalg

import (
	"math/big"
	"testing"
)

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

func TestEGCD(t *testing.T) {
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
		{
			a: 4294967295, b: 42949,
			d: 1,
		},
	}
	for _, test := range tests {
		d, u, v, neg := EGCD(test.a, test.b)
		if d != test.d {
			t.Errorf("gcd(%v,%v) = %v, want %v", test.a, test.b, d, test.d)
		}
		tmp1 := new(big.Int).Mul(new(big.Int).SetUint64(test.a), new(big.Int).SetUint64(u)) // a*u
		tmp2 := new(big.Int).Mul(new(big.Int).SetUint64(test.b), new(big.Int).SetUint64(v)) // b*v
		tmp1.Sub(tmp1, tmp2)                                                                // a*u-b*v
		if neg {
			tmp1.Neg(tmp1)
		}
		if tmp1.Cmp(tmp2.SetUint64(d)) != 0 {
			t.Errorf("%v*(%v)+%v*(%v) = %v", test.a, u, test.b, v, tmp1)
		}
	}
}

func BenchmarkGCD(b *testing.B) {
	var d uint64
	for i := 0; i < b.N; i++ {
		d = GCD(18446744073709551615, 493884732703523455)
	}
	b.Log(d)
}

func BenchmarkEGCD(b *testing.B) {
	var d uint64
	for i := 0; i < b.N; i++ {
		d, _, _, _ = EGCD(18446744073709551615, 493884732703523455)
	}
	b.Log(d)
}
