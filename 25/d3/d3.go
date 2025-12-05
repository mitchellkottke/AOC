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

func part1(banks []string) {
	total := 0

	for _, bank := range banks {
		first_ind := 0
		second_ind := 0
		max_jolt := '0'

		for battery_ind := 0; battery_ind < len(bank)-1; battery_ind++ {
			battery_jolt := rune(bank[battery_ind])
			if battery_jolt <= max_jolt {
				continue
			}

			max_jolt = battery_jolt
			first_ind = battery_ind

			if battery_jolt == '9' {
				break
			}
		}

		max_jolt = '0'

		for battery_ind := first_ind + 1; battery_ind < len(bank); battery_ind++ {
			battery_jolt := rune(bank[battery_ind])
			if battery_jolt <= max_jolt {
				continue
			}

			max_jolt = battery_jolt
			second_ind = battery_ind

			if max_jolt == '9' {
				break
			}
		}

		bank_jolts_str := bank[first_ind:first_ind+1] + bank[second_ind:second_ind+1]
		fmt.Printf("For bank %s got %s jolts\n", bank, bank_jolts_str)
		bank_jolts, err := strconv.Atoi(bank_jolts_str)
		if err != nil {
			fmt.Printf("Failed to convert string to number: %s", bank_jolts_str)
			os.Exit(1)
		}
		total += bank_jolts
	}

	fmt.Printf("Total jolts: %d\n", total)
}

func main() {
	banks := getBanks("test.txt")

	part1(banks)
}
