package main

import (
	"advent_of_code/2023/models"
	"advent_of_code/2023/models_imp"
	"advent_of_code/2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func printGameLine(gL models.GameLine) {
	fmt.Printf("GameLine number %d | ", gL.GameNumber)
	fmt.Printf("GameLine line: %s\n", gL.Line)
	fmt.Println("\tGameLine Sets:")
	for i := range gL.Subsets {
		fmt.Printf("\t\tSet %d:\n", i)
		for j := range gL.Subsets[i].Colours {
			fmt.Printf("\t\t\t%d %s\n", gL.Subsets[i].Colours[j].Count, gL.Subsets[i].Colours[j].Colour)
		}
	}
}

func removeSpacesFromArray(arr []string) []string {
	for i := range arr {
		arr[i] = strings.TrimSpace(arr[i])
	}
	return arr
}

func setGameNumber(gL *models.GameLine) {
	res, err := strconv.Atoi(regexp.MustCompile("[0-9]+").FindAllString(gL.Line, 1)[0])
	utils.Check(err)
	gL.Line = strings.Split(gL.Line, ":")[1]
	gL.GameNumber = res
}

func getColour(set string) models.Colour {
	split_vals := removeSpacesFromArray(strings.Split(set, " "))

	count, err := strconv.Atoi(split_vals[0])
	utils.Check(err)
	colour := split_vals[1]

	return models_imp.NewColour(count, colour)
}

func setSubsets(gL *models.GameLine) {

	// Split vals represents our sets in our subset
	split_subset := removeSpacesFromArray(strings.Split(gL.Line, ";"))

	// For each 'set' in the split_subset
	for _, set := range split_subset {
		// Get the individual sets in subset
		split_sets := removeSpacesFromArray(strings.Split(set, ","))

		// Declare and initialise a 'Set'
		new_set := models.Set{}

		// For each pair in each set
		for _, split_set := range split_sets {
			// Add the `Colour` to the new set's Colours array
			new_set.Colours = append(new_set.Colours, getColour(split_set))
		}

		// Append the new set to the GameLine Set array
		gL.Subsets = append(gL.Subsets, new_set)
	}
}

func check_colour(given string, expected string) {
	if given != expected {
		panic(fmt.Sprintf("Expected colour was %s, actual was %s", expected, given))
	}
	fmt.Printf("Given colour of %s matches expected colour of %s\n", given, expected)
}

func parseGameLine(gL models.GameLine, colour_limits [3]int) bool {
	setGameNumber(&gL)
	setSubsets(&gL)

	setMaximums(&gL)

	for i := range gL.Maximums {
		if gL.Maximums[i] > colour_limits[i] {
			return false
		}
	}

	return true
}

func setMaximums(gL *models.GameLine) {
	for i := range gL.Subsets {
		for j := range gL.Subsets[i].Colours {
			switch colour := gL.Subsets[i].Colours[j].Colour; colour {
			case "red":
				if gL.Subsets[i].Colours[j].Count > gL.Maximums[0] {
					gL.Maximums[0] = gL.Subsets[i].Colours[j].Count
				}

			case "green":
				if gL.Subsets[i].Colours[j].Count > gL.Maximums[1] {
					gL.Maximums[1] = gL.Subsets[i].Colours[j].Count
				}

			case "blue":
				if gL.Subsets[i].Colours[j].Count > gL.Maximums[2] {
					gL.Maximums[2] = gL.Subsets[i].Colours[j].Count
				}

			}
		}
	}
}

func main() {
	var data []models.GameLine
	data, err := utils.GetLines("data.txt")
	utils.Check(err)

	colour_limits := []int{12, 13, 14}

	results := []bool{}
	for i := range data {
		results = append(results, parseGameLine(data[i], [3]int(colour_limits)))
	}

	fmt.Printf("The following lines are possible:")
	total := 0
	for i := range results {
		if results[i] {
			fmt.Printf(" %d", i+1)
			total += (i + 1)
		}
	}
	fmt.Printf("\nTotal is %d\n", total)

}
