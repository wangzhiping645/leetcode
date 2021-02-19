package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{"III"}, 3},
		{"12", args{"XII"}, 12},
		{"80", args{"LXXX"}, 80},
		{"199", args{"CXCIX"}, 199},
		{"1437", args{"MCDXXXVII"}, 1437},
		{"3333", args{"MMMCCCXXXIII"}, 3333},
		{"3999", args{"MMMCMXCIX"}, 3999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt2(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
