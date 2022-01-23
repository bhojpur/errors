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

import "time"

func ExampleToBoolean() {
	// Returns the boolean value represented by the string.
	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	// Any other value returns an error.
	_, _ = ToBoolean("false")  // false, nil
	_, _ = ToBoolean("T")      // true, nil
	_, _ = ToBoolean("123123") // false, error
}

func ExampleToInt() {
	_, _ = ToInt(1.0)     // 1, nil
	_, _ = ToInt("-124")  // -124, nil
	_, _ = ToInt("false") // 0, error
}

func ExampleToFloat() {
	_, _ = ToFloat("-124.2e123") // -124.2e123, nil
	_, _ = ToFloat("false")      // 0, error
}

func ExampleToString() {
	_ = ToString(new(interface{}))        // 0xc000090200
	_ = ToString(time.Second + time.Hour) // 1h1s
	_ = ToString(123)                     // 123
}

func ExampleToJSON() {
	_, _ = ToJSON([]int{1, 2, 3})          // [1, 2, 3]
	_, _ = ToJSON(map[int]int{1: 2, 2: 3}) // { "1": 2, "2": 3 }
	_, _ = ToJSON(func() {})               // error
}
