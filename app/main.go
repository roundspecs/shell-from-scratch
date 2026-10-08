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
			} else {
				fmt.Println(args[1] + ": not found")
			}
		default:
			fmt.Printf("%v: command not found\n", command)
		}
	}
}
