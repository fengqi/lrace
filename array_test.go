package lrace

import "testing"

func TestArrayCompare(t *testing.T) {
	type strCase struct {
		arr1   []string
		arr2   []string
		want   bool
		strict bool
	}

	cases := []strCase{
		{
			arr1:   []string{"a", "b", "c"},
			arr2:   []string{"a", "b", "c"},
			want:   true,
			strict: false,
		},
		{
			arr1:   []string{"b", "a", "c"},
			arr2:   []string{"c", "b", "a"},
			want:   true,
			strict: false,
		},
		{
			arr1:   []string{"b", "a", "c"},
			arr2:   []string{"c", "b", "a"},
			want:   false,
			strict: true,
		},
		{
			arr1:   []string{"a", "b", "c"},
			arr2:   []string{"a", "b", "d"},
			want:   false,
			strict: true,
		},
	}
	for _, item := range cases {
		give := ArrayCompare(item.arr1, item.arr2, item.strict)
		if give != item.want {
			t.Errorf("ArrayCompare(%v, %v, %v) give: %v, want: %v", item.arr1, item.arr2, item.strict, give, item.want)
		}
	}
}
