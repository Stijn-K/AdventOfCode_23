package common

import (
	"bufio"
	"os"
)

func ReadInputFile(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

func ReadInputFileAsRunes(filePath string) [][]rune {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var data [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}
