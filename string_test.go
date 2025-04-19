package lrace

import "testing"

func TestStringSplitWith(t *testing.T) {
	type testCase struct {
		str     string
		sep     []string
		split   []string
		exclude []string
	}

	cases := []testCase{
		{
			str:     "a.b-c.d",
			sep:     []string{".", "-"},
			split:   []string{"a", "b-c", "d"},
			exclude: []string{"b-c"},
		},
		{
			str:     "a.b-c@d",
			sep:     []string{".", "-", "@"},
			split:   []string{"a", "b", "c", "d"},
			exclude: []string{},
		},
	}
	for _, item := range cases {
		give := StringSplitWith(item.str, item.sep, item.exclude)
		if !ArrayCompare(item.split, give, true) {
			t.Errorf("SplitWith(%s, %v) give: %v, want: %v", item.str, item.sep, give, item.split)
		}
	}
}
