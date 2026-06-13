package validanagram

import "testing"

type args struct {
	s string
	t string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"Test 1", args{s: "anagram", t: "nagaram"}, true},
	{"Test 2", args{s: "rat", t: "car"}, false},
	{"Test 3", args{s: "pet", t: "pest"}, false},
	{"Test 4", args{s: "липа", t: "пила"}, true},
}

func Test_isAnagram(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagramV2(t *testing.T) {
	for _, tt := range tests[:len(tests)-1] { // avoid last test case (with Unicode characters)
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramV2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
