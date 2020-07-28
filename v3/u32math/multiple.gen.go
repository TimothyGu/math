package u32math

// Code generated by generate/generate.go from multiple.go.tmpl. DO NOT EDIT.

// RoundToMultiple returns the multiple of y closest to x, rounding ties away
// from zero.
// If x + y/2 exceeds MaxValue, then the result is undefined.
//
// Special cases:
//     RoundToMultiple(x, 0) = x
func RoundToMultiple(x, y uint32) uint32 {
	if y <= 1 {
		return x
	}

	return ((x + y/2) / y) * y
}

// CeilToMultiple returns the least multiple of y that is at least x.
// If x+y exceeds MaxValue, then the result is undefined.
//
// Special cases:
//     CeilToMultiple(x, 0) = x
func CeilToMultiple(x, y uint32) uint32 {
	if y <= 1 {
		return x
	}

	return ((x + y - 1) / y) * y
}

// FloorToMultiple returns the greatest multiple of y that is at most x.
//
// Special cases:
//     FloorToMultiple(x, 0) = x
func FloorToMultiple(x, y uint32) uint32 {
	if y <= 1 {
		return x
	}

	return (x / y) * y
}
