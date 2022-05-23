package string_sum

import (
	"reflect"
	"testing"
)

func TestStringSum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1",
			args{"3 + 5"},
			"8",
		},
		{"t2", args{"- 3 + - 5"}, "-8"},
		{"t3", args{"- 30 - 5"}, "-35"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := StringSum(tt.args.input); err != nil || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSum() = %v, want %v, %v", got, tt.want, err)
			}
		})
	}
}
