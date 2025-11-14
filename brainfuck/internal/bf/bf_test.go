package bf

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNew(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	tapeSizeWant := 30_000
	tapePosition := 0
	if i == nil {
		t.Fatal("client should not be null")
	}

	if i.tape == nil {
		t.Fatal("tape should not be null")
	}

	if len(i.tape) != tapeSizeWant {
		t.Errorf("tape is %d and should be %d", len(i.tape), tapeSizeWant)
	}

	if i.pos != tapePosition {
		t.Errorf("tape position got %d want %d", i.pos, tapePosition)
	}
}

func TestLoadInstructions(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	want := ">>>>>"
	i.LoadInstructions(want)

	if want != i.instructions {
		t.Errorf("want %s, go %s", want, i.instructions)
	}
}

func TestExecuteInstructionsError(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.LoadInstructions("")
	err = i.ExecutionLoop()
	if err == nil {
		t.Error("expected error for empty instructions, got nil")
	}
}

func TestExecuteInstructionsShiftRight(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	inst := ">x>"
	want := 2
	i.LoadInstructions(inst)
	err = i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftRight(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	want := 5
	for range want {
		i.shiftRight()
	}

	if i.pos != want {
		t.Errorf("got position %d want position %d", i.pos, want)
	}
}

func TestExecuteInstructionsShiftLeft(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	inst := ">x><"
	want := 1
	i.LoadInstructions(inst)
	err = i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftLeft(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	want := 3
	for range 5 {
		i.shiftRight()
	}

	for range 2 {
		i.shiftLeft()
	}

	if i.pos != want {
		t.Errorf("got position %d want position %d", i.pos, want)
	}
}

func TestShiftLeftNegative(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.shiftLeft()
	want := 29_999

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestShiftRightOverflow(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	want := 0
	i.pos = 29_999
	i.shiftRight()

	if i.pos != want {
		t.Errorf("got %d, want %d", i.pos, want)
	}
}

func TestExecuteInstructionsIncrement(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	inst := ">+"
	want := byte(1)
	i.LoadInstructions(inst)
	err = i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error got error")
	}

	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestIncrement(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.increment()
	want := byte(1)
	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestExecuteInstructionsDecrement(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	inst := ">-"
	want := byte(255)
	i.LoadInstructions(inst)
	err = i.ExecutionLoop()
	if err != nil {
		t.Errorf("expected no error got error %v", err)
	}

	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestDecrement(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.decrement()
	want := byte(255)
	if i.tape[i.pos] != want {
		t.Errorf("got %d, want %d", i.tape[i.pos], want)
	}
}

func TestOutput(t *testing.T) {
	buffer := &bytes.Buffer{}
	i, err := New(buffer, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.decrement()
	want := string(byte(255))
	i.output()
	got := buffer.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExecuteInstructionOutput(t *testing.T) {
	buffer := &bytes.Buffer{}
	i, err := New(buffer, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.LoadInstructions(">-.")
	want := string(byte(255))
	err = i.ExecutionLoop()
	if err != nil {
		t.Errorf("expected no error got error %v", err)
	}
	got := buffer.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestParserLoop(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		want         string
	}{
		{
			name:         "skip loop output B",
			instructions: "[impossible]+++++++++++[>++++++<-]>.",
			want:         "B",
		},
		{
			name:         "output A",
			instructions: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			want:         "Hello World!\n",
		},
		{
			name:         "output 5",
			instructions: "+++++[>++++++++++<-]>+++.",
			want:         "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			i, err := New(buffer, strings.NewReader("a"))
			if err != nil {
				t.Errorf("error constructing new interperter %v", err)
			}
			i.LoadInstructions(tt.instructions)
			err = i.ParserLoop()
			if err != nil {
				t.Error("expected no error, got error")
			}
			err = i.ExecutionLoop()
			if err != nil {
				t.Error("expected no error, got error")
			}
			if buffer.String() != tt.want {
				t.Errorf("got %q want %q", buffer.String(), tt.want)
			}
		})
	}
}

func TestClearInstructions(t *testing.T) {
	buffer := &bytes.Buffer{}
	i, err := New(buffer, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.LoadInstructions("+++++++++")
	i.ClearInstructions()
	err = i.ParserLoop()
	if err == nil {
		t.Error("Error is nil and should not be")
	}
}

func TestReadChar(t *testing.T) {
	buffer := &bytes.Buffer{}
	want := "a"
	wantRune, _ := utf8.DecodeRuneInString(want)
	input := strings.NewReader(want)

	i, err := New(buffer, input)
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	char, err := i.readChar()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if char != wantRune {
		t.Errorf("expected %c got %c", wantRune, char)
	}
}

func TestCollectUserInput(t *testing.T) {
	buffer := &bytes.Buffer{}
	want := "a"
	input := strings.NewReader(want)
	inst := ",."
	i, err := New(buffer, input)
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.LoadInstructions(inst)
	err = i.ParserLoop()
	if err != nil {
		t.Error("expected no error, got error")
	}
	err = i.ExecutionLoop()
	if err != nil {
		t.Error("expected no error, got error")
	}
	got := buffer.String()
	if got != want {
		t.Errorf("got %s - want %s", got, want)
	}
}

func TestUnevenLoopsError(t *testing.T) {
	tests := []struct {
		name string
		inst string
		err  bool
	}{
		{
			name: "extra [",
			inst: "[[]",
			err:  true,
		},
		{
			name: "extra ]",
			inst: "[]]",
			err:  true,
		},
		{
			name: "no error loop",
			inst: "[[[[[[[[[]]]]]]]]]",
			err:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
			if err != nil {
				t.Errorf("error constructing new interperter %v", err)
			}
			i.LoadInstructions(tt.inst)
			err = i.ParserLoop()
			if err == nil && tt.err {
				t.Error("expected error not none")
			}
		})
	}
}

func TestFlushStackAndLoops(t *testing.T) {
	i, err := New(&bytes.Buffer{}, strings.NewReader("a"))
	if err != nil {
		t.Errorf("error constructing new interperter %v", err)
	}
	i.LoadInstructions("[[[[[[[][[][[")
	_ = i.ParserLoop()
	i.ClearInstructions()
	i.LoadInstructions("[]")
	err = i.ParserLoop()
	if err != nil {
		t.Error("Got error expected none")
	}
}

func TestNilReader(t *testing.T) {
	_, err := New(&bytes.Buffer{}, nil)
	if err == nil {
		t.Error("expected error got none")
	}
}

func TestNilWriter(t *testing.T) {
	_, err := New(nil, strings.NewReader(""))
	if err == nil {
		t.Error("expected error got none")
	}
}

type errorReader struct{}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func TestReadError(t *testing.T) {
	i, err := New(&bytes.Buffer{}, errorReader{})
	if err != nil {
		t.Fatal("failed to created interperter")
	}
	i.LoadInstructions(",.")
	_ = i.ParserLoop()
	err = i.ExecutionLoop()
	if err == nil {
		t.Error("expexted error, not none")
	}
}
