package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trevmoody/aoc23/util"
	"testing"
)

func Benchmark_part2(b *testing.B) {
	realInput := util.GetFileAsLines("input")

	for i := 0; i < b.N; i++ {
		part2(*realInput)
	}
}

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
			name: "Test Data",
			args: args{*testInput},
			want: 46,
		},
		{
			name: "Real Input Data",
			args: args{*realInput},
			want: 4917124,
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
