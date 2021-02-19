package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"3", args{3}, "III"},
		{"12", args{12}, "XII"},
		{"80", args{80}, "LXXX"},
		{"199", args{199}, "CXCIX"},
		{"1437", args{1437}, "MCDXXXVII"},
		{"3333", args{3333}, "MMMCCCXXXIII"},
		{"3999", args{3999}, "MMMCMXCIX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman1(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
