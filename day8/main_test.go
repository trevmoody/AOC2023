package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trevmoody/aoc2023/util"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Input",
			args: args{lines: *util.GetFileAsLines("testinput")},
			want: 2,
		},
		{
			name: "Test Input2",
			args: args{lines: *util.GetFileAsLines("testinput2")},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); assert.Equal(t, tt.want, got) == false {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
