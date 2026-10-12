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

		stdoutRedirectIdx := slices.Index(args, ">")
		var stdoutRedirect string
		if stdoutRedirectIdx == -1 {
			stdoutRedirectIdx = slices.Index(args, "1>")
		}
		if stdoutRedirectIdx != -1 {
			stdoutRedirect = args[stdoutRedirectIdx+1]
			args = args[:stdoutRedirectIdx]
		}
		stdout := GetWriter(stdoutRedirect)

		stderrRedirectIdx := slices.Index(args, "2>")
		var stderrRedirect string
		if stderrRedirectIdx != -1 {
			stderrRedirect = args[stderrRedirectIdx+1]
			args = args[:stderrRedirectIdx]
		}
		stderr := GetWriter(stderrRedirect)

		if len(args) == 0 {
			continue
		} else if builtinFunc, ok := builtin.BuiltinCommands[args[0]]; ok {
			_ = builtinFunc(args[1:], stdout, stderr)
		} else if path := builtin.GetPath(args[0]); path != "" {
			cmd := exec.Command(args[0], args[1:]...)
			b, err := cmd.Output()
			if err != nil {
				_, _ = fmt.Fprint(stderr, string(err.(*exec.ExitError).Stderr))
			}
			_, _ = fmt.Fprint(stdout, string(b))
		} else {
			_, _ = fmt.Fprintf(stderr, "%v: command not found\n", command)
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
