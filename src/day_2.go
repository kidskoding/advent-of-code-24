package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunDay2() {
	filepath := "input/day_2.txt"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count_1 := 0
	count_2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		lst := []int{}

		for _, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				log.Fatal(err)
			}
			lst = append(lst, num)
		}

		if Day2_1(lst) {
			count_1++
		}

		if Day2_2(lst) {
			count_2++
		}
	}

	fmt.Println(count_1)
	fmt.Println(count_2)
}

func Day2_1(lst []int) bool {
	return isSafe(lst, true) || isSafe(lst, false)
}

func Day2_2(lst []int) bool {
	if isSafe(lst, true) || isSafe(lst, false) {
		return true
	}
	for i := 0; i < len(lst); i++ {
		temp := append([]int{}, lst[:i]...)
		temp = append(temp, lst[i+1:]...)
		if isSafe(temp, true) || isSafe(temp, false) {
			return true
		}
	}
	return false
}

func isSafe(lst []int, inc bool) bool {
	for i := 0; i < len(lst)-1; i++ {
		diff := lst[i+1] - lst[i]
		if inc {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || -diff > 3 {
				return false
			}
		}
	}
	return true
}
