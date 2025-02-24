package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunDay5() {
	filepath := "input/day_5.txt"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		words := strings.Split(line, "|")
		if len(words) != 2 {
			log.Fatal("Invalid input")
		}

		num_1, err := strconv.Atoi(strings.TrimSpace(words[0]))
		if err != nil {
			log.Fatal(err)
		}
		num_2, err := strconv.Atoi(strings.TrimSpace(words[1]))
		if err != nil {
			log.Fatal(err)
		}

		temp := []int{num_1, num_2}
		rules = append(rules, temp)
	}

	sum := 0
	sum_2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		nums := []int{}

		for _, word := range words {
			num, err := strconv.Atoi(strings.TrimSpace(word))
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		if Day5_1(nums, rules) {
			sum += nums[len(nums)/2]
		} else {
			sort.Slice(nums, func(i, j int) bool {
				for _, rule := range rules {
					if nums[i] == rule[0] && nums[j] == rule[1] {
						return true
					}
				}
				return false
			})
			sum_2 += nums[len(nums)/2]
		}
	}

	fmt.Println(sum)
	fmt.Println(sum_2)
}

func Day5_1(nums []int, rules [][]int) bool {
	for i := 0; i < len(nums)-1; i++ {
		num1 := nums[i]
		num2 := nums[i+1]
		found := false
		for _, rule := range rules {
			if num1 == rule[0] && num2 == rule[1] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
