package utils

import (
	"advent_of_code/2023/models"
	"advent_of_code/2023/models_imp"
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetLines(filename string) ([]models.GameLine, error) {
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
