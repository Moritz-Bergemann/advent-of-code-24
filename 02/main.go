package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const FILE_PATH = "02/input-sample.txt"

const FILE_PATH = "02/input.txt"

func readFile(path string) [][]int {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		readingsStr := strings.Split(line, " ")

		readings := make([]int, len(readingsStr))
		for idx, readingStr := range readingsStr {
			reading, err := strconv.Atoi(readingStr)

			if err != nil {
				panic(err)
			}

			readings[idx] = reading
		}

		reports = append(reports, readings)
	}

	return reports
}

func validateReport(report []int) bool {
	isAsc := true
	for idx := range report {
		if idx == len(report)-1 {
			break
		}

		if report[idx] > report[idx+1] {
			isAsc = false
			break
		}
	}

	isDesc := true
	for idx := range report {
		if idx == len(report)-1 {
			break
		}

		if report[idx] < report[idx+1] {
			isDesc = false
			break
		}
	}

	safeGaps := true
	for idx := range report {
		if idx == len(report)-1 {
			break
		}

		gap := report[idx] - report[idx+1]

		if gap < 0 {
			gap = -gap
		}

		if gap > 3 || gap < 1 {
			safeGaps = false
			break
		}
	}

	if (isAsc || isDesc) && safeGaps {
		return true
	}
	return false
}

func taskA() {
	reports := readFile(FILE_PATH)

	numSafeReports := 0

	for _, report := range reports {
		safe := validateReport(report)

		if safe {
			numSafeReports++
		}
	}

	fmt.Printf("Part A: %d\n", numSafeReports)
}

func taskB() {
	reports := readFile(FILE_PATH)

	numSafeReports := 0

	for _, report := range reports {
		safe := false

		for levelIdx := range report {
			reportWithoutLevel := make([]int, len(report))
			copy(reportWithoutLevel, report)

			reportWithoutLevel = append(reportWithoutLevel[:levelIdx], reportWithoutLevel[levelIdx+1:]...)

			safeWithoutLevel := validateReport(reportWithoutLevel)

			if safeWithoutLevel {
				safe = true
				break
			}
		}

		if safe {
			numSafeReports++
		}
	}

	fmt.Printf("Part B: %d\n", numSafeReports)
}

func main() {
	taskA()
	taskB()
}
