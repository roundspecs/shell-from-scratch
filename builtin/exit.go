package builtin

import (
	"io"
	"os"
)

func Exit(args []string, w io.Writer) error {
	os.Exit(0)
	return nil
}
