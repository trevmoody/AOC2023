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
			want: 114,
		},
		//{
		//	name: "Real Input",
		//	args: args{lines: *util.GetFileAsLines("input")},
		//	want: 1882395907,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); assert.Equal(t, tt.want, got) == false {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
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
		//{
		//	name: "Real Input",
		//	args: args{lines: *util.GetFileAsLines("input")},
		//	want: 1005,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); assert.Equal(t, tt.want, got) == false {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processLine(t *testing.T) {

	type args struct {
		valuesList []int
	}
	tests := []struct {
		name       string
		args       args
		wantNewVal int
	}{
		{
			name:       "all zero",
			args:       args{[]int{0, 0, 0, 0}},
			wantNewVal: 0,
		},
		{
			name:       "all 3",
			args:       args{[]int{3, 3, 3, 3, 3, 3}},
			wantNewVal: 3,
		},
		{
			name:       "want 18",
			args:       args{[]int{0, 3, 6, 9, 12, 15}},
			wantNewVal: 18,
		}, {
			name:       "want 7",
			args:       args{[]int{2, 3, 4, 5, 6}},
			wantNewVal: 7,
		},
		{
			name:       "want 28",
			args:       args{[]int{1, 3, 6, 10, 15, 21}},
			wantNewVal: 28,
		}, {
			name:       "want 68",
			args:       args{[]int{10, 13, 16, 21, 30, 45}},
			wantNewVal: 68,
		}, {
			name:       "want -2",
			args:       args{[]int{6, 4, 2, 0}},
			wantNewVal: -2,
		},
		{
			name:       "want 5",
			args:       args{[]int{45, 30, 21, 16, 13, 10}},
			wantNewVal: 5,
		},
		{
			name:       "18090394",
			args:       args{[]int{2, -5, 0, 46, 184, 502, 1146, 2356, 4534, 8376, 15141, 27226, 49408, 91460, 173440, 335957, 659454, 1299601, 2551270, 4962878, 9538508}},
			wantNewVal: 18090394,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantNewVal, processLine(tt.args.valuesList), "processLine(%v)", tt.args.valuesList)
		})
	}
}
