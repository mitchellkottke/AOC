package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start string
	end   string
}

func parseRanges(file_path string) []Range {
	var ranges []Range
	var r string
	var err error
	var file *os.File

	file, err = os.Open(file_path)
	if err != nil {
		fmt.Printf("Error reading file")
		os.Exit(1)
	}
	defer file.Close()

	//Input is expected to be on one line
	//Reading via reader saves allocating a large buffer
	//	and reading the whole line into memory
	reader := bufio.NewReader(file)
	for r, err = reader.ReadString(','); err == nil; r, err = reader.ReadString(',') {
		parts := strings.Split(r[:len(r)-1], "-")
		ranges = append(ranges, Range{parts[0], parts[1]})
	}

	if err == io.EOF {
		//ReadString returns error for EOF when section doesn't end with delim
		//Append last range after EOF
		parts := strings.Split(r, "-")
		ranges = append(ranges, Range{parts[0], parts[1]})
	} else {
		fmt.Printf("Error while parsing file: %s\n", err.Error())
		os.Exit(1)
	}

	return ranges
}

// Find all numbers in ranges that are the same when split in half and sum them
// Ex: 123123 is true, 1231234 is false
// My answer is: 24157613387
func part1(ranges []Range) {
	total := 0

	for _, r := range ranges {
		end, err := strconv.Atoi(r.end)
		if err != nil {
			fmt.Printf("Error while parsing int: %s\n", err.Error())
			os.Exit(1)
		}

		for current, err := strconv.Atoi(r.start); err == nil && current <= end; current++ {
			current_str := strconv.Itoa(current)
			if len(current_str)%2 != 0 {
				continue
			}

			middle := len(current_str) / 2
			if current_str[:middle] != current_str[middle:] {
				continue
			}

			total += current
		}
	}

	fmt.Printf("Part 1 total is: %d\n", total)
}

// Find all numbers in ranges that are made up of repeating sets of digets
// Ex: 262626 is valid, 262623 is invalid, 2222222 is valid
// My answer is: 33832678380
func part2(ranges []Range) {
	var total int64 = 0

	//For every number range
	for _, r := range ranges {
		end, err := strconv.Atoi(r.end)
		if err != nil {
			fmt.Printf("Error while parsing int: %s\n", err.Error())
			os.Exit(1)
		}

		//For every number in the range
		for current, err := strconv.Atoi(r.start); err == nil && current <= end; current++ {
			current_str := strconv.Itoa(current)

			//For every possible split count of the number
			//In this case num_split = 2 means split in half
			for num_split := len(current_str); num_split > 1; num_split-- {
				//Splits are only valid if it's an even split
				if len(current_str)%num_split != 0 {
					continue
				}

				split := len(current_str) / num_split
				found_valid_split := true

				//Verify every split matches the next
				for left_start := 0; left_start < len(current_str)-split; left_start += split {
					right_start := left_start + split

					if current_str[left_start:right_start] != current_str[right_start:right_start+split] {
						found_valid_split = false
						break
					}
				}

				if found_valid_split {
					total += int64(current)
					break
				}
			}
		}
	}

	fmt.Printf("Part 2 total is: %d\n", total)
}

func main() {
	ranges := parseRanges("test.txt")
	part1(ranges)
	part2(ranges)
}
