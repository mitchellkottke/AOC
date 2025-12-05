package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getBanks(file_path string) []string {
	var banks []string

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		banks = append(banks, scanner.Text())
	}

	return banks
}

func getMaxJolts(banks []string, num_digits int) int64 {
	total := int64(0)

	for _, bank := range banks {
		last_ind := -1
		bank_jolts_str := ""

		for curr_diget := 1; curr_diget <= num_digits; curr_diget++ {
			max_jolt := '0'

			for battery_ind := last_ind + 1; battery_ind < len(bank)-num_digits+curr_diget; battery_ind++ {
				battery_jolt := rune(bank[battery_ind])
				if battery_jolt <= max_jolt {
					continue
				}

				max_jolt = battery_jolt
				last_ind = battery_ind

				if max_jolt == '9' {
					break
				}
			}

			bank_jolts_str += bank[last_ind : last_ind+1]
		}

		bank_jolts, err := strconv.ParseInt(bank_jolts_str, 10, 64)
		if err != nil {
			fmt.Printf("Failed to convert string to number: %s", bank_jolts_str)
			os.Exit(1)
		}
		total += bank_jolts
	}

	return total
}

// Answer from my test set: 16973
func part1(banks []string) {
	fmt.Printf("Part 1 jolts: %d\n", getMaxJolts(banks, 2))
}

// Answer from my test set: 168027167146027
func part2(banks []string) {
	fmt.Printf("Part 2 jolts: %d\n", getMaxJolts(banks, 12))
}

func main() {
	banks := getBanks("test.txt")

	part1(banks)
	part2(banks)
}
