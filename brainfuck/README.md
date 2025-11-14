# Brainfuck Interpreter

A Brainfuck interpreter implementation in Go, built using Test-Driven Development.

## What is Brainfuck?

Brainfuck is an esoteric programming language with only 8 commands, operating on a 30,000-byte tape with a movable pointer.

## Why?

So why would I write this? I have used unit testing for years in my code. However, a former Director of Engineering always told me I needed to start doing TDD (Test Driven Development). I did not get the differences between writing my tests before or after my code. However, reading Modern Software Engineering by David Farley I came to understand the difference. When you start out using TDD you keep your code loosely coupled and your tests behavior driven and not implementation driven. Also instead of DDD (Domain Driven Design) seeming like work, it seems to naturally flow with TDD.

## Implementation Details

- 30,000-byte tape with wrapping at boundaries
- Two-phase execution: parsing (bracket matching) and execution
- Stack-based bracket matching during parse phase
- Jump table (map) for O(1) loop jumps during execution

## Commands

| Command | Description                                           |
| ------- | ----------------------------------------------------- |
| `>`     | Move pointer right                                    |
| `<`     | Move pointer left                                     |
| `+`     | Increment cell value                                  |
| `-`     | Decrement cell value                                  |
| `.`     | Output current cell as ASCII                          |
| `,`     | Input character to current cell                       |
| `[`     | Jump past matching `]` if cell is 0                   |
| `]`     | Jump back to matching `[` if cell is non-zero         |

## Running Tests

```bash
# Basic tests with coverage
just test

# Tests with race detection and extra checks
just test-extra

# Full test suite with linting
just test-full

# Open coverage report in browser
just test-open
```

