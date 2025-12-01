package main

import "testing"

func TestLeft(t *testing.T) {
	tests := []struct {
		name  string
		inst  int
		want  int
		start int
	}{
		{
			name:  "L1",
			inst:  1,
			want:  49,
			start: 50,
		},
		{
			name:  "Zero",
			inst:  50,
			want:  0,
			start: 50,
		},
		{
			name:  "wrap around to 75",
			inst:  75,
			want:  75,
			start: 50,
		},
		{
			name:  "L68 rotation",
			inst:  68,
			want:  82,
			start: 50,
		},
		{
			name:  "L30 rotation",
			inst:  30,
			want:  52,
			start: 82,
		},
		{
			name:  "huge",
			inst:  300,
			want:  50,
			start: 50,
		},
	}

	for _, tt := range tests {
		s := new(tt.start)
		t.Run(tt.name, func(t *testing.T) {
			got := s.left(tt.inst)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name  string
		inst  int
		want  int
		start int
	}{
		{
			name:  "R1",
			inst:  1,
			want:  51,
			start: 50,
		},
		{
			name:  "Zero",
			inst:  50,
			want:  0,
			start: 50,
		},
		{
			name:  "25",
			inst:  75,
			want:  25,
			start: 50,
		},
		{
			name:  "R48 from 52",
			inst:  48,
			want:  0,
			start: 52,
		},
		{
			name:  "large",
			inst:  550,
			want:  0,
			start: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(tt.start)
			got := s.right(tt.inst)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	s := new(50)
	err := s.load("data")
	if err != nil {
		t.Errorf("Unable to load file %s", err.Error())
	}
}

func TestParser(t *testing.T) {
	s := new(50)
	want := 3
	err := s.load("data")
	if err != nil {
		t.Errorf("unable to load file %s", err.Error())
	}

	got, err := s.parse()
	if err != nil {
		t.Error("got error expected none")
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "normal",
			num:  40,
			want: 40,
		},
		{
			name: "high by 1",
			num:  150,
			want: 50,
		},
		{
			name: "hight by 5",
			num:  550,
			want: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(tt.num)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
