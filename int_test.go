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
		})
	}
}

func TestInt64Nil(t *testing.T) {
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
			if res := tc.subject.Nil(); res != tc.want {
				t.Errorf("got: %v, want: %v", res, tc.want)
				return
			}
		})
	}
}
