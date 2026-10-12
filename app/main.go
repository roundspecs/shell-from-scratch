package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
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
		redirectIdx := slices.Index(args, ">")
		var redirect string
		if redirectIdx == -1 {
			redirectIdx = slices.Index(args, "1>")
		}
		if redirectIdx != -1 {
			redirect = args[redirectIdx+1]
			args = args[:redirectIdx]
		}
		w := GetWriter(redirect)

		if len(args) == 0 {
			continue
		} else if builtinFunc, ok := builtin.BuiltinCommands[args[0]]; ok {
			_ = builtinFunc(args[1:], w)
		} else if path := builtin.GetPath(args[0]); path != "" {
			cmd := exec.Command(args[0], args[1:]...)
			b, err := cmd.Output()
			if err != nil {
				fmt.Print(string(err.(*exec.ExitError).Stderr))
			}
			_, _ = fmt.Fprint(w, string(b))
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
	backSlash := false
	arg := strings.Builder{}
	for {
		_, err := prompt.Read(char)
		if err != nil {
			args = append(args, arg.String())
			break
		}

		if backSlash {
			if doubleQuote && !strings.ContainsRune(`"\`, rune(char[0])) {
				arg.WriteByte('\\')
			}
			arg.WriteByte(char[0])
			backSlash = false
			continue
		}
		if !doubleQuote && char[0] == '\'' {
			singleQuote = !singleQuote
			continue
		}
		if !singleQuote && char[0] == '"' {
			doubleQuote = !doubleQuote
			continue
		}
		if singleQuote {
			arg.WriteByte(char[0])
			continue
		}
		if doubleQuote {
			if char[0] == '\\' {
				backSlash = true
				continue
			}
			arg.WriteByte(char[0])
			continue
		}
		if char[0] == '\\' {
			backSlash = true
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

func GetWriter(redirect string) io.Writer {
	if redirect == "" {
		return os.Stdin
	}
	idx := strings.LastIndexByte(redirect, '/')
	dir := redirect[:idx]
	if dir != "" {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.Create(redirect)
	if err != nil {
		fmt.Println(err)
		return os.Stdin
	}
	return f
}
