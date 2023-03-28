package nullable

import "testing"

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
