package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"exit", "echo", "type"}

	for {
		fmt.Print("$ ")
		command, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		args := strings.Fields(command)
		switch args[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "type":
			if slices.Contains(builtinCommands, args[1]) {
				fmt.Println(args[1], "is a shell builtin")
			} else if path := getPath(args[1]); path != "" {
				fmt.Println(args[1], "is", path)
			} else {
				fmt.Println(args[1] + ": not found")
			}
		default:
			fmt.Printf("%v: command not found\n", command)
		}
	}
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
		if mode.IsRegular() && mode&0111 != 0 {
			return filePath
		}
	}
	return ""
}

