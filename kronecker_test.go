package intalg

import "testing"

func TestKronecker(t *testing.T) {
	var tests = []struct {
		a, b int64
		r    int
	}{
		{
			a: 0,
			b: 0,
			r: 0,
		},
		{
			a: 1,
			b: 0,
			r: 1,
		},
		{
			a: -1,
			b: 0,
			r: 1,
		},
		{
			a: 3,
			b: 0,
			r: 0,
		},
		{
			a: 0,
			b: -1,
			r: 1,
		},
		{
			a: 3,
			b: -1,
			r: 1,
		},
		{
			a: -3,
			b: -1,
			r: -1,
		},
		{
			a: -2,
			b: 2,
			r: 0,
		},
		{
			a: -1,
			b: 2,
			r: 1,
		},
		{
			a: 0,
			b: 2,
			r: 0,
		},
		{
			a: 1,
			b: 2,
			r: 1,
		},
		{
			a: 2,
			b: 2,
			r: 0,
		},
		{
			a: 3,
			b: 2,
			r: -1,
		},
		{
			a: 4,
			b: 2,
			r: 0,
		},
		{
			a: 5,
			b: 2,
			r: -1,
		},
		{
			a: 9,
			b: 2,
			r: 1,
		},
		{
			a: -1,
			b: 3,
			r: -1,
		},
		{
			a: 0,
			b: 3,
			r: 0,
		},
		{
			a: 1,
			b: 3,
			r: 1,
		},
		{
			a: 2,
			b: 3,
			r: -1,
		},
		{
			a: 2,
			b: 15,
			r: 1,
		},
		{
			a: 3,
			b: 15,
			r: 0,
		},
		{
			a: 4,
			b: 15,
			r: 1,
		},
	}
	for _, test := range tests {
		r := Kronecker(test.a, test.b)
		if r != test.r {
			t.Errorf("(%v/%v) got %v, want %v", test.a, test.b, r, test.r)
		}
	}
}
