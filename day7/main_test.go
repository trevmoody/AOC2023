package main

import (
	"github.com/trevmoody/aoc2023/util"
	"testing"
)

func Test_less(t *testing.T) {

	type args struct {
		h1 hand
		h2 hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "KK677 and KTJJT",
			args: args{newHand("KK677 123"), newHand("KTJJT 123")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := less(tt.args.h1, tt.args.h2); got != tt.want {
				t.Errorf("less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	testInput := util.GetFileAsLines("testinput")
	realInput := util.GetFileAsLines("input")

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
			args: args{lines: *testInput},
			want: 6440,
		},
		{
			name: "Real Input",
			args: args{lines: *realInput},
			want: 6440,
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
