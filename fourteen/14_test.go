package fourteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTilt1Column(t *testing.T) {
	p := platform{{'O'}, {'O'}, {'.'}, {'O'}, {'.'}, {'O'}, {'.'}, {'.'}, {'#'}, {'#'}}
	tilt(p)
	expected := platform{{'O'}, {'O'}, {'0'}, {'O'}, {'.'}, {'.'}, {'.'}, {'.'}, {'#'}, {'#'}}
	assert.Equal(t, expected, p)
}

func TestTilt(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	printPlatform(p)
	tilt(p)
	println()
	printPlatform(p)

	expected := parse("fourteen/14-ex-1-expected.txt")
	assert.Equal(t, expected, p)
}

func printPlatform(p platform) {
	for _, row := range p {
		for _, r := range row {
			print(string(r))
		}
		println()
	}
}
