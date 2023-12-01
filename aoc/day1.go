package aoc

/* Part 2 of this problem could be solved with a regexp
 * that contains a positive lookahead. Numeric strings can
 * overlap (e.g. "sevenine" 79) and a normal regexp will
 * not be able to match all strings.
 *
 * (?=(\d|one|two|three|four|five|six|seven|eight|nine))
 *
 * Golang stdlib engine for regular expressions (RE2) does
 * not support lookaheads. Instead of fetching a third-party
 * library I have implemented this with a simple manual search.
 */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func toDigit(in byte) int {
	return int(in - '0')
}

var (
	ReNumber = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
)

func tryGetNumber(in string, out *int) bool {
	matches := ReNumber.FindAllString(in, 1)
	if len(matches) == 0 {
		return false
	}
	switch matches[0] {
	case "one":
		*out = 1
	case "two":
		*out = 2
	case "three":
		*out = 3
	case "four":
		*out = 4
	case "five":
		*out = 5
	case "six":
		*out = 6
	case "seven":
		*out = 7
	case "eight":
		*out = 8
	case "nine":
		*out = 9
	default:
		*out = toDigit(matches[0][0])
	}
	return true
}

func Day1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalFirst := 0
	totalSecond := 0

	first := regexp.MustCompile(`\d`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := first.FindAllString(line, -1)
		totalFirst += toDigit(matches[0][0])*10 + toDigit(matches[len(matches)-1][0])

		var number int
		if tryGetNumber(line, &number) {
			totalSecond += number * 10
		} else {
			log.Fatalf("Number not found in %v", line)
		}

		for i := len(line) - 1; i >= 0; i-- {
			if tryGetNumber(line[i:], &number) {
				totalSecond += number
				break
			}
		}
	}
	fmt.Println("Day1.1:", totalFirst)
	fmt.Println("Day1.2:", totalSecond)
}
