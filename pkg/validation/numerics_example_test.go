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

func ExampleAbs() {
	_ = Abs(-123.3e1) // 123.3e1
	_ = Abs(+0)       // 0
	_ = Abs(321)      // 321
}

func ExampleSign() {
	_ = Sign(-123) // -1
	_ = Sign(123)  // 1
	_ = Sign(0)    // 0
}

func ExampleIsNegative() {
	_ = IsNegative(-123) // true
	_ = IsNegative(0)    // false
	_ = IsNegative(123)  // false
}

func ExampleIsPositive() {
	_ = IsPositive(-123) // false
	_ = IsPositive(0)    // false
	_ = IsPositive(123)  // true
}

func ExampleIsNonNegative() {
	_ = IsNonNegative(-123) // false
	_ = IsNonNegative(0)    // true
	_ = IsNonNegative(123)  // true
}

func ExampleIsNonPositive() {
	_ = IsNonPositive(-123) // true
	_ = IsNonPositive(0)    // true
	_ = IsNonPositive(123)  // false
}

func ExampleInRangeInt() {
	_ = InRangeInt(10, -10, 10) // true
	_ = InRangeInt(10, 10, 20)  // true
	_ = InRangeInt(10, 11, 20)  // false
}

func ExampleInRangeFloat32() {
	_ = InRangeFloat32(10.02, -10.124, 10.234) // true
	_ = InRangeFloat32(10.123, 10.123, 20.431) // true
	_ = InRangeFloat32(10.235, 11.124, 20.235) // false
}

func ExampleInRangeFloat64() {
	_ = InRangeFloat64(10.02, -10.124, 10.234) // true
	_ = InRangeFloat64(10.123, 10.123, 20.431) // true
	_ = InRangeFloat64(10.235, 11.124, 20.235) // false
}

func ExampleInRange() {
	_ = InRange(10, 11, 20)             // false
	_ = InRange(10.02, -10.124, 10.234) // true
	_ = InRange("abc", "a", "cba")      // true
}

func ExampleIsWhole() {
	_ = IsWhole(1.123) // false
	_ = IsWhole(1.0)   // true
	_ = IsWhole(10)    // true
}

func ExampleIsNatural() {
	_ = IsNatural(1.123) // false
	_ = IsNatural(1.0)   // true
	_ = IsNatural(-10)   // false
}
