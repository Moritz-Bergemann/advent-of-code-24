package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// const FILE_PATH = "./04/input-sample.txt"

// const FILE_PATH = "./04/input-debug.txt"

const FILE_PATH = "./04/input.txt"

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rows := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows
}

func getReverseString(in string) string {
	out := make([]rune, len(in))

	for idx, char := range in {
		reverseIdx := len(in) - 1 - idx

		out[reverseIdx] = char
	}

	return string(out)
}

func scanBackSlash(rows []string, startRow int, startColumn int) string {
	rowIdx := startRow
	columnIdx := startColumn

	diagonalChars := make([]rune, 0)

	for rowIdx < len(rows) && columnIdx < len(rows[0]) && rowIdx >= 0 && columnIdx >= 0 {
		diagonalChars = append(diagonalChars, rune(rows[rowIdx][columnIdx]))

		rowIdx++
		columnIdx++
	}

	return string(diagonalChars)
}

func scanForwardSlash(rows []string, startRow int, startColumn int) string {
	rowIdx := startRow
	columnIdx := startColumn

	diagonalChars := make([]rune, 0)

	for rowIdx < len(rows) && columnIdx < len(rows[0]) && rowIdx >= 0 && columnIdx >= 0 {
		diagonalChars = append(diagonalChars, rune(rows[rowIdx][columnIdx]))

		rowIdx++
		columnIdx--
	}

	return string(diagonalChars)
}

func rowsToScannableLines(rows []string) []string {
	scannableLines := make([]string, 0)

	// Rows (& reverse)
	for _, row := range rows {
		rowReverse := getReverseString(row)

		scannableLines = append(scannableLines, row, rowReverse)
	}

	// Columns (& reverse)
	for columnIdx := range rows[0] {
		columnChars := make([]rune, len(rows))

		for rowIdx := 0; rowIdx < len(rows); rowIdx++ {
			columnChars[rowIdx] = rune(rows[rowIdx][columnIdx])
		}

		column := string(columnChars)

		columnReverse := getReverseString(column)

		scannableLines = append(scannableLines, column, columnReverse)
	}

	// '\' Diagonals (& reverse)
	rowIndexesBackSlash := make([]int, 0)
	columnIndexesBackSlash := make([]int, 0)

	for rowIndex := range rows {
		// Skip 0,0 because the next one covers it
		if rowIndex == 0 {
			continue
		}
		rowIndexesBackSlash = append(rowIndexesBackSlash, rowIndex)
		columnIndexesBackSlash = append(columnIndexesBackSlash, 0)
	}
	for columnIndex := range rows[0] {
		rowIndexesBackSlash = append(rowIndexesBackSlash, 0)
		columnIndexesBackSlash = append(columnIndexesBackSlash, columnIndex)
	}

	for i := 0; i < len(rowIndexesBackSlash); i++ {
		rowIndex := rowIndexesBackSlash[i]
		columnIndex := columnIndexesBackSlash[i]

		diagonal := scanBackSlash(rows, rowIndex, columnIndex)

		if len(diagonal) >= 4 {
			diagonalReverse := getReverseString(diagonal)

			scannableLines = append(scannableLines, diagonal, diagonalReverse)
		}
	}

	// '/' Diagonals (& reverse)
	rowIndexesForwardSlash := make([]int, 0)
	columnIndexesForwardSlash := make([]int, 0)

	// Skip 0,0 because the next one covers it
	for columnIndex := range rows[0] {
		rowIndexesForwardSlash = append(rowIndexesForwardSlash, 0)
		columnIndexesForwardSlash = append(columnIndexesForwardSlash, columnIndex)
	}
	for rowIndex := range rows {
		if rowIndex == 0 {
			continue
		}
		rowIndexesForwardSlash = append(rowIndexesForwardSlash, rowIndex)
		columnIndexesForwardSlash = append(columnIndexesForwardSlash, len(rows[0])-1)
	}

	for i := 0; i < len(rowIndexesForwardSlash); i++ {
		rowIndex := rowIndexesForwardSlash[i]
		columnIndex := columnIndexesForwardSlash[i]

		diagonal := scanForwardSlash(rows, rowIndex, columnIndex)

		if len(diagonal) >= 4 {
			diagonalReverse := getReverseString(diagonal)

			scannableLines = append(scannableLines, diagonal, diagonalReverse)
		}
	}

	return scannableLines
}

func task1() {
	fileRows := readFile(FILE_PATH)

	lines := rowsToScannableLines(fileRows)

	xmasCount := 0

	for _, line := range lines {

		xmasCountInString := len(strings.Split(line, "XMAS")) - 1

		xmasCount += xmasCountInString
	}

	fmt.Printf("Task 1: %d\n", xmasCount)
}

func checkForXmas(rows []string, rowIdx int, colIdx int) bool {
	// Check bounds are safe
	if rowIdx-1 < 0 || rowIdx+1 > len(rows)-1 {
		return false
	}
	if colIdx-1 < 0 || colIdx+1 > len(rows[0])-1 {
		return false
	}

	if rows[rowIdx][colIdx] != 'A' {
		return false
	}

	topLeft := rows[rowIdx-1][colIdx-1]
	topRight := rows[rowIdx-1][colIdx+1]
	bottomLeft := rows[rowIdx+1][colIdx-1]
	bottomRight := rows[rowIdx+1][colIdx+1]

	// '\' MAS check
	if topLeft == 'M' {
		if bottomRight != 'S' {
			return false
		}
	} else if topLeft == 'S' {
		if bottomRight != 'M' {
			return false
		}
	} else {
		return false
	}

	// '/' MAS check
	if topRight == 'M' {
		if bottomLeft != 'S' {
			return false
		}
	} else if topRight == 'S' {
		if bottomLeft != 'M' {
			return false
		}
	} else {
		return false
	}

	return true
}

func task2() {
	rows := readFile(FILE_PATH)

	xmasCount := 0

	for rowIdx := range rows {
		for colIdx := range rows[0] {
			hasXmas := checkForXmas(rows, rowIdx, colIdx)

			if hasXmas {
				xmasCount++
			}
		}
	}

	fmt.Printf("Task 2: %d\n", xmasCount)
}

func main() {
	task1()
	task2()
}
