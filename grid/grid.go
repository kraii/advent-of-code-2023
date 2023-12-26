package grid

import "fmt"

type Point struct {
	X, Y int
}

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

func PrintDir(dir Direction) string {
	switch dir {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic(badDirection(dir))
	}
}

func Move1(p Point, d Direction) Point {
	return Move(p, d, 1)
}

func Move(p Point, d Direction, distance int) Point {
	switch d {
	case Up:
		return Point{p.X, p.Y - distance}
	case Down:
		return Point{p.X, p.Y + distance}
	case Left:
		return Point{p.X - distance, p.Y}
	case Right:
		return Point{p.X + distance, p.Y}
	}
	panic(badDirection(d))
}

func badDirection(d Direction) string {
	return fmt.Sprintf("What the heck direction is %d", d)
}

func TurnLeft(current Direction) Direction {
	switch current {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	default:
		panic(badDirection(current))
	}
}

func TurnRight(current Direction) Direction {
	switch current {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	default:
		panic(badDirection(current))
	}
}

func InRange[T any](grid [][]T, p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(grid) && p.X < len(grid[0])
}
