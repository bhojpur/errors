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
	"fmt"
	"testing"
)

func TestErrorsToString(t *testing.T) {
	t.Parallel()
	customErr := &Error{Name: "Custom Error Name", Err: fmt.Errorf("stdlib error")}
	customErrWithCustomErrorMessage := &Error{Name: "Custom Error Name 2", Err: fmt.Errorf("Bad stuff happened"), CustomErrorMessageExists: true}

	var tests = []struct {
		param1   Errors
		expected string
	}{
		{Errors{}, ""},
		{Errors{fmt.Errorf("Error 1")}, "Error 1"},
		{Errors{fmt.Errorf("Error 1"), fmt.Errorf("Error 2")}, "Error 1;Error 2"},
		{Errors{customErr, fmt.Errorf("Error 2")}, "Custom Error Name: stdlib error;Error 2"},
		{Errors{fmt.Errorf("Error 123"), customErrWithCustomErrorMessage}, "Bad stuff happened;Error 123"},
	}
	for _, test := range tests {
		actual := test.param1.Error()
		if actual != test.expected {
			t.Errorf("Expected Error() to return '%v', got '%v'", test.expected, actual)
		}
	}
}
