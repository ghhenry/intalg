package intalg

// GCD calculates the greatest common divisor of a and b using the binary GCD method.
func GCD(a, b uint64) uint64 {
	// make sure that a >= b
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	// one classical reduction step to optimize the case that a and b have very different length
	a, b = b, a%b
	if b == 0 {
		return a
	}
	// deal with factors of 2
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
	// here a and b are both odd
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

// func check(msg string, u, uo, v, vo uint64, a0, b0, a, b uint64, neg bool) {
// 	var t1, t2, t3, t4 big.Int
// 	fmt.Printf("\n%s\n", msg)
// 	fmt.Printf("%20v %20v %20v %20v\n", u, v, a0, a)
// 	fmt.Printf("%20v %20v %20v %20v\n", uo, vo, b0, b)
// 	t1.SetUint64(u)  // u
// 	t2.SetUint64(a0) // a0
// 	t1.Mul(&t1, &t2) // u*a0
// 	t3.SetUint64(b0) // b0
// 	t4.SetUint64(v)  // v
// 	t4.Mul(&t3, &t4) // v*b0
// 	if neg {
// 		t1.Sub(&t4, &t1) // -u*a0 + v*b0
// 	} else {
// 		t1.Sub(&t1, &t4) // u*a0 - v*b0
// 	}
// 	t4.SetUint64(a) // a
// 	if t4.Cmp(&t1) != 0 {
// 		fmt.Println("a incorrect")
// 	}
// 	t1.SetUint64(uo) // uo
// 	t1.Mul(&t1, &t2) // uo*a0
// 	t4.SetUint64(vo) // vo
// 	t4.Mul(&t3, &t4) // vo*b0
// 	if neg {
// 		t1.Sub(&t1, &t4) // uo*a0 - vo*b0
// 	} else {
// 		t1.Sub(&t4, &t1) // -uo*a0 + vo*b0
// 	}
// 	t4.SetUint64(b) // b
// 	if t4.Cmp(&t1) != 0 {
// 		fmt.Println("b incorrect")
// 	}
// }

// EGCD calculates the greatest common divisor of a and b using the binary GCD method.
// It also returns numbers u and v such that a*u - b*v = +-d.
// If neg is true the sign is -d, otherwise +d.
func EGCD(a, b uint64) (d uint64, u, v uint64, neg bool) {
	var uc, vc uint64
	u, v = 1, 0
	var uo, vo uint64 = 0, 1
	a0, b0 := a, b
	// make sure that a >= b
	if a < b {
		a, b = b, a
		u, uo = uo, u
		v, vo = vo, v
		neg = !neg
		//check("exchange", u, uo, v, vo, b, a, a, b, neg)
	}
	if b == 0 {
		return a, u, v, neg
	}
	// common factors of 2
	p2 := uint(0)
	for a&1 == 0 && b&1 == 0 {
		a >>= 1
		b >>= 1
		p2++
	}
	a0 >>= p2
	b0 >>= p2
	uc = b0>>1 + b0&1
	vc = a0>>1 + a0&1
	//check("setup", u, uo, v, vo, a0, b0, a, b, neg)
	// one classical reduction step to optimize the case that a and b have very different length
	q := a / b
	a, b = b, a-q*b
	u, uo = uo, u+q*uo
	v, vo = vo, v+q*vo
	neg = !neg
	//check("classic step", u, uo, v, vo, a0, b0, a, b, neg)
	if b == 0 {
		return a << p2, u, v, neg
	}
	// at least one of a,b is odd here, make sure b is odd
	if b&1 == 0 {
		a, b = b, a
		u, uo = uo, u
		v, vo = vo, v
		neg = !neg
	}
	for {
		for a&1 == 0 {
			a >>= 1
			if u&1 == 1 || v&1 == 1 {
				u = u>>1 + uc
				v = v>>1 + vc
			} else {
				u >>= 1
				v >>= 1
			}
			//check("a even", u, uo, v, vo, a0, b0, a, b, neg)
		}
		// a and b are both odd here
		var t uint64
		if a < b {
			a, b = b, a
			u, uo = uo, u
			v, vo = vo, v
			neg = !neg
			//check("exchange", u, uo, v, vo, a0, b0, a, b, neg)
		}
		t = a - b
		if t == 0 {
			return a << p2, u, v, neg
		}
		a = t
		u += uo
		v += vo
		// fix possible overflow
		if u < uo || v < vo {
			u -= b0
			v -= a0
		}
		//check("sub", u, uo, v, vo, a0, b0, a, b, neg)
	}
}
