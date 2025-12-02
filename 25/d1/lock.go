package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getDirectionsFile(file_path string) []int {
	var directions []int
	var mult int

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Error reading file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text()[0:1] == "L" {
			mult = -1
		} else {
			mult = 1
		}

		num, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			fmt.Printf("Error parsing direction: %s", scanner.Text())
			os.Exit(1)
		}

		directions = append(directions, num*mult)
	}

	return directions
}

// Number of times zero is landed on
// Answer for my test set is 1139
func part1() {
	directions := getDirectionsFile("test.txt")
	curr_num := 50
	zero_count := 0

	fmt.Printf("Got %d instructions\n", len(directions))

	for _, direction := range directions {
		curr_num += direction % 100
		if curr_num >= 100 {
			curr_num -= 100
		} else if curr_num < 0 {
			curr_num += 100
		}

		if curr_num == 0 {
			zero_count++
		}
	}

	fmt.Printf("There were %d times we hit zero\n", zero_count)
}

// Number of times zero is passed
// Answer for my test set is 6684
func part2() {
	directions := getDirectionsFile("test.txt")
	curr_num := 50
	last_num := 50
	zero_count := 0

	fmt.Printf("Got %d instructions\n", len(directions))

	for _, direction := range directions {
		curr_num += direction

		fmt.Printf("%d", curr_num)

		if curr_num <= 0 && last_num != 0 {
			zero_count++
		}
		if curr_num < 0 {
			zero_count += curr_num / 100 * -1
		} else {
			zero_count += curr_num / 100
		}

		curr_num = curr_num % 100

		if curr_num < 0 {
			curr_num += 100
		}
		fmt.Printf(", %d\n", zero_count)

		last_num = curr_num
	}

	fmt.Printf("There were %d times we clicked zero\n", zero_count)
}

func main() {
	part1()
	part2()
}
