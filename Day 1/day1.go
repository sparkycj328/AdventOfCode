package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := make([]int, 0)
	windows := make([]int, 0)
	s = open(s)
	// fmt.Println(s)
	iterateWindows(s, windows)
}

// Opens a file and reads in the numbers line by line
func open(s []int) []int {
	f, err := os.Open("sonar.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	// Splits on newlines
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println(err)
		}
		s = append(s, num)
	}
	return s
}

// Determine how many times number increase
func iterate(s []int) int {

	// Keep track how many times the next number is bigger than the previous
	increased := 0
	for i, _ := range s {
		if i <= 1996 {
			if s[i+1] > s[i] {
				increased += 1
			} else {

			}
		} else {
			break
		}

	}
	return increased
}

func iterateWindows(s []int, windows []int) {

	// Keep track how many times the next number is bigger than the previous
	increased := 0
	for i, _ := range s {
		if i <= 1997 {
			sum := s[i] + s[i+1] + s[i+2]
			windows = append(windows, sum)
			increased += 1
		} else {
			break
		}

	}
	increased = iterate(windows)
	fmt.Println(increased)
}
