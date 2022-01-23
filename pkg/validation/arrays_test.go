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

func TestEach(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	acc := 0
	data := []interface{}{1, 2, 3, 4, 5}
	var fn Iterator = func(value interface{}, index int) {
		acc = acc + value.(int)
	}
	Each(data, fn)
	if acc != 15 {
		t.Errorf("Expected Each(..) to be %v, got %v", 15, acc)
	}
}

func TestMap(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	data := []interface{}{1, 2, 3, 4, 5}
	var fn ResultIterator = func(value interface{}, index int) interface{} {
		return value.(int) * 3
	}
	result := Map(data, fn)
	for i, d := range result {
		if d != fn(data[i], i) {
			t.Errorf("Expected Map(..) to be %v, got %v", fn(data[i], i), d)
		}
	}
}

func TestFind(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	findElement := 96
	data := []interface{}{1, 2, 3, 4, findElement, 5}
	var fn1 ConditionIterator = func(value interface{}, index int) bool {
		return value.(int) == findElement
	}
	var fn2 ConditionIterator = func(value interface{}, index int) bool {
		value, _ = value.(string)
		return value == "govalidator"
	}
	val1 := Find(data, fn1)
	val2 := Find(data, fn2)
	if val1 != findElement {
		t.Errorf("Expected Find(..) to be %v, got %v", findElement, val1)
	}
	if val2 != nil {
		t.Errorf("Expected Find(..) to be %v, got %v", nil, val2)
	}
}

func TestFilter(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	answer := []interface{}{2, 4, 6, 8, 10}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	result := Filter(data, fn)
	for i := range result {
		if result[i] != answer[i] {
			t.Errorf("Expected Filter(..) to be %v, got %v", answer[i], result[i])
		}
	}
}

func TestCount(t *testing.T) {
	// TODO Maybe refactor?
	t.Parallel()
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	count := 5
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	result := Count(data, fn)
	if result != count {
		t.Errorf("Expected Count(..) to be %v, got %v", count, result)
	}
}
