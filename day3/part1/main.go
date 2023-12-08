package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type symbol struct {
	symbol    string
	rowNumber int
	colNumber int
}

func main() {
	var engineSchematic [][]string

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var tokens []string
		for _, r := range line {
			tokens = append(tokens, string(r))
		}
		engineSchematic = append(engineSchematic, tokens)
	}

	symbols := identify_symbols(engineSchematic)
	sumOfAllPartNumbers := sumAdjacentNumbers(engineSchematic, symbols)

	fmt.Println(sumOfAllPartNumbers)
}

func identify_symbols(engineSchematic [][]string) map[symbol]bool {
	numRows, numCols := len(engineSchematic), len(engineSchematic[0])

	symbols := map[symbol]bool{}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			token := engineSchematic[row][col]
			if _, err := strconv.Atoi(token); err != nil && token != "." {
				symbol := symbol{
					symbol:    token,
					rowNumber: row,
					colNumber: col,
				}
				symbols[symbol] = true
			}
		}
	}

	return symbols
}

func sumAdjacentNumbers(engineSchematic [][]string, symbols map[symbol]bool) int {
	numRows, numCols := len(engineSchematic), len(engineSchematic[0])
	sum := 0
	counted := make(map[string]bool)

	for sym := range symbols {
		for row := max(0, sym.rowNumber-1); row <= min(numRows-1, sym.rowNumber+1); row++ {
			for col := max(0, sym.colNumber-1); col <= min(numCols-1, sym.colNumber+1); col++ {
				if row == sym.rowNumber && col == sym.colNumber {
					continue // Skip the symbol itself
				}

				if _, err := strconv.Atoi(engineSchematic[row][col]); err == nil {
					fullNumber, posID := findFullNumber(engineSchematic, row, col, numRows, numCols)
					if !counted[posID] {
						num, _ := strconv.Atoi(fullNumber)
						sum += num
						counted[posID] = true
					}
				}
			}
		}
	}

	return sum
}

func findFullNumber(engineSchematic [][]string, row, col, numRows, numCols int) (string, string) {
	number := engineSchematic[row][col]
	left, right := col, col

	for left > 0 && isDigit(engineSchematic[row][left-1]) {
		left--
		number = engineSchematic[row][left] + number
	}

	for right < numCols-1 && isDigit(engineSchematic[row][right+1]) {
		right++
		number = number + engineSchematic[row][right]
	}

	posID := fmt.Sprintf("%d-%d:%d", row, left, right)
	return number, posID
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
