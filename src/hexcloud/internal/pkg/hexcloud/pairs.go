package hexcloud

import "math"

func Pair(x, y int64) int64 {
	x = SignedTransform(x)
	y = SignedTransform(y)

	return SzudzikPair(x, y)
}

func Unpair(z int64) (x, y int64) {
	x, y = SzudzikUnPair(z)

	x = UnSignedTransform(x)
	y = UnSignedTransform(y)

	return x, y
}

func SignedTransform(x int64) int64 {
	if x >= 0 {
		return x * 2
	} else {
		return (-2 * x) - 1
	}
}

func UnSignedTransform(x int64) int64 {
	if x%2 == 0 {
		return x / 2
	} else {
		return x/-2 - 1
	}
}

func SzudzikUnPair(z int64) (x, y int64) {
	b := int64(math.Floor(math.Sqrt(float64(z))))
	a := z - b*b

	if a < b {
		x = a
		y = b
	} else {
		x = b
		y = a - b
	}

	return x, y
}

func SzudzikPair(x, y int64) int64 {
	if x >= y {
		return (x * x) + x + y
	} else {
		return (y * y) + x
	}
}
