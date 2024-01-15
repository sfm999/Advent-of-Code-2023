package utils

import (
	"advent_of_code/2023/models"
	"advent_of_code/2023/models_imp"
	"bufio"
	"fmt"
	"os"
)

// Check whether error was nil, if so, panic!
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetGameLines(filename string) ([]models.GameLine, error) {
	readFile, err := os.Open(filename)
	Check(err)

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	var fileLines []models.GameLine

	for fileScanner.Scan() {
		fileLines = append(fileLines, *models_imp.NewGameLine(fileScanner.Text()))
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	return fileLines, nil
}

func GetStringLines(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	Check(err)

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	return fileLines, nil
}

func Get2dRuneArray(filename string) ([][]rune, error) {
	readFile, err := os.Open(filename)
	Check(err)

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var grid [][]rune
	for fileScanner.Scan() {
		line := fileScanner.Text()

		row := []rune(line)
		grid = append(grid, row)
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func PrintStringArray(arr []string) {
	for _, val := range arr {
		fmt.Println(val)
	}
}
