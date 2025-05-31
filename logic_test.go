package lrace

import (
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	type args[T any] struct {
		expr  bool
		left  T
		right T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "true",
			args: args[string]{expr: "a" == "b", left: "a", right: "b"},
			want: "b",
		},
		{
			name: "true",
			args: args[string]{expr: "a" == "a", left: "a", right: "b"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.expr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
