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
		command = strings.TrimSpace(command[:len(command)-1])

		args := strings.Split(command, " ")
		switch args[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "type":
			if slices.Contains(builtinCommands, args[1]) {
				fmt.Println(args[1], "is a shell builtin")
			} else if path := isInPath(args[1]); path != "" {
				fmt.Println(args[1], "is", path)
			} else {
				fmt.Println(args[1] + ": not found")
			}
		default:
			fmt.Printf("%v: command not found\n", command)
		}
	}
}

func isInPath(command string) string {
	paths := strings.SplitSeq(os.Getenv("PATH"), ":")
	for path := range paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error reading directories in ", path+":", err.Error())
		}
		for _, entry := range entries {
			executable, err := isExecutable(entry)
			if err != nil {
				fmt.Println("Error reading info about ", entry.Name()+":", err.Error())
			}
			if !executable {
				continue
			}
			if entry.Name() == command {
				return path + "/"
			}
		}
	}
	return ""
}

func isExecutable(entry os.DirEntry) (bool, error) {
	info, err := entry.Info()
	if err != nil {
		return false, err
	}
	mode := info.Mode()
	return mode.IsRegular() && mode&0111 != 0, nil
}
