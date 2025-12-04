package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parseBank(input string) (high int, rerr error) {
	nums := make([]int, len(input))
	for i := range input {
		v := rune(input[i])
		if unicode.IsDigit(v) {
			nums[i] = int(v - '0')
		}
	}

	changed := true

	for changed {
		length := len(nums)
		changed = false
		for pos := 0; pos < length-1; pos++ {
			if length == 12 {
				break
			}

			if nums[pos] < nums[pos+1] {
				nums = append(nums[:pos], nums[pos+1:]...)
				changed = true
				break
			}

		}
	}

	nums = nums[:12]
	for _, n := range nums {
		high = high*10 + n
	}

	return
}

func loadFile(name string) (data []string, err error) {
	file, ferr := os.Open(name)
	if ferr != nil {
		err = fmt.Errorf("unable to open file %s - %w", name, ferr)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return
}

func parseData(data []string) (total int, err error) {
	for _, str := range data {
		val, perr := parseBank(str)
		if perr != nil {
			err = fmt.Errorf("failed to parse %s %w", str, perr)
			return
		}
		total += val
	}
	return
}

func main() {
	data, _ := loadFile("real_data")
	count, _ := parseData(data)
	fmt.Println(count)
}
