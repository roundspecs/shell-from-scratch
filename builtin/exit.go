package builtin

import (
	"io"
	"os"
)

func Exit(args []string, stdout, stderr io.Writer) error {
	os.Exit(0)
	return nil
}
