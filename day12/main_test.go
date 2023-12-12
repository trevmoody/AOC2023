package main

import "testing"

func Test_count(t *testing.T) {
	type args struct {
		conditionRecord string
		check           []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1.", args: args{".", []int{1}}, want: 0},
		{name: "1#", args: args{"#", []int{1}}, want: 1},
		{name: "1?", args: args{"?", []int{1}}, want: 1},
		{name: "2.", args: args{"..", []int{1}}, want: 0},
		{name: "2#", args: args{"##", []int{1}}, want: 0},
		{name: "2?", args: args{"??", []int{1}}, want: 2},
		// group size 2
		{name: "1.", args: args{".", []int{2}}, want: 0},
		{name: "1#", args: args{"#", []int{2}}, want: 0},
		{name: "1?", args: args{"?", []int{2}}, want: 0},
		{name: "2.", args: args{"..", []int{2}}, want: 0},
		{name: "2#", args: args{"##", []int{2}}, want: 1},
		{name: "2?", args: args{"??", []int{2}}, want: 1},
		{name: "testInput1", args: args{"???.###", []int{1, 1, 3}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.conditionRecord, tt.args.check, make(map[state]int)); got != tt.want {
				t.Errorf("getPossibleConditionRecordsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
