package main

import (
	"bufio"
	"fmt"
	"os"
)

const FILE_PATH = "01-b/input.txt"

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

	occurenceMap := make(map[int]int)

	for i := 0; i < len(locationIds2); i++ {
		locationId := locationIds2[i]

		occurencesSoFar, exists := occurenceMap[locationId]

		if !exists {
			occurencesSoFar = 0
		}

		occurenceMap[locationId] = occurencesSoFar + 1
	}

	similarityScore := 0

	for i := 0; i < len(locationIds1); i++ {
		locationId := locationIds1[i]

		occurences, existed := occurenceMap[locationId]

		if !existed {
			occurences = 0
		}

		similarityScore += locationId * occurences
	}

	fmt.Printf("Similarity Score: %d\n", similarityScore)
}
