package intalg

import (
	"fmt"
	"testing"
)

func TestNaiveP(t *testing.T) {
	t.Skip()
	p := uint64(18446744073709551557)
	d := uint64(78647315)
	for {
		fmt.Println("testing", d)
		maxy := Sqrt64(p / d)
		for y := uint64(1); y <= maxy; y++ {
			r := d * y * y
			l := p - r
			x := Sqrt64(l)
			if x*x == l {
				fmt.Println("solution", x, y)
				return
			}
		}
		d++
	}
}

func TestNaive4P(t *testing.T) {
	t.Skip()
	p := uint64(4611686018427387847)
	d := uint64(667435272)
	for {
		fmt.Println("testing", d)
		maxy := Sqrt64(4 * p / d)
		for y := uint64(1); y <= maxy; y++ {
			r := d * y * y
			l := 4*p - r
			x := Sqrt64(l)
			if x*x == l {
				fmt.Println("solution", x, y)
				return
			}
		}
		switch d & 3 {
		case 0:
			d += 3
		case 1:
			d += 2
		case 2:
			d++
		case 3:
			d++
		}
	}
}

func TestCornacchia64(t *testing.T) {
	var tests = []struct {
		d, p uint64
		ok   bool
	}{
		{d: 1, p: 2, ok: true},
		{d: 1, p: 3, ok: false},
		{d: 2, p: 3, ok: true},
		{d: 1, p: 5, ok: true},
		{d: 2, p: 5, ok: false},
		{d: 3, p: 5, ok: false},
		{d: 4, p: 5, ok: true},
		{d: 1, p: 7, ok: false},
		{d: 2, p: 7, ok: false},
		{d: 3, p: 7, ok: true},
		{d: 4, p: 7, ok: false},
		{d: 5, p: 7, ok: false},
		{d: 6, p: 7, ok: true},
		{d: 1, p: 18446744073709551557, ok: true},
		{d: 2, p: 18446744073709551557, ok: false},
		{d: 3, p: 18446744073709551557, ok: false},
		{d: 4, p: 18446744073709551557, ok: true},
		{d: 6, p: 18446744073709551557, ok: false},
		{d: 9, p: 18446744073709551557, ok: false},
		{d: 10, p: 18446744073709551557, ok: false},
		{d: 11, p: 18446744073709551557, ok: false},
		{d: 13, p: 18446744073709551557, ok: true},
		{d: 78618724, p: 18446744073709551557, ok: false},
		{d: 78618727, p: 18446744073709551557, ok: false},
		{d: 78618728, p: 18446744073709551557, ok: false},
		{d: 78618733, p: 18446744073709551557, ok: false},
		{d: 78618734, p: 18446744073709551557, ok: false},
		{d: 78618735, p: 18446744073709551557, ok: false},
		{d: 78618737, p: 18446744073709551557, ok: false},
		{d: 78618739, p: 18446744073709551557, ok: false},
		{d: 78648772, p: 18446744073709551557, ok: true},
	}
	for _, test := range tests {
		x, y, ok := Cornacchia64(test.d, test.p)
		if ok != test.ok {
			t.Errorf("d=%v, p=%v: wrong existence", test.d, test.p)
			continue
		}
		if !ok {
			continue
		}
		t.Logf("d=%v, p=%v: x=%v, y=%v", test.d, test.p, x, y)
		if x*x+test.d*y*y != test.p {
			t.Errorf("d=%v, p=%v: wrong solution x=%v, y=%v", test.d, test.p, x, y)
		}
	}
}

func TestCornacchia4p64(t *testing.T) {
	var tests = []struct {
		d, p uint64
		ok   bool
	}{
		{d: 3, p: 2, ok: false},
		{d: 4, p: 2, ok: true},
		{d: 7, p: 2, ok: true},
		{d: 3, p: 3, ok: true},
		{d: 4, p: 3, ok: false},
		{d: 7, p: 3, ok: false},
		{d: 8, p: 3, ok: true},
		{d: 11, p: 3, ok: true},
		{d: 3, p: 4611686018427387847, ok: true},
		{d: 4, p: 4611686018427387847, ok: false},
		{d: 7, p: 4611686018427387847, ok: false},
		{d: 8, p: 4611686018427387847, ok: false},
		{d: 11, p: 4611686018427387847, ok: false},
		{d: 12, p: 4611686018427387847, ok: true},
		{d: 66271, p: 4611686018427387847, ok: true},
		{d: 67107, p: 4611686018427387847, ok: true},
		{d: 667589016, p: 4611686018427387847, ok: true},
	}
	for _, test := range tests {
		x, y, ok := Cornacchia4p64(test.d, test.p)
		if ok != test.ok {
			t.Errorf("d=%v, p=%v: wrong existence", test.d, test.p)
			continue
		}
		if !ok {
			continue
		}
		t.Logf("d=%v, p=%v: x=%v, y=%v", test.d, test.p, x, y)
		if x*x+test.d*y*y != 4*test.p {
			t.Errorf("d=%v, p=%v: wrong solution x=%v, y=%v", test.d, test.p, x, y)
		}
	}
}
