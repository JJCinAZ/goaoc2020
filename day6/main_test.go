package main

import "testing"

func Test_countUnanomAnswers(t *testing.T) {
	type args struct {
		grp Group
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{grp: Group{People: []Answers{"abc"}}}, 3},
		{"test2", args{grp: Group{People: []Answers{"a", "b", "c"}}}, 0},
		{"test3", args{grp: Group{People: []Answers{"ab", "ac"}}}, 1},
		{"test4", args{grp: Group{People: []Answers{"a", "a"}}}, 1},
		{"test5", args{grp: Group{People: []Answers{"b"}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnanomAnswers(tt.args.grp); got != tt.want {
				t.Errorf("countUnanomAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
