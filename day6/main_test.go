package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trevmoody/aoc2023/util"
	"testing"
)

func Test_part1(t *testing.T) {
	testInput := util.GetFileAsLines("testinput")
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{*testInput},
			want: 288,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.strings); assert.Equal(t, got, tt.want) == false {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "realInput",
			want: 21039729,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, part2(), "part2()")
		})
	}
}
