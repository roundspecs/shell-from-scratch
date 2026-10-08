# Shell

A lightweight, custom command-line shell written from scratch in Go. Built to explore operating systems concepts, REPL mechanics, command parsing, and process execution.

## Overview

An interactive Unix-like shell environment implemented in Go. It provides a standard Read-Eval-Print Loop (REPL), parses user input, and evaluates built-in shell commands and program execution flows without relying on existing shell wrappers.

This project was built to gain a deep, hands-on understanding of how modern shells (such as Bash and Zsh) interact with operating system primitives, handle I/O, and manage command lifecycles.

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (1.22 or higher recommended)
- A POSIX-compatible environment (Linux, macOS) or WSL on Windows

### Installation

Clone the repository:

```bash
git clone https://github.com/roundspecs/shell-from-scratch.git
cd codecrafters-shell-go
```

### Running the Shell

You can run the shell directly using `go run`:

```bash
go run app/main.go
```

Or build a standalone binary:

```bash
# Build binary
go build -o bin/shell app/main.go

# Run
./bin/shell
```

---

## Usage Example

```text
$ echo Hello, world!
Hello, world!

$ type echo
echo is a shell builtin

$ type exit
exit is a shell builtin

$ type my_script
my_script: not found

$ nonexistent-command
nonexistent-command: command not found

$ exit
```

