package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question226 struct {
	para226
	ans226
}

// para 是参数
// one 代表第一个参数
type para226 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans226 struct {
	one []int
}

func Test_Problem226(t *testing.T) {

	qs := []question226{

		{
			para226{[]int{}},
			ans226{[]int{}},
		},

		{
			para226{[]int{1}},
			ans226{[]int{1}},
		},

		{
			para226{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans226{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 226------------------------\n")

	for _, q := range qs {
		_, p := q.ans226, q.para226
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(invertTree(root)))
	}
	fmt.Printf("\n\n\n")
}

func Test_distanceK(t *testing.T) {
	type args struct {
		root []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{4, 2, 7, 1, 3, 6, 9}, 0}, []int{4}},
		{"2", args{[]int{4, 2, 7, 1, 3, 6, 9}, 1}, []int{2, 7}},
		{"3", args{[]int{4, 2, 7, 1, 3, 6}, 2}, []int{1, 3, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result = []int{}
			if got := distanceK(structures.Ints2TreeNode(tt.args.root), tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
