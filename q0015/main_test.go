package main

import (
	"fmt"
	"reflect"
	"testing"
)

//func Test_threeSum(t *testing.T) {
//	type args struct {
//		nums []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want [][]int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("threeSum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

type question15 struct {
	para15
	ans15
}

// para 是参数
// one 代表第一个参数
type para15 struct {
	a []int
}

// ans 是答案
// one 代表第一个答案
type ans15 struct {
	one [][]int
}

func Test_threeSum(t *testing.T) {

	qs := []question15{

		{
			para15{[]int{0, 0, 0}},
			ans15{[][]int{{0, 0, 0}}},
		},

		{
			para15{[]int{-1, 0, 1, 2, -1, -4}},
			ans15{[][]int{{-1, 0, 1}, {-1, -1, 2}}},
		},

		{
			para15{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			ans15{[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		},

		{
			para15{[]int{5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5}},
			ans15{[][]int{{-10, 2, 8}, {-10, 3, 7}, {-10, 4, 6}, {-10, 5, 5}, {-8, 2, 6}, {-8, 3, 5}, {-8, 4, 4}, {-7, -1, 8},
				{-7, 2, 5}, {-7, 3, 4}, {-6, -2, 8}, {-6, -1, 7}, {-6, 2, 4}, {-5, -3, 8}, {-5, -2, 7}, {-5, -1, 6}, {-5, 2, 3},
				{-4, -3, 7}, {-4, -2, 6}, {-4, -1, 5}, {-4, 2, 2}, {-3, -3, 6}, {-3, -2, 5}, {-3, -1, 4}, {-2, -1, 3}}},
		},
		{
			para15{[]int{0, 0, 0, 0}},
			ans15{[][]int{{0, 0, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 15------------------------\n")

	for _, q := range qs {
		_, p := q.ans15, q.para15
		got := threeSum(p.a)
		fmt.Printf("【input】:%v 【output】:%v  【len(output)】:%d  【len(answer)】:%d \n", p, got, len(got), len(q.one))
		if !reflect.DeepEqual(got, q.ans15) {
			t.Errorf("threeSum() = %v, want %v", got, q.ans15)
		}
	}
	fmt.Printf("\n\n\n")
}
