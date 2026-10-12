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
		args, stdout, stderr := extractIO(args)

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

func extractIO(args []string) (argsNew []string, stdout io.Writer, stderr io.Writer) {
	stdout, stderr = os.Stdout, os.Stderr
	var path string

	stdoutRedirectIdx := slices.Index(args, ">")
	if stdoutRedirectIdx == -1 {
		stdoutRedirectIdx = slices.Index(args, "1>")
	}
	if stdoutRedirectIdx != -1 {
		path = args[stdoutRedirectIdx+1]
		argsNew = args[:stdoutRedirectIdx]
		stdout = GetWriter(path, true)
		return
	}

	stderrRedirectIdx := slices.Index(args, "2>")
	if stderrRedirectIdx != -1 {
		path = args[stderrRedirectIdx+1]
		argsNew = args[:stderrRedirectIdx]
		stderr = GetWriter(path, true)
		return
	}

	stdoutAppendIdx := slices.Index(args, ">>")
	if stdoutAppendIdx == -1 {
		stdoutAppendIdx = slices.Index(args, "1>>")
	}
	if stdoutAppendIdx != -1 {
		path = args[stdoutAppendIdx+1]
		argsNew = args[:stdoutAppendIdx]
		stdout = GetWriter(path, false)
		return
	}

	stderrAppendIdx := slices.Index(args, "2>>")
	if stderrAppendIdx != -1 {
		path = args[stderrAppendIdx+1]
		argsNew = args[:stderrAppendIdx]
		stderr = GetWriter(path, false)
		return
	}
	return args, stdout, stderr
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

func GetWriter(path string, truncate bool) io.Writer {
	idx := strings.LastIndexByte(path, '/')
	if idx != -1 {
		dir := path[:idx]
		if dir != "" {
			_ = os.MkdirAll(dir, os.ModePerm)
		}
	}
	var flag int
	if truncate {
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	} else {
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	}
	f, _ := os.OpenFile(path, flag, 0o644)
	return f
}
