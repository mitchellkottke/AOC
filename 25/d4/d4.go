package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRollMap(file_path string) [][]bool {
	rolls := [][]bool{}

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Error opening file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row_str := scanner.Text()
		curr_row := []bool{}

		for _, roll := range row_str {
			if roll == '@' {
				curr_row = append(curr_row, true)
			} else {
				curr_row = append(curr_row, false)
			}
		}

		rolls = append(rolls, curr_row)
	}

	return rolls
}

func getRollNeighborCount(rolls [][]bool) [][]int {
	counts := make([][]int, len(rolls))
	for ind := range rolls {
		counts[ind] = make([]int, len(rolls[0]))
	}

	for row_ind, row := range rolls {
		for col_ind, roll := range row {
			if !roll {
				continue
			}

			for upd_row := row_ind - 1; upd_row <= row_ind+1; upd_row++ {
				if upd_row < 0 || upd_row >= len(rolls) {
					continue
				}

				for upd_col := col_ind - 1; upd_col <= col_ind+1; upd_col++ {
					if upd_col < 0 || upd_col >= len(rolls[0]) || (upd_col == col_ind && upd_row == row_ind) {
						continue
					}

					counts[upd_row][upd_col]++
				}
			}
		}
	}

	return counts
}

// Answer for my test set: 1320
func part1(rolls [][]bool) {
	counts := getRollNeighborCount(rolls)
	total := 0

	for row_ind, row := range counts {
		for col_ind, count := range row {
			if count < 4 && rolls[row_ind][col_ind] {
				total++
			}
		}
	}

	fmt.Printf("Part 1 number of rolls: %d\n", total)
}

// Answer for my test set: 8354
func part2(rolls [][]bool) {
	total := 0

	for {
		counts := getRollNeighborCount(rolls)
		round_count := 0

		for row_ind, row := range counts {
			for col_ind, count := range row {
				if count < 4 && rolls[row_ind][col_ind] {
					round_count++
					rolls[row_ind][col_ind] = false
				}
			}
		}

		total += round_count

		if round_count == 0 {
			break
		}
	}

	fmt.Printf("Part 2 number of rolls: %d", total)
}

func main() {
	rolls := getRollMap("test.txt")

	part1(rolls)
	part2(rolls)
}
