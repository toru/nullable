package nullable

import "testing"

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
		{"with NULL string", String{Valid: false}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"with NULL string + non-empty value", String{String: "sneaky", Valid: false}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"with empty string", String{String: "", Valid: true}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"with non-empty string", String{String: "hello", Valid: true}, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
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
