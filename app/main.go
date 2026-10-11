package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/builtin"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		args := getArgs(strings.NewReader(command))

		if len(args) == 0 {
			continue
		} else if builtinFunc, ok := builtin.BuiltinCommands[args[0]]; ok {
			_ = builtinFunc(args[1:])
		} else if path := builtin.GetPath(args[0]); path != "" {
			cmd := exec.Command(args[0], args[1:]...)
			b, _ := cmd.Output()
			fmt.Print(string(b))
		} else {
			fmt.Printf("%v: command not found\n", command)
		}
	}
}

func getArgs(prompt io.Reader) []string {
	args := []string{}
	char := make([]byte, 1)
	singleQuote := false
	doubleQuote := false
	arg := strings.Builder{}
	for {
		_, err := prompt.Read(char)
		if err != nil {
			args = append(args, arg.String())
			break
		}

		if !doubleQuote && char[0] == '\'' {
			singleQuote = !singleQuote
			continue
		}
		if !singleQuote && char[0] == '"' {
			doubleQuote = !doubleQuote
			continue
		}
		if singleQuote || doubleQuote {
			arg.WriteByte(char[0])
			continue
		}

		if char[0] == ' ' {
			if arg.Len() > 0 {
				args = append(args, arg.String())
				arg.Reset()
				continue
			}
		} else {
			arg.WriteByte(char[0])
		}
	}
	return args
}
