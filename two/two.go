package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2019/day/2

func intCode(list []int) []int {
	copy := make([]int, len(list))
	for i, v := range list {
		copy[i] = v
	}
	list = copy

	pos := 0
	for pos < len(list) {
		opcode := list[pos]
		if opcode == 99 {
			return list
		}

		var f int
		a := list[pos+1]
		b := list[pos+2]
		c := list[pos+3]

		switch opcode {
		case 1:
			f = list[a] + list[b]
		case 2:
			f = list[a] * list[b]
		default:
			log.Fatal("bad opt code", opcode)
		}

		list[c] = f
		pos += 4

	}
	log.Fatal("out of the loop")
	return list
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, ",")

		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// list[1] = 12
	// list[2] = 2
	// fmt.Println(list)

	mem := list[0]
	wantedOutput := 19690720

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			list[0] = mem
			list[1] = noun
			list[2] = verb

			ic := intCode(list)
			if ic[0] == wantedOutput {
				fmt.Printf("%02d%02d\n", noun, verb)
				return
			}
		}
	}

	fmt.Println("did not find an answer")

}
