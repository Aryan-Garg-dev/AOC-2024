package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var pl = fmt.Println
var pf = fmt.Printf

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func checkVertical(i, j int, list []string) int {
	var count int = 0
	if i+3 < len(list) {
		if list[i][j] == 'X' && list[i+1][j] == 'M' && list[i+2][j] == 'A' && list[i+3][j] == 'S' {
			count++
		}
	}
	if i-3 >= 0 {
		if list[i][j] == 'X' && list[i-1][j] == 'M' && list[i-2][j] == 'A' && list[i-3][j] == 'S' {
			count++
		}
	}
	return count
}

func checkHorizontal(i, j int, list []string) int {
	var count int = 0
	if j+3 < len(list[i]) {
		if list[i][j] == 'X' && list[i][j+1] == 'M' && list[i][j+2] == 'A' && list[i][j+3] == 'S' {
			count++
		}
	}
	if j-3 >= 0 {
		if list[i][j] == 'X' && list[i][j-1] == 'M' && list[i][j-2] == 'A' && list[i][j-3] == 'S' {
			count++
		}
	}
	return count
}

func checkDiagonal(i, j int, list []string) int {
	var count int = 0
	if j+3 < len(list[i]) && i+3 < len(list) {
		if list[i][j] == 'X' && list[i+1][j+1] == 'M' && list[i+2][j+2] == 'A' && list[i+3][j+3] == 'S' {
			count++
		}
	}
	if j-3 >= 0 && i-3 >= 0 {
		if list[i][j] == 'X' && list[i-1][j-1] == 'M' && list[i-2][j-2] == 'A' && list[i-3][j-3] == 'S' {
			count++
		}
	}
	if j+3 < len(list[i]) && i-3 >= 0 {
		if list[i][j] == 'X' && list[i-1][j+1] == 'M' && list[i-2][j+2] == 'A' && list[i-3][j+3] == 'S' {
			count++
		}
	}
	if j-3 >= 0 && i+3 < len(list) {
		if list[i][j] == 'X' && list[i+1][j-1] == 'M' && list[i+2][j-2] == 'A' && list[i+3][j-3] == 'S' {
			count++
		}
	}
	return count
}

func checkMas(i, j int, list []string) bool {
	if i+2 < len(list) && j+2 < len(list) {
		return (list[i][j] == 'M' && list[i][j+2] == 'S' && list[i+1][j+1] == 'A' && list[i+2][j] == 'M' && list[i+2][j+2] == 'S') ||
			(list[i][j] == 'M' && list[i][j+2] == 'M' && list[i+1][j+1] == 'A' && list[i+2][j] == 'S' && list[i+2][j+2] == 'S') ||
			(list[i][j] == 'S' && list[i][j+2] == 'M' && list[i+1][j+1] == 'A' && list[i+2][j] == 'S' && list[i+2][j+2] == 'M') ||
			(list[i][j] == 'S' && list[i][j+2] == 'S' && list[i+1][j+1] == 'A' && list[i+2][j] == 'M' && list[i+2][j+2] == 'M')
	}
	return false
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)
	input := string(data)
	inputList := strings.Split(strings.TrimSpace(input), "\n")
	totalCount := 0
	masCount := 0
	for i := 0; i < len(inputList); i++ {
		for j := 0; j < len(inputList[i]); j++ {
			count := checkVertical(i, j, inputList) +
				checkDiagonal(i, j, inputList) +
				checkHorizontal(i, j, inputList)
			totalCount += count
			if checkMas(i, j, inputList) {
				masCount++
			}
		}
	}
	pl(totalCount)
	pl(masCount)
}
