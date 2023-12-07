package main

import (
	"os"
	"testing"
)

func Test_processFile(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processFile(tt.args.file)
		})
	}
}
