package src

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunDay1() {
	filename := "input/day_1.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	leftlist := []int{}
	rightlist := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		leftNum, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatal(err)
		}
		rightNum, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}

		leftlist = append(leftlist, leftNum)
		rightlist = append(rightlist, rightNum)
	}

	fmt.Println(day1_1(leftlist, rightlist))
	fmt.Println(day1_2(leftlist, rightlist))
}

func day1_1(leftlist []int, rightlist []int) int {
	sort.Ints(leftlist)
	sort.Ints(rightlist)

	sum := 0
	for i := 0; i < len(leftlist); i++ {
		distance := int(math.Abs(float64(leftlist[i] - rightlist[i])))
		sum += distance
	}

	return sum
}

func day1_2(leftlist []int, rightlist []int) int {
	sum := 0
	for i := 0; i < len(leftlist); i++ {
		elem := leftlist[i]
		count := 0
		for j := 0; j < len(rightlist); j++ {
			if rightlist[j] == elem {
				count++
			}
		}
		sum += elem * count
	}
	return sum
}
