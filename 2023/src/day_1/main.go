package main

import (
	"advent_of_code/2023/models"
	"advent_of_code/2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

func part_one(data []models.GameLine) (int, error) {
	total := 0
	for _, gL := range data {

		// Compile regex string
		numbers := regexp.MustCompile(`\d`).FindAllString(gL.Line, -1)
		if numbers == nil {
			return -1, fmt.Errorf("no numbers found in the line: %s", gL.Line)
		}

		// Check for two numbers
		if len(numbers) > 1 {
			num, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
			if err != nil {
				return -1, err
			}
			total += num
		}

		// Check for one number
		if len(numbers) == 1 {
			num, err := strconv.Atoi(numbers[0] + numbers[0])
			if err != nil {
				return -1, err
			}
			total += num
		}
	}
	return total, nil
}

func main() {
	var data []models.GameLine
	data, err := utils.GetLines("data.txt")
	utils.Check(err)

	for _, gL := range data {
		fmt.Println(gL)
	}

	res, err := part_one(data)
	utils.Check(err)
	fmt.Println(res)
}
