//go:build exclude_tests

package main

import (
	"fmt"
	"os"

	"brainfuck/internal/bf"
)

func main() {
	i, err := bf.New(os.Stdout, os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	i.LoadInstructions("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")

	_ = i.Compile()
	_ = i.VM()
}
