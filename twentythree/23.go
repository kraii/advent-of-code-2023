package twentythree

import (
	. "aoc/grids"
)

func solve(file string) int {
	grid := ParseGrid(file)

	start := Point{X: findStartOrEnd(grid[0]), Y: 0}
	end := Point{X: findStartOrEnd(grid[len(grid)-1]), Y: len(grid) - 1}

	return search(grid, start, end, 0, Make2DSlice[bool](len(grid[0]), len(grid)))
}

func search(grid [][]rune, cur Point, end Point, distance int, visited [][]bool) int {
	if cur == end {
		return distance
	}
	visited[cur.Y][cur.X] = true
	maxDist := 0
	for _, point := range findAvailableMoves(grid, cur) {
		if !visited[point.Y][point.X] {
			maxDist = max(maxDist, search(grid, point, end, distance+1, visited))
		}
	}
	visited[cur.Y][cur.X] = false
	return maxDist
}

var directionMap = map[rune]Direction{
	'^': Up,
	'>': Right,
	'v': Down,
	'<': Left,
}

func findAvailableMoves(grid [][]rune, pos Point) []Point {
	currentTile := grid[pos.Y][pos.X]
	if currentTile == '.' {
		var moves []Point
		for _, direction := range Directions {
			newPos := Move1(pos, direction)
			if traversable(grid, newPos) {
				moves = append(moves, newPos)
			}
		}
		return moves
	}

	direction, pres := directionMap[currentTile]
	if !pres {
		message := "wot is" + string(currentTile)
		panic(message)
	}

	newPos := Move1(pos, direction)
	if traversable(grid, newPos) {
		return []Point{newPos}
	}
	return nil
}

func traversable(grid [][]rune, newPos Point) bool {
	return InRange(grid, newPos) && grid[newPos.Y][newPos.X] != '#'
}

func findStartOrEnd(row []rune) int {
	for x, r := range row {
		if r == '.' {
			return x
		}
	}
	panic("Couldn't find start or end")
}
