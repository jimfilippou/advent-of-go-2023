package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	ReplaceWordsWithDigits bool
}

func getCalibration(line string) int {

	// First and last are digits O(1)
	firstAndLastAreDigits, _ := regexp.MatchString(`^\d(\w+)\d$`, string(line))
	if firstAndLastAreDigits {
		concatenated := string(line[0]) + string(line[len(line)-1])
		value, _ := strconv.Atoi(concatenated)
		return value
	}

	// Start looking for a digit match, at this point I cant predict the complexity, I think it's O(n)
	leftIterator, rightIterator := 0, len(line)-1
	leftMatch, _ := regexp.MatchString(`^[0-9]$`, string(line[leftIterator]))
	rightMatch, _ := regexp.MatchString(`^[0-9]$`, string(line[rightIterator]))

	for {
		// If both are digits, return
		if leftMatch == true && rightMatch == true {
			concatenated := string(line[leftIterator]) + string(line[rightIterator])
			value, _ := strconv.Atoi(concatenated)
			return value
		}

		// If for some reason the iterators meet, check if it's a digit and return
		if leftIterator == rightIterator {
			isNumber, _ := regexp.MatchString(`^[0-9]$`, string(line[leftIterator]))
			if !isNumber {
				panic("The iterators met but the value is not a number, are you drunk?")
			}
			concatenated := string(line[leftIterator]) + string(line[rightIterator])
			number, _ := strconv.Atoi(concatenated)
			return number
		}

		// If the left is not a digit, move it to the right
		if !leftMatch {
			leftIterator++
			leftMatch, _ = regexp.MatchString(`^\d$`, string(line[leftIterator]))
		}

		// If the right is not a digit, move it to the left
		if !rightMatch {
			rightIterator--
			rightMatch, _ = regexp.MatchString(`^\d$`, string(line[rightIterator]))
		}
	}
}

func main() {
	// Load configuration
	var conf Configuration
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatalf("unable to load configuration: %v", err)
		return
	}

	fmt.Println("Advent of Code 2023 - Day 1")
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		return
	}

	// Split the lines
	lines := strings.Split(string(file), "\n")

	// Set the total
	total := 0

	// Iterate over the lines
	for _, line := range lines {
		fmt.Printf(`
			Line before: %s
			Calibration before: %d
		`, line, getCalibration(line))
		if conf.ReplaceWordsWithDigits {
			line = strings.ReplaceAll(line, "one", "o1e")
			line = strings.ReplaceAll(line, "two", "t2o")
			line = strings.ReplaceAll(line, "three", "t3e")
			line = strings.ReplaceAll(line, "four", "4")
			line = strings.ReplaceAll(line, "five", "5e")
			line = strings.ReplaceAll(line, "six", "6")
			line = strings.ReplaceAll(line, "seven", "7n")
			line = strings.ReplaceAll(line, "eight", "e8t")
			line = strings.ReplaceAll(line, "nine", "n9e")
		}
		calibration := getCalibration(line)
		fmt.Printf(`
			Line after: %s
			Calibration after: %d
		`, line, calibration)
		total += calibration
	}

	fmt.Println("(╯°□°）╯︵ ┻━┻ Answer is: ", total)
}
