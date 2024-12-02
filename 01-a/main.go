package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const FILE_PATH = "01-a/input.txt"

func readFile(path string) ([]int, []int) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	locationIds1 := make([]int, 0)
	locationIds2 := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var id1 int
		var id2 int

		_, err := fmt.Sscanf(scanner.Text(), "%d   %d", &id1, &id2)

		if err != nil {
			panic(err)
		}

		locationIds1 = append(locationIds1, id1)
		locationIds2 = append(locationIds2, id2)
	}

	return locationIds1, locationIds2
}

func main() {
	locationIds1, locationIds2 := readFile(FILE_PATH)

	slices.Sort(locationIds1)
	slices.Sort(locationIds2)

	sum := 0

	for i := 0; i < len(locationIds1); i++ {
		difference := locationIds1[i] - locationIds2[i]

		if difference < 0 {
			difference = -difference
		}

		sum += difference
	}

	fmt.Printf("Sum: %d\n", sum)
}
