package main

import (
	"reflect"
	"testing"
)

func TestBankParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "third",
			input: "234234234234278",
			want:  434234234278,
		},
		{
			name:  "second",
			input: "811111111111119",
			want:  811111111119,
		},
		{
			name:  "First",
			input: "987654321111111",
			want:  987654321111,
		},
		{
			name:  "fourth",
			input: "818181911112111",
			want:  888911112111,
		},
		{
			name:  "edge",
			input: "11111111111129",
			want:  111111111129,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBank(tt.input)
			if err != nil {
				t.Error("got error want none")
			}

			if got != tt.want {
				t.Errorf("\ns: %s \nw: %d \ng: %d", tt.input, tt.want, got)
			}
		})
	}
}

func TestLoadFile(t *testing.T) {
	name := "data"
	want := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	got, err := loadFile(name)
	if err != nil {
		t.Error("got error want none")
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \n%v \nwant \n%v", got, want)
	}
}

func TestParseData(t *testing.T) {
	data := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	want := 3121910778619
	got, err := parseData(data)
	if err != nil {
		t.Errorf("got error expected none")
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
