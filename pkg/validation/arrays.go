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

// Iterator is the function that accepts element of slice/array and its index
type Iterator func(interface{}, int)

// ResultIterator is the function that accepts element of slice/array and its index and returns any result
type ResultIterator func(interface{}, int) interface{}

// ConditionIterator is the function that accepts element of slice/array and its index and returns boolean
type ConditionIterator func(interface{}, int) bool

// ReduceIterator is the function that accepts two element of slice/array and returns result of merging those values
type ReduceIterator func(interface{}, interface{}) interface{}

// Some validates that any item of array corresponds to ConditionIterator. Returns boolean.
func Some(array []interface{}, iterator ConditionIterator) bool {
	res := false
	for index, data := range array {
		res = res || iterator(data, index)
	}
	return res
}

// Every validates that every item of array corresponds to ConditionIterator. Returns boolean.
func Every(array []interface{}, iterator ConditionIterator) bool {
	res := true
	for index, data := range array {
		res = res && iterator(data, index)
	}
	return res
}

// Reduce boils down a list of values into a single value by ReduceIterator
func Reduce(array []interface{}, iterator ReduceIterator, initialValue interface{}) interface{} {
	for _, data := range array {
		initialValue = iterator(initialValue, data)
	}
	return initialValue
}

// Each iterates over the slice and apply Iterator to every item
func Each(array []interface{}, iterator Iterator) {
	for index, data := range array {
		iterator(data, index)
	}
}

// Map iterates over the slice and apply ResultIterator to every item. Returns new slice as a result.
func Map(array []interface{}, iterator ResultIterator) []interface{} {
	var result = make([]interface{}, len(array))
	for index, data := range array {
		result[index] = iterator(data, index)
	}
	return result
}

// Find iterates over the slice and apply ConditionIterator to every item. Returns first item that meet ConditionIterator or nil otherwise.
func Find(array []interface{}, iterator ConditionIterator) interface{} {
	for index, data := range array {
		if iterator(data, index) {
			return data
		}
	}
	return nil
}

// Filter iterates over the slice and apply ConditionIterator to every item. Returns new slice.
func Filter(array []interface{}, iterator ConditionIterator) []interface{} {
	var result = make([]interface{}, 0)
	for index, data := range array {
		if iterator(data, index) {
			result = append(result, data)
		}
	}
	return result
}

// Count iterates over the slice and apply ConditionIterator to every item. Returns count of items that meets ConditionIterator.
func Count(array []interface{}, iterator ConditionIterator) int {
	count := 0
	for index, data := range array {
		if iterator(data, index) {
			count = count + 1
		}
	}
	return count
}
