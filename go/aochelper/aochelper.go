package aochelper

import (
	"fmt"
	"os"
)

const inputDir = "../input"

func ReadInputFile(day int) string {
	fileName := fmt.Sprintf("%s/day%02d.txt", inputDir, day)
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("error reading input file %s: %s", fileName, err))
	}
	return string(data)
}
