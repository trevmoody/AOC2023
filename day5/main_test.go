package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trevmoody/aoc2023/util"
	"testing"
)

func Test_part2(t *testing.T) {

	testInput := util.GetFileAsLines("./testinput")

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Data",
			args: args{*testInput},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); assert.Equal(t, tt.want, got, "they should be equal") == false {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
