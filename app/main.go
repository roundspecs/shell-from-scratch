package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("$ ")
	command, err := scanner.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Printf("%v: command not found", command[:len(command)-1])
}
