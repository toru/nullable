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

func TestNewInt64(t *testing.T) {
	testCases := []struct {
		label string
		input int64
		want  bool
	}{
		{"with negative integer", -1, true},
		{"with positive integer", 1, true},
		{"with zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			val := NewInt64(tc.input)

			if val.Valid != tc.want {
				t.Errorf("got: %v, want: %v", val.Valid, tc.want)
				return
			}
			if val.Int64 != tc.input {
				t.Errorf("got: %v, want: %v", val.Int64, tc.input)
				return
			}
		})
	}
}

func TestInt64Scan(t *testing.T) {
	testCases := []struct {
		label       string
		input       int64
		wantSuccess bool
	}{
		{"with valid integer", 1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			var val Int64
			if err := val.Scan(tc.input); (err == nil) != tc.wantSuccess {
				t.Error(err)
				return
			}
		})
	}
}

func TestInt64Null(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int64
		want    bool
	}{
		{"with NULL integer", Int64{Valid: false}, true},
		{"with negative integer", Int64{Int64: -1, Valid: true}, false},
		{"with positive integer", Int64{Int64: 1, Valid: true}, false},
		{"with zero", Int64{Int64: 0, Valid: true}, false},
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

func TestNewInt32(t *testing.T) {
	testCases := []struct {
		label string
		input int32
		want  bool
	}{
		{"with negative integer", -1, true},
		{"with positive integer", 1, true},
		{"with zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			val := NewInt32(tc.input)

			if val.Valid != tc.want {
				t.Errorf("got: %v, want: %v", val.Valid, tc.want)
				return
			}
			if val.Int32 != tc.input {
				t.Errorf("got: %v, want: %v", val.Int32, tc.input)
				return
			}
		})
	}
}

func TestInt32Scan(t *testing.T) {
	testCases := []struct {
		label       string
		input       any
		wantSuccess bool
	}{
		{"with valid integer", 1, true},
		{"with overflow integer", 4294967296, false},
		{"with underflow integer", -4294967296, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			var val Int32
			if err := val.Scan(tc.input); (err == nil) != tc.wantSuccess {
				t.Error(err)
				return
			}
		})
	}
}

func TestInt64HexString(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int64
		want    string
	}{
		{"with NULL integer", Int64{Valid: false}, ""},
		{"with non-null integer", Int64{Int64: 128, Valid: true}, "80"},
		{"with another non-null integer", Int64{Int64: 12345, Valid: true}, "3039"},
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

func TestInt32Null(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int32
		want    bool
	}{
		{"with NULL integer", Int32{Valid: false}, true},
		{"with negative integer", Int32{Int32: -1, Valid: true}, false},
		{"with positive integer", Int32{Int32: 1, Valid: true}, false},
		{"with zero", Int32{Int32: 0, Valid: true}, false},
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

func TestInt32HexString(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int32
		want    string
	}{
		{"with NULL integer", Int32{Valid: false}, ""},
		{"with non-null integer", Int32{Int32: 128, Valid: true}, "80"},
		{"with another non-null integer", Int32{Int32: 12345, Valid: true}, "3039"},
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

func TestNewInt16(t *testing.T) {
	testCases := []struct {
		label string
		input int16
		want  bool
	}{
		{"with negative integer", -1, true},
		{"with positive integer", 1, true},
		{"with zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			val := NewInt16(tc.input)

			if val.Valid != tc.want {
				t.Errorf("got: %v, want: %v", val.Valid, tc.want)
				return
			}
			if val.Int16 != tc.input {
				t.Errorf("got: %v, want: %v", val.Int16, tc.input)
				return
			}
		})
	}
}

func TestInt16Scan(t *testing.T) {
	testCases := []struct {
		label       string
		input       any
		wantSuccess bool
	}{
		{"with valid integer", 1, true},
		{"with overflow integer", 65536, false},
		{"with underflow integer", -65536, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			var val Int16
			if err := val.Scan(tc.input); (err == nil) != tc.wantSuccess {
				t.Error(err)
				return
			}
		})
	}
}

func TestInt16Null(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int16
		want    bool
	}{
		{"with NULL integer", Int16{Valid: false}, true},
		{"with negative integer", Int16{Int16: -1, Valid: true}, false},
		{"with positive integer", Int16{Int16: 1, Valid: true}, false},
		{"with zero", Int16{Int16: 0, Valid: true}, false},
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

func TestInt16HexString(t *testing.T) {
	testCases := []struct {
		label   string
		subject Int16
		want    string
	}{
		{"with NULL integer", Int16{Valid: false}, ""},
		{"with non-null integer", Int16{Int16: 128, Valid: true}, "80"},
		{"with another non-null integer", Int16{Int16: 12345, Valid: true}, "3039"},
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
