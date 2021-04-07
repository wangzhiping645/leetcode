package q0863

import (
	"reflect"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

func Test_distanceK(t *testing.T) {
	type args struct {
		root   []int
		target []int
		K      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, 2}, []int{7, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(structures.Ints2TreeNode(tt.args.root),
				structures.Ints2TreeNode(tt.args.target), tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distanceK2(t *testing.T) {
	type args struct {
		root   []int
		target []int
		K      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, 2}, []int{7, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK2(structures.Ints2TreeNode(tt.args.root),
				structures.Ints2TreeNode(tt.args.target), tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
