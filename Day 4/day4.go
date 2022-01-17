package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Storage struct {
	NumbersCalled []int
	Row           Rows
	boards        Board
}
type Rows struct {
	numbers []int
}

type Board struct {
	rows []Rows
}

func main() {
	s := Storage{}
	file := open()
	file2 := open2()
	s.calledNumbers(file)
	s.creatingBoards(file2)
}

// Open() will open the file containing the data and return it back to main function
func open() *os.File {
	f, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// Open2() will open the file containing the secondary data and return it back to main function
func open2() *os.File {
	f, err := os.Open("./data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// calledNumbers() will receive a file and retrieve the first line of numbers to be called
// storing each number in a slice
func (s *Storage) calledNumbers(f *os.File) {
	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	lines := strings.Split(str, "\n")
	numbersCalled := lines[0]
	numbersCalled = strings.TrimSuffix(numbersCalled, "\r")
	strArr := strings.Split(numbersCalled, ",")

	for _, str := range strArr {
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		s.NumbersCalled = append(s.NumbersCalled, number)
	}

	fmt.Println(s.NumbersCalled)

}

// creatingBoards() will create boards based on a 5x5 basis
func (s *Storage) creatingBoards(f *os.File) {
	// board := make([]int, 0)
	// boards := make([][]int, 0)

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)

	// Create horizontal rows
	lines := strings.Split(str, "\n\n")
	for _, line := range lines {
		rowStr := strings.Fields(line)
		for _, rowInts := range rowStr {
			number, err := strconv.Atoi(rowInts)
			if err != nil {
				log.Fatal(err)
			}
			var numInts []int
			numInts = append(numInts, number)
			row := Rows{numInts}
			s.boards.rows = append(s.boards.rows, row)
		}
	}

	fmt.Println(s.boards.rows)

}
