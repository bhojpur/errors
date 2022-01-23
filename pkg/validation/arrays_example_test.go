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

func ExampleFilter() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Filter(data, fn) // result = []interface{}{2, 4, 6, 8, 10}
}

func ExampleCount() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Count(data, fn) // result = 5
}

func ExampleMap() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn ResultIterator = func(value interface{}, index int) interface{} {
		return value.(int) * 3
	}
	_ = Map(data, fn) // result = []interface{}{1, 6, 9, 12, 15}
}

func ExampleEach() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn Iterator = func(value interface{}, index int) {
		println(value.(int))
	}
	Each(data, fn)
}

func ExampleFind() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int) == 4
	}
	_ = Find(data, fn) // result = 4
}
