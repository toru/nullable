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

import "testing"

func TestStringSet(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		input   string
		want    error
	}{
		{"with empty string", String{}, "", nil},
		{"with non-empty string", String{}, "miraculous", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			tc.subject.Set(tc.input)

			if !tc.subject.Valid {
				t.Errorf("got: %v, want: %v", tc.subject.Valid, true)
				return
			}
			if tc.subject.String != tc.input {
				t.Errorf("got: %v, want: %v", tc.subject.String, tc.input)
				return
			}
		})
	}
}

func TestStringValue(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    error
	}{
		{"with NULL string", String{Valid: false}, nil},
		{"with empty string", String{String: "", Valid: true}, nil},
		{"with non-empty string", String{String: "hello", Valid: true}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if _, err := tc.subject.Value(); err != tc.want {
				t.Errorf("got: %v, want: %v", err, tc.want)
				return
			}
		})
	}
}

func TestStringScan(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		value   string
		want    error
	}{
		{"with NULL string", String{Valid: false}, "", nil},
		{"with empty string", String{String: "", Valid: true}, "", nil},
		{"with non-empty string", String{String: "hello", Valid: true}, "", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if err := tc.subject.Scan(tc.value); err != tc.want {
				t.Errorf("got: %v, want: %v", err, tc.want)
				return
			}
		})
	}
}

func TestStringPresent(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    bool
	}{
		{"with NULL string", String{Valid: false}, false},
		{"with empty string", String{String: "", Valid: true}, false},
		{"with non-empty string", String{String: "hello", Valid: true}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.Present(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}

func TestStringNull(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    bool
	}{
		{"with NULL string", String{Valid: false}, true},
		{"with empty string", String{String: "", Valid: true}, false},
		{"with non-empty string", String{String: "hello", Valid: true}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.Null(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}

func TestStringNil(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    bool
	}{
		{"with NULL string", String{Valid: false}, true},
		{"with empty string", String{String: "", Valid: true}, false},
		{"with non-empty string", String{String: "hello", Valid: true}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.Nil(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}

func TestStringEmpty(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    bool
	}{
		{"with NULL string", String{Valid: false}, true},
		{"with empty string", String{String: "", Valid: true}, true},
		{"with non-empty string", String{String: "hello", Valid: true}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.Empty(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}

func TestStringHexString(t *testing.T) {
	testCases := []struct {
		label   string
		subject String
		want    string
	}{
		{"with NULL string", String{Valid: false}, ""},
		{"with NULL string + non-empty value", String{String: "sneaky", Valid: false}, ""},
		{"with empty string", String{String: "", Valid: true}, ""},
		{"with non-empty string", String{String: "hello", Valid: true}, "68656c6c6f"},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			if res := tc.subject.HexString(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}
