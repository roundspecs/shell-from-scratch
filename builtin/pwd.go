package builtin

import (
	"fmt"
	"io"
	"os"
)

func Pwd(args []string, stdout, stderr io.Writer) error {
	wd, _ := os.Getwd()
	_, _ = fmt.Fprintln(stdout, wd)
	return nil
}
