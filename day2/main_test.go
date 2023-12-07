package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	line := "Game day1: day3 blue, day4 red; day1 red, day2 green, 6 blue; day2 green"

	result := parse(line)

	try1 := map[string]int{"blue": 3, "red": 4}
	try2 := map[string]int{"blue": 6, "red": 1, "green": 2}

	assert.Equal(t, try1, try2, "they should be equal")

	if result.id != 1 {
		t.Errorf("id not matching, %d => %d", result.id, 1)
	}

	if reflect.DeepEqual(result, game{1, []map[string]int{try1, try2}}) == false {
		t.Errorf("fooked")
	}

}
