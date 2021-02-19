package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 4, 5, 18, 17, 6}}, 17},
		{"2", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"3", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxArea1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 4, 5, 18, 17, 6}}, 17},
		{"2", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"3", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea1(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
