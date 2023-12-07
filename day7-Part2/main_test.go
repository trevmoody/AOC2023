package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trevmoody/aoc23/util"
	"testing"
)

func Test_part2(t *testing.T) {
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
			want: 5905,
		},
		{
			name: "Real Input",
			args: args{lines: *realInput},
			want: 254115617,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHand(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want hand
	}{
		{
			name: "J9A5A",
			args: args{"J9A5A 123"},
			want: hand{cards: "J9A5A", bid: 123, handType: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHandFromLine(tt.args.line); assert.Equal(t, tt.want, got, "fooked") == false {
				t.Errorf("newHandFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
