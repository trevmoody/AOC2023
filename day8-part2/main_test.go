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
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); assert.Equal(t, tt.want, got) == false {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLeastCommonMultiple(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{2, 5}},
			want: 10,
		},
	}
	// TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getLeastCommonMultiple(tt.args.numbers), "getLeastCommonMultiple(%v)", tt.args.numbers)
		})
	}
}
