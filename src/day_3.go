package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func RunDay3() {
	file, err := os.Open("input/day_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum1 := 0
	sum2 := 0
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		matches := mulRe.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			n1, _ := strconv.Atoi(line[match[2]:match[3]])
			n2, _ := strconv.Atoi(line[match[4]:match[5]])
			product := n1 * n2

			mostRecentMatch := GetMostRecentMatch(line[:match[0]])
			if mostRecentMatch == "do" {
				enabled = true
			} else if mostRecentMatch == "don't" {
				enabled = false
			}

			if enabled {
				sum2 += product
			}
			sum1 += product
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
func GetMostRecentMatch(line string) string {
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	doMatches := doRe.FindAllStringIndex(line, -1)
	dontMatches := dontRe.FindAllStringIndex(line, -1)

	var mostRecentMatch string
	var mostRecentPos int
	if len(doMatches) > 0 || len(dontMatches) > 0 {
		if len(doMatches) > 0 && doMatches[len(doMatches)-1][1] > mostRecentPos {
			mostRecentMatch = "do"
			mostRecentPos = doMatches[len(doMatches)-1][1]
		}

		if len(dontMatches) > 0 && dontMatches[len(dontMatches)-1][1] > mostRecentPos {
			mostRecentMatch = "don't"
			mostRecentPos = dontMatches[len(dontMatches)-1][1]
		}
	}

	return mostRecentMatch
}
