package builtin

import (
	"fmt"
	"os"
	"strings"
)

func Cd(args []string) error {
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
