package retry

import "math"

type Interval func(attempts int) int

func Fixed(factor int) Interval {
	return func(attempts int) int {
		return factor + attempts*0
	}
}

func Exponential(base int) Interval {
	return func(attempts int) int {
		return int(math.Pow(float64(base), float64(attempts)))
	}
}

func Binary() Interval {
	return Exponential(2)
}

func Linear(scale int) Interval {
	return func(attempts int) int {
		return scale * attempts
	}
}

func Polinomial(base int) Interval {
	return func(attempts int) int {
		return int(math.Pow(float64(attempts), float64(base)))
	}
}
