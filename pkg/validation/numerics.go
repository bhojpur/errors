package validation

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"math"
)

// Abs returns absolute value of number
func Abs(value float64) float64 {
	return math.Abs(value)
}

// Sign returns signum of number: 1 in case of value > 0, -1 in case of value < 0, 0 otherwise
func Sign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// IsNegative returns true if value < 0
func IsNegative(value float64) bool {
	return value < 0
}

// IsPositive returns true if value > 0
func IsPositive(value float64) bool {
	return value > 0
}

// IsNonNegative returns true if value >= 0
func IsNonNegative(value float64) bool {
	return value >= 0
}

// IsNonPositive returns true if value <= 0
func IsNonPositive(value float64) bool {
	return value <= 0
}

// InRangeInt returns true if value lies between left and right border
func InRangeInt(value, left, right interface{}) bool {
	value64, _ := ToInt(value)
	left64, _ := ToInt(left)
	right64, _ := ToInt(right)
	if left64 > right64 {
		left64, right64 = right64, left64
	}
	return value64 >= left64 && value64 <= right64
}

// InRangeFloat32 returns true if value lies between left and right border
func InRangeFloat32(value, left, right float32) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRangeFloat64 returns true if value lies between left and right border
func InRangeFloat64(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRange returns true if value lies between left and right border, generic type to handle int, float32, float64 and string.
// All types must the same type.
// False if value doesn't lie in range or if it incompatible or not comparable
func InRange(value interface{}, left interface{}, right interface{}) bool {
	switch value.(type) {
	case int:
		intValue, _ := ToInt(value)
		intLeft, _ := ToInt(left)
		intRight, _ := ToInt(right)
		return InRangeInt(intValue, intLeft, intRight)
	case float32, float64:
		intValue, _ := ToFloat(value)
		intLeft, _ := ToFloat(left)
		intRight, _ := ToFloat(right)
		return InRangeFloat64(intValue, intLeft, intRight)
	case string:
		return value.(string) >= left.(string) && value.(string) <= right.(string)
	default:
		return false
	}
}

// IsWhole returns true if value is whole number
func IsWhole(value float64) bool {
	return math.Remainder(value, 1) == 0
}

// IsNatural returns true if value is natural number (positive and whole)
func IsNatural(value float64) bool {
	return IsWhole(value) && IsPositive(value)
}
