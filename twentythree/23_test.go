package twentythree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = "twentythree/ex.txt"
const input = "twentythree/input.txt"

func TestExample1(t *testing.T) {
	assert.Equal(t, 94, solve(example))
}

func TestFindStart(t *testing.T) {
	assert.Equal(t, 1, findStartOrEnd([]rune("#.#####################")))
}

func TestSolvePart1(t *testing.T) {
	println(solve(input))
}
