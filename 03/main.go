package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const FILE_PATH = "./03/input.txt"

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseMult(mult string) int {
	var num1 int
	var num2 int

	_, err := fmt.Sscanf(mult, "mul(%d,%d)", &num1, &num2)
	if err != nil {
		panic(err)
	}

	return num1 * num2
}

func part1() {
	lines := readFile(FILE_PATH)

	regex := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")

	sum := 0

	for _, line := range lines {
		matches := regex.FindAllString(line, -1)

		for _, match := range matches {
			sum += parseMult(match)
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part2() {
	lines := readFile(FILE_PATH)

	regex := regexp.MustCompile("(mul\\([0-9]+,[0-9]+\\))|(do\\(\\))|(don't\\(\\))")

	active := true
	sum := 0

	for _, line := range lines {
		matches := regex.FindAllString(line, -1)

		for _, match := range matches {
			switch match {
			case "do()":
				active = true
			case "don't()":
				active = false
			default:
				if active {
					sum += parseMult(match)
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	part1()

	part2()
}
