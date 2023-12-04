package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type buckets struct {
	red   int
	blue  int
	green int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	powerSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := scanner.Text()
		colonIndex := strings.Index(game, ":")
		powerSum += getSetPower(game[colonIndex+1:])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(powerSum)
}

func getGameId(gameStr string) int {
	gameStrSplit := strings.Split(gameStr, " ")
	if num, err := strconv.Atoi(gameStrSplit[1]); err == nil {
		return num
	}
	return -1
}

func getSetPower(gameSets string) int {
	cubeColors := buckets{
		red:   1,
		blue:  1,
		green: 1,
	}

	cubesSets := strings.Split(gameSets, "; ")
	for _, cubeSet := range cubesSets {
		trimmedCubeSet := strings.TrimSpace(cubeSet)
		colorBuckets := strings.Split(trimmedCubeSet, ", ")
		updatedColors := countCubes(colorBuckets)
		cubeColors.red = max(updatedColors.red, cubeColors.red)
		cubeColors.blue = max(updatedColors.blue, cubeColors.blue)
		cubeColors.green = max(updatedColors.green, cubeColors.green)
	}

	return cubeColors.red * cubeColors.blue * cubeColors.green
}

func countCubes(colorBuckets []string) buckets {
	colors := buckets{
		red:   0,
		blue:  0,
		green: 0,
	}

	for _, cubes := range colorBuckets {
		cubeAndColor := strings.Split(cubes, " ")
		numCubes, err := strconv.Atoi(cubeAndColor[0])
		if err != nil {
			fmt.Println(err)
		}
		switch cubeAndColor[1] {
		case "blue":
			colors.blue += numCubes
		case "red":
			colors.red += numCubes
		case "green":
			colors.green += numCubes
		}

	}

	return colors
}
