package intalg

import "testing"

func TestModMul64(t *testing.T) {
	var tests = []struct {
		a, b, q, r uint64
	}{
		{a: 0, b: 2, q: 3, r: 0},
		{a: 1, b: 2, q: 3, r: 2},
		{a: 2, b: 2, q: 3, r: 1},
		{a: 18446744073709551592, b: 18446744073709551508, q: 18446744073709551615, r: 2461},
	}
	for _, test := range tests {
		r := ModMul64(test.a, test.b, test.q)
		if r != test.r {
			t.Errorf("%v * %v mod %v: got %v, want %v", test.a, test.b, test.q, r, test.r)
		}
	}
}

func TestModPow32(t *testing.T) {
	var tests = []struct {
		a, q, r uint32
		n       int64
	}{
		{a: 0, n: 2, q: 3, r: 0},
		{a: 1, n: 2, q: 3, r: 1},
		{a: 2, n: 2, q: 3, r: 1},
		{a: 42949, n: 27, q: 4294967295, r: 2488881274},
		{a: 42949, n: -27, q: 4294967295, r: 4194068974},
		{a: 42949, n: -1, q: 4294967295, r: 278204359},
		{a: 42949, n: -9223372036854775808, q: 4294967295, r: 1},
	}
	for _, test := range tests {
		r := ModPow32(test.a, test.n, test.q)
		if r != test.r {
			t.Errorf("%v ^ %v mod %v: got %v, want %v", test.a, test.n, test.q, r, test.r)
		}
	}
}

func TestModPow64(t *testing.T) {
	var tests = []struct {
		a, q, r uint64
		n       int64
	}{
		{a: 0, n: 2, q: 3, r: 0},
		{a: 1, n: 2, q: 3, r: 1},
		{a: 2, n: 2, q: 3, r: 1},
		{a: 42949, n: 27, q: 4294967295, r: 2488881274},
		{a: 42949, n: -27, q: 4294967295, r: 4194068974},
		{a: 42949, n: -1, q: 4294967295, r: 278204359},
		{a: 42949, n: -9223372036854775808, q: 4294967295, r: 1},
		{a: 42949, n: 107, q: 18446744073709551615, r: 13247271014208205834},
		{a: 42949, n: -107, q: 18446744073709551615, r: 8465843072778963274},
	}
	for _, test := range tests {
		r := ModPow64(test.a, test.n, test.q)
		if r != test.r {
			t.Errorf("%v ^ %v mod %v: got %v, want %v", test.a, test.n, test.q, r, test.r)
		}
	}
}

func TestModSqrt(t *testing.T) {
	var tests = []struct {
		a, p uint64
		ok   bool
	}{
		{a: 0, p: 3, ok: true},
		{a: 1, p: 3, ok: true},
		{a: 2, p: 3, ok: false},
		{a: 0, p: 5, ok: true},
		{a: 1, p: 5, ok: true},
		{a: 2, p: 5, ok: false},
		{a: 3, p: 5, ok: false},
		{a: 4, p: 5, ok: true},
		{a: 0, p: 17, ok: true},
		{a: 1, p: 17, ok: true},
		{a: 2, p: 17, ok: true},
		{a: 3, p: 17, ok: false},
		{a: 4, p: 17, ok: true},
		{a: 5, p: 17, ok: false},
		{a: 6, p: 17, ok: false},
		{a: 7, p: 17, ok: false},
		{a: 8, p: 17, ok: true},
		{a: 9, p: 17, ok: true},
		{a: 13157734221201575605, p: 18446744073709551557, ok: true}, // p is largest 64 bit prime
		{a: 12816110470792687436, p: 18446744073709551521, ok: true}, // p is largest 64 bit prime which uses Tonelli Shanks
	}
	for _, test := range tests {
		r, ok := ModSqrt64(test.a, test.p)
		if ok != test.ok {
			t.Errorf("sqrt(%v) mod %v, existence wrong", test.a, test.p)
			continue
		}
		if !ok {
			continue
		}
		if s := ModMul64(r, r, test.p); s != test.a {
			t.Errorf("sqrt(%v) = %v mod %v is wrong, r^2 = %v", test.a, r, test.p, s)
		}
	}
}
