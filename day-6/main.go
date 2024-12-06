package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var pl = fmt.Println
var pf = fmt.Printf

func parseInput(fileName string) []string {
	var areaMap []string

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		areaMap = append(areaMap, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return areaMap
}

type vec2 struct {
	x, y int
}

var up = vec2{0, -1}
var down = vec2{0, 1}
var left = vec2{-1, 0}
var right = vec2{1, 0}

func getDirection(arrow byte) vec2 {
	switch arrow {
	case '^':
		return up
	case '>':
		return right
	case 'V':
		return down
	case '<':
		return left
	default:
		return vec2{0, 0}
	}
}

func switchDirection(guardDir *vec2) {
	var newDir vec2
	switch *guardDir {
	case up:
		newDir = right
	case right:
		newDir = down
	case down:
		newDir = left
	case left:
		newDir = up
	default:
		newDir = vec2{0, 0}
	}
	*guardDir = newDir
}

func markPosition(areaMap []string, nextPosition vec2) {
	if nextPosition.x == 0 {
		areaMap[nextPosition.y] = "X" + areaMap[nextPosition.y][nextPosition.x+1:]
	} else if nextPosition.x == len(areaMap[nextPosition.y])-1 {
		areaMap[nextPosition.y] = areaMap[nextPosition.y][:nextPosition.x] + "X"
	} else {
		areaMap[nextPosition.y] = areaMap[nextPosition.y][:nextPosition.x] + "X" + areaMap[nextPosition.y][nextPosition.x+1:]
	}
}

func findGuardStepCount(areaMap []string, guardPos, guardDir vec2) (int, int) {
	var count, loops int
	var nextPosition vec2

	localMap := make([]string, len(areaMap))
	copy(localMap, areaMap)

	for guardPos.y > 0 && guardPos.y < len(areaMap)-1 && guardPos.x > 0 && guardPos.x < len(areaMap[guardPos.y])-1 {
		nextPosition = vec2{guardPos.x + guardDir.x, guardPos.y + guardDir.y}
		if areaMap[nextPosition.y][nextPosition.x] == '#' {
			switchDirection(&guardDir)
			// loops += simulateLoop(areaMap, guardPos, guardDir)
			continue
		} else {
			guardPos.x, guardPos.y = nextPosition.x, nextPosition.y
			if localMap[nextPosition.y][nextPosition.x] != 'X' {
				count++
				markPosition(localMap, nextPosition)
			}
		}
	}
	markPosition(localMap, guardPos)
	return count, loops
}

func isInALoop(areaMap []string, guardPos, guardDir vec2) bool {
	var nextPosition vec2
	var visitedObs = 0
	for guardPos.y > 0 && guardPos.y < len(areaMap)-1 && guardPos.x > 0 && guardPos.x < len(areaMap[guardPos.y])-1 {
		nextPosition = vec2{guardPos.x + guardDir.x, guardPos.y + guardDir.y}
		if areaMap[nextPosition.y][nextPosition.x] == '#' {
			switchDirection(&guardDir)
			continue
		} else if areaMap[nextPosition.y][nextPosition.x] == 'O' {
			switchDirection(&guardDir)
			visitedObs++
			if visitedObs == 2 {
				return true
			}
			continue
		} else {
			guardPos.x, guardPos.y = nextPosition.x, nextPosition.y
		}
	}
	return false
}

func simulateLoop(areaMap []string, guardPos, guardDir vec2) int {
	var count int
	r, c := guardPos.y+guardDir.y, guardPos.x+guardDir.x
	for areaMap[r][c] != '#' {
		originalRow := areaMap[r]
		areaMap[r] = areaMap[r][:c] + "O" + areaMap[r][c+1:]
		if ok := isInALoop(areaMap, guardPos, guardDir); ok {
			count++
		}
		areaMap[r] = originalRow
		r, c = r+guardDir.y, c+guardDir.x
		if !(r >= 0 && r <= len(areaMap)-1 && c >= 0 && c <= len(areaMap[r])-1) {
			break
		}
	}
	return count
}

func main() {
	areaMap := parseInput("input.txt")
	var guardPos vec2
	var guardDir vec2
	for i, row := range areaMap {
		j := strings.IndexAny(row, "^><v")
		if j != -1 {
			guardPos = vec2{j, i}
			guardDir = getDirection(areaMap[i][j])
		}
	}
	count, obstructions := findGuardStepCount(areaMap, guardPos, guardDir)
	pl(count)
	pl(obstructions)

}
