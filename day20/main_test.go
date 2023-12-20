package main

import (
	"github.com/trevmoody/aoc2023/util"
	"testing"
)

func Test_part1(t *testing.T) {
	testinput := *util.GetFileAsLines("testinput")
	testinput2 := *util.GetFileAsLines("testinput2")

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{testinput},
			want: 32000000,
		},
		{
			name: "testinput2",
			args: args{testinput2},
			want: 11687500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
