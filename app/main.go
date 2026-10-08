package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		command = strings.TrimSpace(command[:len(command)-1])
		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%v: command not found\n", command)
		}
	}
}
