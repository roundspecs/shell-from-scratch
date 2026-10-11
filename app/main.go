package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandFunc func(args []string) error

var builtinCommands map[string]CommandFunc

func init() {
	builtinCommands = map[string]CommandFunc{
		"exit": builtinExit,
		"echo": builtinEcho,
		"type": builtinType,
		"pwd":  builtinPwd,
		"cd":   builtinCd,
	}
}

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
		args := strings.Fields(command)

		if len(args) == 0 {
			continue
		} else if builtinFunc, ok := builtinCommands[args[0]]; ok {
			_ = builtinFunc(args[1:])
		} else if path := getPath(args[0]); path != "" {
			cmd := exec.Command(args[0], args[1:]...)
			b, _ := cmd.Output()
			fmt.Print(string(b))
		} else {
			fmt.Printf("%v: command not found\n", command)
		}
	}
}

func builtinExit(args []string) error {
	os.Exit(0)
	return nil
}

func builtinEcho(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func builtinType(args []string) error {
	if _, ok := builtinCommands[args[0]]; ok {
		fmt.Println(args[0], "is a shell builtin")
	} else if path := getPath(args[0]); path != "" {
		fmt.Println(args[0], "is", path)
	} else {
		fmt.Println(args[0] + ": not found")
	}
	return nil
}

func builtinPwd(args []string) error {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	return nil
}

func builtinCd(args []string) error {
	if len(args) == 0 {
		args = append(args, os.Getenv("HOME"))
	} else if strings.HasPrefix(args[0], "~") {
		args[0] = os.Getenv("HOME") + args[0][1:]
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println("cd:", args[0]+": No such file or directory")
	}
	return nil
}

func getPath(command string) string {
	paths := strings.SplitSeq(os.Getenv("PATH"), ":")
	for path := range paths {
		filePath := path + "/" + command
		info, err := os.Stat(filePath)
		if err != nil {
			continue
		}
		mode := info.Mode()
		if mode.IsRegular() && mode&0o111 != 0 {
			return filePath
		}
	}
	return ""
}
