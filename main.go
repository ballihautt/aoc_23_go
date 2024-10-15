package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ballihautt/aoc_23_go/days"
)

var (
	inputFile string
	testMode  bool
	exercise  uint
)

func init() {
	flag.StringVar(&inputFile, "f", "", "Path to the input file.")
	flag.BoolVar(&testMode, "t", false, "Enable test mode for default path.")
	flag.UintVar(&exercise, "d", 1, "The number of the day you want to solve.")
	flag.Parse()

	if inputFile == "" { // no input file, switch to default
		if testMode {
			inputFile = "./tinput/" // default path for test mode
		} else {
			inputFile = "./input/" // default path for inputs
		}
		if exercise < 10 { // paddle with 0 as each file is dayXX
			inputFile += fmt.Sprintf("day0%d", exercise)
		} else {
			inputFile += fmt.Sprintf("day%d", exercise)
		}
		fmt.Printf("Using default path: %s\n", inputFile) // notifying user
	}
}

func main() {

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading `%s`: %v", inputFile, err)
		return
	}
	input := string(inputBytes)

	switch exercise {
	case 1:
		days.Day01(input)
	case 2:
		days.Day02(input)
	case 3:
		days.Day03(input)
	case 4:
		days.Day04(input)
	case 5:
		days.Day05(input)
	case 6:
		days.Day06(input)
	default:
		fmt.Printf("Day number %d is not implemented yet\n", exercise)
	}
}
