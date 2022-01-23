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

import "testing"

func BenchmarkAbs(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = Abs(-123.3e1)
	}
}

func BenchmarkSign(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = Sign(-123.3e1)
	}
}

func BenchmarkIsNegative(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNegative(-123.3e1)
	}
}

func BenchmarkIsPositive(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsPositive(-123.3e1)
	}
}

func BenchmarkIsNonNegative(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNonNegative(-123.3e1)
	}
}

func BenchmarkIsNonPositive(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNonPositive(-123.3e1)
	}
}

func BenchmarkInRangeInt(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeInt(10, -100, 100)
	}
}

func BenchmarkInRangeFloat32(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeFloat32(10, -100, 100)
	}
}

func BenchmarkInRangeFloat64(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeFloat64(10, -100, 100)
	}
}

func BenchmarkInRange(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRange("abc", "a", "cba")
	}
}

func BenchmarkIsWhole(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsWhole(123.132)
	}
}

func BenchmarkIsNatural(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNatural(123.132)
	}
}
