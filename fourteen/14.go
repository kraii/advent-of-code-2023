package fourteen

import "aoc"

type platform [][]rune

func tilt(p platform) {
	for y, row := range p {
		for x, row := range row {
			if row == 'O' {
				rollRock(p, x, y)
			}
		}
	}
}

func rollRock(p platform, x int, startY int) {
	y := startY - 1
	for y >= 0 {
		tileAbove := p[y][x]
		if tileAbove != '.' {
			if y != startY {
				p[y+1][x] = 'O'
				p[startY][x] = '.'
			}
			return
		}
		y--
	}
	// we got to the bottom
	p[0][x] = 'O'
	p[startY][x] = '.'
}

func parse(file string) platform {
	scanner := aoc.OpenScanner(file)
	var result platform

	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result
}
