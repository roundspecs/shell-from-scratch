package builtin

import (
	"fmt"
	"io"
	"os"
)

func Pwd(args []string, w io.Writer) error {
	wd, _ := os.Getwd()
	_, _ = fmt.Fprintln(w, wd)
	return nil
}
