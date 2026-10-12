package builtin

import (
	"fmt"
	"io"
)

func Type(args []string, stdout, stderr io.Writer) error {
	if _, ok := BuiltinCommands[args[0]]; ok {
		_, _ = fmt.Fprintln(stdout, args[0], "is a shell builtin")
	} else if path := GetPath(args[0]); path != "" {
		_, _ = fmt.Fprintln(stdout, args[0], "is", path)
	} else {
		_, _ = fmt.Fprintln(stderr, args[0]+": not found")
	}
	return nil
}
