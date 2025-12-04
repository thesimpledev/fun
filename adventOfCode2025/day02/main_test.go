package main

import (
	"reflect"
	"testing"
)

func TestSplitID(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "11-22 bad",
			input: []int{11, 22},
			want:  33,
		},
		{
			name:  "95-115",
			input: []int{95, 115},
			want:  210,
		},
		{
			name:  "998-1012",
			input: []int{998, 1012},
			want:  2009,
		},
		{
			name:  "1188511880-1188511890",
			input: []int{1188511880, 1188511890},
			want:  1188511885,
		},
		{
			name:  "222220-222224",
			input: []int{222220, 222224},
			want:  222222,
		},
		{
			name:  "1698522-1698528",
			input: []int{1698522, 1698528},
			want:  0,
		},
		{
			name:  "446443-446449",
			input: []int{446443, 446449},
			want:  446446,
		},
		{
			name:  "38593856-38593862",
			input: []int{38593856, 38593862},
			want:  38593859,
		},
		{
			name:  "2121212118-2121212124",
			input: []int{2121212118, 2121212124},
			want:  2121212121,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateRange(tt.input[0], tt.input[1])
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name:  "11-22, 95-115",
			input: "11-22, 95-115",
			want:  [][]int{{11, 22}, {95, 115}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.input)
			if err != nil {
				t.Error("got error want none")
				return
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestLoader(t *testing.T) {
	fileName := "data_test"
	want := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124\n"

	got, err := loader(fileName)
	if err != nil {
		t.Error("got error want no error")
	}

	if got != want {
		t.Errorf("got \n%s want \n%s", got, want)
	}
}

func TestManager(t *testing.T) {
	fileName := "data_test"
	want := 4174379265

	got := manager(fileName)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestIsInvalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "111111",
			input: "111111",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInvalid(tt.input)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
