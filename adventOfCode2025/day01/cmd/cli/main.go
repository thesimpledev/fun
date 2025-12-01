package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const max = 99

type safe struct {
	dial         int
	combination  int
	instructions [][]byte
}

func new(start int) safe {
	return safe{
		dial:         start,
		instructions: make([][]byte, 0),
	}
}

func normalize(n int) int {
	return n % 100
}

func (s *safe) left(n int) int {
	n = normalize(n)
	temp := s.dial - n
	if temp >= 0 {
		s.dial = temp
		return temp
	}

	s.dial = max + temp + 1
	return s.dial
}

func (s *safe) right(n int) int {
	n = normalize(n)
	temp := s.dial + n
	if temp <= 99 {
		s.dial = temp
		return s.dial
	}
	s.dial = temp - max - 1
	return s.dial
}

func (s *safe) load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("unable to open file %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s.instructions = append(s.instructions, slices.Clone(scanner.Bytes()))
	}

	return nil
}

func (s *safe) parse() (int, error) {
	if s.instructions == nil {
		return 0, errors.New("instructions must be set first")
	}

	for _, instruction := range s.instructions {
		direction := instruction[0]
		amount, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			return 0, fmt.Errorf("unable to parse %b: %w", instruction, err)
		}
		if direction == 'L' {
			num := s.left(amount)
			if num == 0 {
				s.combination++
			}
			continue
		}
		num := s.right(amount)
		if num == 0 {
			s.combination++
		}

	}

	return s.combination, nil
}

func main() {
	safe := new(50)
	safe.load("datareal")
	combo, err := safe.parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The Combination is: ", combo)
}
