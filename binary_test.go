// Copyright 2026 Toru Maesaka
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
	"slices"
	"testing"
)

func TestNewBinary(t *testing.T) {
	testCases := []struct {
		label string
		input []byte
		want  bool
	}{
		{"with valid bytes", []byte("x"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {

			val := NewBinary(tc.input)

			if val.Valid != tc.want {
				t.Errorf("got: %v, want: %v", val.Valid, tc.want)
				return
			}

			if !slices.Equal(val.Bytes, tc.input) {
				t.Errorf("got: %v, want: %v", val.Bytes, tc.input)
				return
			}
		})
	}
}

func TestBinaryValue(t *testing.T) {
	testCases := []struct {
		label   string
		subject Binary
		want    error
	}{
		{"with NULL binary", Binary{Valid: false}, nil},
		{"with empty bytes", Binary{Bytes: []byte{}, Valid: true}, nil},
		{"with non-empty bytes", Binary{Bytes: []byte("hello"), Valid: true}, nil},
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

func TestBinaryScan(t *testing.T) {
	testCases := []struct {
		label   string
		input   any
		wantErr bool
		wantVal bool
	}{
		{"with nil", nil, false, false},
		{"with empty bytes", []byte{}, false, true},
		{"with non-empty bytes", []byte("hello"), false, true},
		{"with invalid type", "bad apple", true, false},
	}

	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			var b Binary

			err := b.Scan(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("got: %v, want: %v", err, tc.wantErr)
				return
			}

			if b.Valid != tc.wantVal {
				t.Errorf("got: %v, want: %v", b.Valid, tc.wantVal)
				return
			}
		})
	}
}

func TestBinaryScanCopiesBytes(t *testing.T) {
	var b Binary

	src := []byte("hello")

	if err := b.Scan(src); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	src[0] = 'x'

	if b.Bytes[0] == 'x' {
		t.Error("original bytes still referenced")
		return
	}
}

func TestBinaryPresent(t *testing.T) {
	testCases := []struct {
		label   string
		subject Binary
		want    bool
	}{
		{"with NULL binary", Binary{Valid: false}, false},
		{"with empty binary", Binary{Bytes: []byte(""), Valid: true}, false},
		{"with nil bytes", Binary{Bytes: nil, Valid: true}, false},
		{"with non-empty binary", Binary{Bytes: []byte("hello"), Valid: true}, true},
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

func TestBinaryNull(t *testing.T) {
	testCases := []struct {
		label   string
		subject Binary
		want    bool
	}{
		{"with NULL binary", Binary{Valid: false}, true},
		{"with non-NULL binary", Binary{Bytes: []byte("hello"), Valid: true}, false},
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

func TestBinaryHexString(t *testing.T) {
	testCases := []struct {
		label   string
		subject Binary
		want    string
	}{
		{"with NULL binary", Binary{Valid: false}, ""},
		{"with NULL binary + non-empty bytes", Binary{Bytes: []byte("sneaky"), Valid: false}, ""},
		{"with empty bytes", Binary{Bytes: []byte{}, Valid: true}, ""},
		{"with non-empty bytes", Binary{Bytes: []byte("hello"), Valid: true}, "68656c6c6f"},
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
