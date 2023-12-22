package main

import (
	"fmt"
	"github.com/trevmoody/aoc2023/util"
	"sort"
)

type Brick struct {
	x1 int
	y1 int
	z1 int
	x2 int
	y2 int
	z2 int
}

func parse(lines []string) []Brick {
	var bricks []Brick
	for _, line := range lines {
		b := Brick{}
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.x1, &b.y1, &b.z1, &b.x2, &b.y2, &b.z2)
		bricks = append(bricks, b)
	}

	return bricks
}

func (b Brick) intersect(above Brick) bool {
	return b.x1 <= above.x2 && b.x2 >= above.x1 && b.y1 <= above.y2 && b.y2 >= above.y1
}

func (b Brick) drop() Brick {
	return Brick{x1: b.x1, x2: b.x2, y1: b.y1, y2: b.y2, z1: b.z1 - 1, z2: b.z2 - 1}
}

func fallDown(bricks []Brick) {
	for brickId := 0; brickId < len(bricks); brickId++ {
		canDrop := true

		for bricks[brickId].z1 > 1 && canDrop {
			for belowBrickID := brickId - 1; belowBrickID >= 0; belowBrickID-- {
				if bricks[belowBrickID].z2 == bricks[brickId].z1-1 && bricks[belowBrickID].intersect(bricks[brickId]) {
					canDrop = false
					break
				}
			}
			if canDrop {
				droppedBrick := bricks[brickId].drop()
				bricks[brickId] = droppedBrick
			}
		}
	}
}

func part1(bricks []Brick) {

	safeCount := 0
	tmp := make([]Brick, len(bricks))
	for i := range bricks {
		copy(tmp, bricks)
		tmp[i] = Brick{
			x1: 0,
			y1: 0,
			z1: 0,
			x2: 0,
			y2: 0,
			z2: 0,
		}

		if disintegrateCount(tmp) == 0 {
			safeCount++
		}
	}
	fmt.Printf("Part 1 Got %d safe bricks to disintegrate\n", safeCount)
}

func part2(bricks []Brick) {
	tmp := make([]Brick, len(bricks))
	fallCount := 0
	for i := range bricks {
		copy(tmp, bricks)
		tmp[i] = Brick{
			x1: 0,
			y1: 0,
			z1: 0,
			x2: 0,
			y2: 0,
			z2: 0,
		}
		fallCount += disintegrateCount(tmp)
	}

	fmt.Printf("Part 2 %d", fallCount)

}

func disintegrateCount(bricks []Brick) int {
	fallenBrickId := map[int]bool{}
	for brickId := 0; brickId < len(bricks); brickId++ {
		canDrop := true
		for bricks[brickId].z1 > 1 && canDrop {
			for belowBrickID := brickId - 1; belowBrickID >= 0; belowBrickID-- {
				if bricks[belowBrickID].z2 == bricks[brickId].z1-1 && bricks[belowBrickID].intersect(bricks[brickId]) {
					canDrop = false
					break
				}
			}
			if canDrop {
				droppedBrick := bricks[brickId].drop()
				bricks[brickId] = droppedBrick
				fallenBrickId[brickId] = true
			}
		}
	}

	return len(fallenBrickId)
}

func main() {
	lines := *util.GetFileAsLines("input")

	bricks := parse(lines)

	// get the lowest z's to start of the list.
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 < bricks[j].z1
	})

	fallDown(bricks)

	part1(bricks)
	part2(bricks)
}
