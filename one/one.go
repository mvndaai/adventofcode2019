package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return mass/3 - 2

}

func fuel2(mass int) int {
	m := fuel(mass)
	if m < 0 {
		m = 0
	}
	if m > 0 {
		m += fuel2(m)
	}
	return m
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
		// sum += fuel(i)
		sum += fuel2(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
