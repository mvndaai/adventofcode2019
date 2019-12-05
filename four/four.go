package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func intToSlice(n int) []int {
	var r []int
	for {
		if n == 0 {
			break
		}
		r = append([]int{n % 10}, r...)
		n = n / 10
	}
	return r
}

func validNumber(n int) bool {
	s := intToSlice(n)

	sideByside := false
	// fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		// fmt.Println(s[i], s[i+1], s[i] == s[i+1], s[i] > s[i+1])

		if s[i] == s[i+1] {
			sideByside = true
		}
		if s[i] > s[i+1] {
			// fmt.Println("ending here")
			return false
		}
	}
	if !sideByside {
		// fmt.Println("not side by side")
		return false
	}

	return true
}

func validNumber2(n int) bool {
	s := intToSlice(n)

	sideByside := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}

		if s[i] == s[i+1] {
			t := true
			if i-1 >= 0 {
				if s[i-1] == s[i] {
					t = false
				}
			} else {
				log.Println("a", i, n)
			}
			if i+2 < len(s) {
				if s[i+2] == s[i] {
					t = false
				}
			}

			if t {
				sideByside = t
			}
		}
	}

	if !sideByside {
		return false
	}

	return true
}

func main() {
	input := "137683-596253"
	i := strings.Split(input, "-")

	start, err := strconv.Atoi(i[0])
	if err != nil {
		log.Fatal("could not convert first input", err)
	}
	end, err := strconv.Atoi(i[1])
	if err != nil {
		log.Fatal("could not convert first input", err)
	}

	count := 0
	for i := start; i <= end; i++ {
		if validNumber2(i) {
			count++
		}
	}
	fmt.Println(count)
	// Failed with 1065 (too low)
	// Failed with 1372 (too high)
}
