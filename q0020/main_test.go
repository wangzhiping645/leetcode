package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"()[]{}"}, true},
		{"2", args{"(]"}, false},
		{"3", args{"({[]})"}, true},
		{"4", args{"(){[({[]})]}"}, true},
		{"5", args{"((([[[{{{"}, false},
		{"6", args{"(())]]"}, false},
		{"7", args{""}, true},
		{"8", args{"["}, false},
		{"9", args{"(("}, false},
		{"10", args{"(()("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
