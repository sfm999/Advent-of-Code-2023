package main

import (
	"advent_of_code/2023/utils"
	"fmt"
	"unicode"
)

type Pair struct {
	width  int
	height int
}

// func check_height(field models.Grid, height_index int) bool {
// 	return height_index >= 0 && height_index < field.Height
// }

// func check_left_valid(width_index int) bool {
// 	return width_index-1 > 0
// }

// func check_right_valid(field models.Grid, width_index int) bool {
// 	return width_index+1 < field.Width
// }

// func check_top_side(field models.Grid, width_index int, height_index int) bool {

// 	// Left and right only care about width

// 	// Nothing to the left
// 	if !check_left_valid(width_index) {
// 		// do something
// 	}

// 	// Nothing to the right
// 	if !check_right_valid(field, width_index) {
// 		// do something
// 	}

// 	return false
// }

// func get_sides_to_check(grid models.Grid, width_index int, height_index int) []Pair {
// 	left := Pair{width: width_index - 1, height: height_index}
// 	right := Pair{width: width_index + 1, height: height_index}

// 	top := Pair{width: width_index, height: height_index - 1}
// 	bot := Pair{width: width_index, height: height_index + 1}

// 	top_left := Pair{width: width_index - 1, height: height_index - 1}
// 	top_right := Pair{width: width_index + 1, height: height_index - 1}

// 	bot_left := Pair{width: width_index - 1, height: height_index + 1}
// 	bot_right := Pair{width: width_index + 1, height: height_index + 1}

// 	return []Pair{left, right, top, bot, top_left, top_right, bot_left, bot_right}
// }

// func check_adjacent(field models.Grid, width_index int, height_index int) {

// 	sides := get_sides_to_check(width_index, height_index)

// }

func duplicate_is_consecutive_number(row int, col int, data []string, bool_chart [][]bool) {
	/*
		What if I simply start at the beginning of each line and find the sequences
	*/
}

// Once we have a sequence, should we then check if the sequence is even valid?
// If we find a number in a line, find the sequence of it then check if any of the
// numbers touch a star
// E.g. Row 0, cols 0-2 are numbers. Pass the row number, and a list of cols, the original data,
// and the bool_chart into a function that checks if ANY of the numbers touch a star.
// If any of them touch a star, set each col in the bool_chart row to True
// func is_consecutive_number(row int, col int, data []string, bool_chart [][]bool) {
// 	// check behind
// 	if col > 0 {
// 		curr := col
// 		for curr >= 0 {
// 			if data[row][curr] < 48 || data[row][curr] > 57 {
// 				break
// 			} else {
// 				bool_chart[row][curr] = true
// 			}
// 			curr -= 1
// 		}
// 	}

// 	// check infront
// 	if col < len(data[row])-1 {
// 		curr := col
// 		for curr < len(data[row]) {
// 			if data[row][curr] < 48 || data[row][curr] > 57 {
// 				break
// 			} else {
// 				bool_chart[row][curr] = true
// 			}
// 			curr += 1
// 		}
// 	}
// }

// row 0, col 0 is expected to produce false, true, false, true, false, false, false, true
// row len(data)-1, col len(data[row])-1 is expected to produce true, false, true, false, true, false, false, false
func is_adjacent(row int, col int, data []string, b_chart [][]bool) bool {
	CHECK_ABOVE := row > 0
	// CHECK_BELOW := row < len(data)-1
	// CHECK_LEFT := col > 0
	// CHECK_RIGHT := col < len(data[row])-1
	// CHECK_DIAGONAL_TOP_LEFT := CHECK_ABOVE && CHECK_LEFT
	// CHECK_DIAGONAL_TOP_RIGHT := CHECK_ABOVE && CHECK_RIGHT
	// CHECK_DIAGONAL_BOT_LEFT := CHECK_BELOW && CHECK_LEFT
	// CHECK_DIAGONAL_BOT_RIGHT := CHECK_BELOW && CHECK_RIGHT

	SYMBOLS := []byte{'@', '#', '+', '-', '&', '*', '=', '%', '/', '$'}

	if CHECK_ABOVE {
		for i := range SYMBOLS {
			if data[row-1][col] == SYMBOLS[i] {
				return true
			}
		}
	}
	return false
}

// Given a row and a set of columns, change the same position in the 2d bool array
// to true to indicate it's touching a star. If a single number in the sequence is touching a star,
// set each position in the sequence to true
func check_adjacent_to_star(row int, cols []int, data []string, b_chart [][]bool) {
	curr := 0
	// For each column to check in row
	for curr < len(cols) {
		// Find out if the current col in row is adjacent to a star
		b_chart[row][cols[curr]] = is_adjacent(row, cols[curr], data, b_chart)

		// If it's adjacent, change each col pos in row to true as any
		// sequence of numbers touching ANY star will make the entire sequence
		// touching a star
		if b_chart[row][cols[curr]] {

			// Changing each col pos in row to true
			for i := range cols {
				b_chart[row][cols[i]] = true
			}
			// Exit function
			return
		}

		// Continues implicitly if col pos wasn't adjacent to a star
		curr += 1
	}
}

/*
This function is intended to scan a given row in our data and
determine the positions of any sequences of numbers.

Do a nested for loop: `i`, `j`. The outer loop (i) is iterating the length of the string.
The inner loop (j) will start only IF we find a number at the index of outer loops
current iteration. Before starting the inner loop, create an array of integers to represent
the col indexes of the sequence of numbers. The first value in cols will be `i`.

In the inner loop, j will start at i+1 and iterate to the end of the string
OR until it finds a non-number (A.K.A breaking the sequence).
For each data[row][j] that is a number, add that index to the cols array.
If it's a non-number, set i to j and break. We set i to j so that we can continue from
where we're up to in the string.
*/
func get_sequence(row int, data []string, b_chart [][]bool) {
	// Iterate whole line
	for i := 0; i < len(data[row])-1; i++ {
		// If current is a digit, start sequence finding
		if unicode.IsDigit(rune(data[row][i])) {

			// Add current position to the cols to check
			cols := []int{i}

			// Starting from next character in line (i+1)
			// Iterate till potentially end of line
			for j := i + 1; j < len(data[row])-1; j++ {

				// If next character is not a digit
				if !unicode.IsDigit(rune(data[row][j])) {

					// Set i to where j got up to, then break the inner loop
					i = j
					break
				}

				// Otherwise add current pos to the cols to check
				cols = append(cols, j)
			}
			fmt.Println(cols)
			check_adjacent_to_star(row, cols, data, b_chart)
		}
		fmt.Println("I equals", i)

	}
}

func main() {

	// Refer to individual characters with string() wrapper
	data, err := utils.GetStringLines("test_data.txt")

	utils.Check(err)

	b_chart := make([][]bool, len(data))

	for i := range b_chart {
		b_chart[i] = make([]bool, len(data[0]))
	}

	// for i := range data {
	// for j := range data[i] {
	// b_chart = append(b_chart, )
	// }
	// }

	// is_consecutive_number(0, 2, data, b_chart)

	// is_adjacent(0, 0, data)
	// is_adjacent(len(data)-1, len(data[len(data)-1])-1, data, b_chart)
	// is_adjacent(2, 2, data, b_chart) // Test check above

	// test_cols := []int{2, 3}
	// check_adjacent_to_star(2, test_cols, data, b_chart)
	// fmt.Println(b_chart[2])

	get_sequence(2, data, b_chart)
	fmt.Println(b_chart[2])

	// fmt.Println(b_chart[0])

	// fmt.Println(b_chart)

}
