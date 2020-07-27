package i8math

// Code generated by generate/generate.go from math_test.go.tmpl. DO NOT EDIT.

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TimothyGu/math/u8math"
)

func ExampleAbs() {
	fmt.Println(Abs(0))
	fmt.Println(Abs(1))
	fmt.Println(Abs(-1))
	fmt.Println(Abs(MaxValue) == MaxValue)
	fmt.Println(Abs(MinValue) == MaxValue)
	// Output:
	// 0
	// 1
	// 1
	// true
	// true
}

func TestDim(t *testing.T) {
	assert.Equal(t, uint8(0), Dim(0, 0))
	assert.Equal(t, uint8(0), Dim(0, 1))
	assert.Equal(t, uint8(1), Dim(1, 0))
	assert.Equal(t, uint8(0), Dim(1, 1))

	assert.Equal(t, uint8(2), Dim(1, -1))
	assert.Equal(t, uint8(0), Dim(-1, 1))
	assert.Equal(t, uint8(0), Dim(-1, -1))

	assert.Equal(t, uint8(0), Dim(MaxValue, MaxValue))
	assert.Equal(t, uint8(0), Dim(0, MaxValue))
	assert.Equal(t, uint8(MaxValue), Dim(MaxValue, 0))
	assert.Equal(t, uint8(0), Dim(MaxValue-1, MaxValue))
	assert.Equal(t, uint8(1), Dim(MaxValue, MaxValue-1))

	assert.Equal(t, uint8(0), Dim(MinValue, MaxValue))
	assert.Equal(t, uint8(u8math.MaxValue), Dim(MaxValue, MinValue))

}

func ExampleMax() {
	fmt.Println(Max(0, 0))
	fmt.Println(Max(0, 1))
	fmt.Println(Max(1, 0))
	fmt.Println(Max(MaxValue, 0) == MaxValue)
	fmt.Println(Max(MinValue, MaxValue) == MaxValue)
	// Output:
	// 0
	// 1
	// 1
	// true
	// true
}

func TestMax(t *testing.T) {
	assert.Equal(t, int8(0), Max(0, 0))
	assert.Equal(t, int8(1), Max(0, 1))
	assert.Equal(t, int8(1), Max(1, 0))
	assert.Equal(t, int8(1), Max(1, 1))

	assert.Equal(t, int8(1), Max(1, -1))
	assert.Equal(t, int8(1), Max(-1, 1))
	assert.Equal(t, int8(-1), Max(-1, -1))

	assert.Equal(t, int8(MaxValue), Max(MaxValue, MaxValue))
	assert.Equal(t, int8(MaxValue), Max(0, MaxValue))
	assert.Equal(t, int8(MaxValue), Max(MaxValue, 0))
	assert.Equal(t, int8(MaxValue), Max(MaxValue-1, MaxValue))
	assert.Equal(t, int8(MaxValue), Max(MaxValue, MaxValue-1))

	assert.Equal(t, int8(MaxValue), Max(MinValue, MaxValue))
	assert.Equal(t, int8(MaxValue), Max(MaxValue, MinValue))

}

func ExampleMin() {
	fmt.Println(Min(1, 1))
	fmt.Println(Min(0, 1))
	fmt.Println(Min(1, 0))
	fmt.Println(Min(MinValue, 0) == MinValue)
	fmt.Println(Min(MinValue, MaxValue) == MinValue)
	// Output:
	// 1
	// 0
	// 0
	// true
	// true
}

func TestMin(t *testing.T) {
	assert.Equal(t, int8(0), Min(0, 0))
	assert.Equal(t, int8(0), Min(0, 1))
	assert.Equal(t, int8(0), Min(1, 0))
	assert.Equal(t, int8(1), Min(1, 1))

	assert.Equal(t, int8(-1), Min(1, -1))
	assert.Equal(t, int8(-1), Min(-1, 1))
	assert.Equal(t, int8(-1), Min(-1, -1))

	assert.Equal(t, int8(MaxValue), Min(MaxValue, MaxValue))
	assert.Equal(t, int8(0), Min(0, MaxValue))
	assert.Equal(t, int8(0), Min(MaxValue, 0))

	assert.Equal(t, int8(MinValue), Min(MinValue, MaxValue))
	assert.Equal(t, int8(MinValue), Min(MaxValue, MinValue))

}
