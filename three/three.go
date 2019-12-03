package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func findPath(path []string) []point {
	var final []point

	current := point{x: 0, y: 0}

	for _, p := range path {
		direction := p[0]
		i, err := strconv.Atoi(p[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch string(direction) {
		case "U":
			for j := 0; j < i; j++ {
				current.y++
				final = append(final, current)
			}
		case "D":
			for j := 0; j < i; j++ {
				current.y--
				final = append(final, current)
			}
		case "L":
			for j := 0; j < i; j++ {
				current.x++
				final = append(final, current)
			}
		case "R":
			for j := 0; j < i; j++ {
				current.x--
				final = append(final, current)
			}
		default:
			log.Fatal("direction not found")
		}

	}

	return final
}

func findManhatten(a, b []point) int {
	// fmt.Println("a", a)
	// fmt.Println("b", b)

	var matches []int
	for _, ap := range a {
		for _, bp := range b {
			// fmt.Println("x", ap.x, bp.x, ap.x == bp.x, "y", ap.y, bp.y, ap.y == bp.y)
			if ap.x == bp.x && ap.y == bp.y {
				// fmt.Println("x", ap.x, bp.x, ap.x == bp.x, "y", ap.y, bp.y, ap.y == bp.y)
				f := math.Abs(float64(ap.x)) + math.Abs(float64(ap.y))
				matches = append(matches, int(f))
			}
		}
	}

	var smallest int

	if len(matches) != 0 {
		smallest = matches[0]
		for _, m := range matches {
			if m < smallest {
				smallest = m
			}
			fmt.Println("smallest", smallest)
		}
	}

	return smallest
}

func manhatten(a, b string) int {
	pathA := strings.Split(a, ",")
	pathB := strings.Split(b, ",")

	foundA := findPath(pathA)
	foundB := findPath(pathB)

	return findManhatten(foundA, foundB)
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	one := manhatten(lines[0], lines[1])
	fmt.Println(one)
}
