package containsduplicate

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"Test 1", args{nums: []int{1, 2, 3, 1}}, true},
	{"Test 2", args{nums: []int{1, 2, 3, 4}}, false},
	{"Test 3", args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
}

func Test_containsDuplicate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
