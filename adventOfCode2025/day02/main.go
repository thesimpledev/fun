package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateRange(start, end int) (output int) {
	for i := start; i <= end; i++ {
		str := strconv.Itoa(i)

		if isInvalid(str) {
			output += i
		}
	}

	return
}

func isInvalid(s string) bool {
	n := len(s)

	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		reps := n / patternLen

		if strings.Repeat(pattern, reps) == s {
			return true
		}
	}
	return false
}

func parse(input string) (output [][]int, err error) {
	sets := strings.Split(input, ",")
	output = make([][]int, 0)

	for _, pair := range sets {
		nums := strings.Split(pair, "-")
		if len(nums) != 2 {
			err = errors.New("must have at 2 numbers")
			return
		}

		start, strerr := strconv.Atoi(strings.TrimSpace(nums[0]))
		if strerr != nil {
			err = fmt.Errorf("unable to parse %s: %w", nums[0], err)
			return
		}

		end, strerr := strconv.Atoi(strings.TrimSpace(nums[1]))
		if strerr != nil {
			err = fmt.Errorf("unable to parse %s: %w", nums[1], err)
		}

		output = append(output, []int{start, end})
	}

	return
}

func loader(name string) (output string, err error) {
	file, ferr := os.ReadFile(name)
	if ferr != nil {
		err = fmt.Errorf("unable to open %s: %w", name, ferr)
		return
	}

	output = string(file)
	return
}

func manager(name string) int {
	data, err := loader(name)
	if err != nil {
		panic(err)
	}

	intSets, err := parse(data)
	if err != nil {
		panic(err)
	}
	total := 0
	for _, set := range intSets {
		total += validateRange(set[0], set[1])
	}

	return total
}

func main() {
	fmt.Println(manager("data_real"))
}
