package aoc

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

func GetFilePath(name string) string {
	_, currentDirName, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(currentDirName), name)
}

func OpenScanner(file string) *bufio.Scanner {
	f, err := os.Open(GetFilePath(file))
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(bufio.NewReader(f))
}
