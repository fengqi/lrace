package lrace

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	type args struct {
		m1 map[string]string
		m2 map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "merge",
			args: args{m1: map[string]string{"1": "111", "2": "222"}, m2: map[string]string{"2": "222"}},
			want: map[string]string{"1": "111", "2": "222"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
