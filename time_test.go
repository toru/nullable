// Copyright 2023 Toru Maesaka
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package nullable

import (
	"testing"
	"time"
)

func TestNewTime(t *testing.T) {
	testCases := []struct {
		label string
		input string
		want  bool
	}{
		{"with valid time", "2012-12-12T12:12:12Z", true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			src, err := time.Parse(time.RFC3339, tc.input)
			if err != nil {
				t.Error(err)
				return
			}

			val := NewTime(src)
			if val.Valid != tc.want {
				t.Errorf("got: %v, want: %v", val.Valid, tc.want)
				return
			}
			if val.Time != src {
				t.Errorf("got: %v, want: %v", val.Time, src)
				return
			}
		})
	}
}

func TestTimeNull(t *testing.T) {
	testCases := []struct {
		label   string
		subject Time
		want    bool
	}{
		{"with NULL time", Time{Valid: false}, true},
		{"with non NULL time", NewTime(time.Now()), false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.Null(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
			if res := tc.subject.Nil(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}
