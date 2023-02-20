package triangle

type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
	if !isExists(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isIsosceles(a, b, c) {
		return Iso
	}

	return Sca
}

func isExists(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0 &&
		a+b >= c &&
		a+c >= b &&
		b+c >= a
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c && a == c
}

func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || c == a
}
