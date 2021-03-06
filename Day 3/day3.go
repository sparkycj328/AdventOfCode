package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type StringArrays struct {
	bitOccurences []int
	gammaString   string
	epsilonString string
	binaryStrings []string
}

func main() {

	arr := StringArrays{}
	arr.open()
	arr.rates()
	arr.openFile()
	arr.oxygen()
}

func (arr *StringArrays) openFile() {
	bytes, err := ioutil.ReadFile("binary.txt")
	if err != nil {
		log.Println(err)
	}
	input := strings.Trim(string(bytes), "\n")
	arr.binaryStrings = strings.Split(input, "\n")

}

//Open() will read in the commands line by line from binary.txt
func (arr *StringArrays) open() {
	size := 12
	// initialize array of 0's, with len = the first line
	arr.bitOccurences = make([]int, size)

	// Opens the instructions.txt file and stores it in f
	f, err := os.Open("binary.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	// Splits on newlines
	scanner := bufio.NewScanner(f)

	// Iterate through each line of f based on newline
	for scanner.Scan() {
		// Further iterate through each line based on the separating space
		v := strings.SplitAfter(scanner.Text(), "")

		for i, s := range v {
			switch s {
			case "1":
				arr.bitOccurences[i] += 1
			case "0":
				arr.bitOccurences[i] -= 1
			}
		}
	}

	fmt.Println(arr.bitOccurences)

	// Handle any errors generated by scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

// Rates() will determine the gamma rates and epsilon rates by determining the binary string
// and then converting thi binary to decimal in order to generate a product of these two
func (arr *StringArrays) rates() {
	for _, bit := range arr.bitOccurences {
		if bit >= 0 {
			arr.gammaString += "1"
			arr.epsilonString += "0"
		} else {
			arr.gammaString += "0"
			arr.epsilonString += "1"
		}

	}
	// fmt.Println(arr.gammaString)
	// fmt.Println(arr.epsilonString)

	gamma, _ := strconv.ParseInt(arr.gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(arr.epsilonString, 2, 64)
	product := gamma * epsilon
	fmt.Println(product)
}

// oxygen() willdetermine the 02 level by scanning bit by bit for the most common integer and storing inputs
// based on if the input contains the most common bit in the specified position
func (arr *StringArrays) oxygen() {
	oArr := arr.binaryStrings
	coArr := arr.binaryStrings

	for i := range arr.binaryStrings[0] {
		if len(oArr) > 1 {
			o2BitOccurences := getBitOccurences(oArr)
			newO2Arr := []string{}

			for _, bits := range oArr {
				if o2BitOccurences[i] >= 0 && bits[i] == '1' {
					newO2Arr = append(newO2Arr, bits)
				} else if o2BitOccurences[i] < 0 && bits[i] == '0' {
					newO2Arr = append(newO2Arr, bits)
				}
			}

			oArr = newO2Arr
		}

		if len(coArr) > 1 {
			co2BitOccurences := getBitOccurences(coArr)
			newCo2Arr := []string{}

			for _, bits := range coArr {
				if co2BitOccurences[i] >= 0 && bits[i] == '0' {
					newCo2Arr = append(newCo2Arr, bits)
				} else if co2BitOccurences[i] < 0 && bits[i] == '1' {
					newCo2Arr = append(newCo2Arr, bits)
				}
			}

			coArr = newCo2Arr
		}
	}
	oArr[0] = strings.Trim(oArr[0], "\r")
	coArr[0] = strings.Trim(coArr[0], "\r")
	o2, err := strconv.ParseInt(oArr[0], 2, 64)
	if err != nil {
		log.Println(err)
	}
	co2, _ := strconv.ParseInt(coArr[0], 2, 64)
	fmt.Println("Part2: o2:", o2, "Co2:", co2, "Lift support:", o2*co2)
}

// Helper function which determines the most occuring bit in each position by adding
// or subtracting 1 from a sum. If negative, 0 is the more common bit, if positive then 1 is most occuring
func getBitOccurences(binaryStrings []string) []int {
	size := len(binaryStrings[0])

	// initialize array of 0's, with len = the first line
	bitOccurences := make([]int, size)

	for _, entry := range binaryStrings {
		for i, bit := range entry {
			switch bit {
			case '1':
				bitOccurences[i] += 1
			case '0':
				bitOccurences[i] -= 1
			}
		}
	}
	return bitOccurences
}
