package i8math

// Code generated by generate/generate.go from multiple.go.tmpl. DO NOT EDIT.

// RoundToMultiple returns the multiple of y closest to x, rounding ties away
// from zero.
// If x + y/2 exceeds MaxValue, or if x - y/2 is less than MinValue, then the
// result is undefined.
//
// Special cases:
//     RoundToMultiple(x, 0) = x
//     RoundToMultiple(x, -1) = x
func RoundToMultiple(x, y int8) int8 {
	if y <= 1 {
		return x
	}

	if x < 0 {
		return ((x - y/2) / y) * y
	}

	return ((x + y/2) / y) * y
}

// CeilToMultiple returns the least multiple of y that is at least x.
// If x+y exceeds MaxValue, then the result is undefined.
//
// Special cases:
//     CeilToMultiple(x, 0) = x
//     CeilToMultiple(x, -1) = x
func CeilToMultiple(x, y int8) int8 {
	if y <= 1 {
		return x
	}

	if x < 0 {
		return (x / y) * y
	}

	return ((x + y - 1) / y) * y
}

// FloorToMultiple returns the greatest multiple of y that is at most x.
// If x-y is less than MinValue, then the result is undefined.
//
// Special cases:
//     FloorToMultiple(x, 0) = x
//     FloorToMultiple(x, -1) = x
func FloorToMultiple(x, y int8) int8 {
	if y <= 1 {
		return x
	}

	if x < 0 {
		return ((x - y + 1) / y) * y
	}

	return (x / y) * y
}
