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
		{
			arr1:   []string{"a", "b"},
			arr2:   []string{"a", "b", "d"},
			want:   false,
			strict: false,
		},
		{
			arr1:   []string{"a", "b", "c"},
			arr2:   []string{"a", "b", "d"},
			want:   false,
			strict: false,
		},
	}
	for _, item := range cases {
		give := ArrayCompare(item.arr1, item.arr2, item.strict)
		if give != item.want {
			t.Errorf("ArrayCompare(%v, %v, %v) give: %v, want: %v", item.arr1, item.arr2, item.strict, give, item.want)
		}
	}
}

func TestHasArrayPrefix(t *testing.T) {
	type args struct {
		prefix string
		arr    []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "g1",
			args: args{"g1_", []string{"a", "b", "c"}},
			want: false,
		},
		{
			name: "g2",
			args: args{"g1_", []string{"g1_a", "b", "c"}},
			want: true,
		},
		{
			name: "g3",
			args: args{"g1_", nil},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasArrayPrefix(tt.args.prefix, tt.args.arr); got != tt.want {
				t.Errorf("HasArrayPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
