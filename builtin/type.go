package builtin

import "fmt"

func Type(args []string) error {
	if _, ok := BuiltinCommands[args[0]]; ok {
		fmt.Println(args[0], "is a shell builtin")
	} else if path := GetPath(args[0]); path != "" {
		fmt.Println(args[0], "is", path)
	} else {
		fmt.Println(args[0] + ": not found")
	}
	return nil
}
