package math

var _ = lcm
var _ = gcdeuclid
var _ = gcdeuclidean

// using gcd
// lcm(a,b) = |ab|/gcd(a,b)
// lcm(21,6) = 21*6/gcd(21,6) = 21*6/3 = 42
func lcm(a, b int) int {
	return a * b / gcdeuclidean(a, b)
}

// gcd(48,18)
// -> gcd(48-18,18) = gcd(30,18) -> gcd(30-18,18) = gcd(12,18)
// -> gcd(12,18-12) = gcd(12,6)  -> gcd(12-6,6)   = gcd(6,6)
// so gcd(48,18)=6
func gcdeuclid(a, b int) int {
	if a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

// gcd(48,18) -> gcd(48,18 mod 48)=gcd(48,18)
// -> gcd(18,48 mod 18)=gcd(18,12)
// -> gcd(12,18 mod 12)=gcd(12,6)
// -> gcd(6,12 mod 6)=gcd(6,0)
// so gcd(48,18)=6
func gcdeuclidean(a, b int) int {
	// gcd(a,b) -> gcd(b mod a, a)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
