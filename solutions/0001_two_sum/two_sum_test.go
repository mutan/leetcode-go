package twoSum

import (
	"reflect"
	"testing"
)

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"Test 1", args{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},
	{"Test 2", args{nums: []int{3, 2, 4}, target: 6}, []int{1, 2}},
	{"Test 3", args{nums: []int{3, 3}, target: 6}, []int{0, 1}},
}

func Test_twoSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumBruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumBruteForce(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
