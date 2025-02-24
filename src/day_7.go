package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunDay7() {
	input := "input/day_7.txt"
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		result, _ := strconv.Atoi(line[0:strings.Index(line, ":")])
		nums := strings.Split(line[strings.Index(line, ":")+2:], " ")
		if IsCalibrated(nums, result) {
			sum += result
		}
	}
	fmt.Println(sum)
}

func IsCalibrated(nums []string, result int) bool {
	return evaluate(nums, 0, 0, result, true) ||
		evaluate(nums, 0, 1, result, true)
}
func evaluate(nums []string, index int, current int, result int, isAdd bool) bool {
	if index == len(nums) {
		return current == result
	}

	num, _ := strconv.Atoi(nums[index])
	if index == 0 {
		return evaluate(nums, index+1, num, result, true) ||
			evaluate(nums, index+1, num, result, false)
	}

	if isAdd {
		return evaluate(nums, index+1, current+num, result, true) ||
			evaluate(nums, index+1, current+num, result, false)
	} else {
		return evaluate(nums, index+1, current*num, result, true) ||
			evaluate(nums, index+1, current*num, result, false)
	}
}
