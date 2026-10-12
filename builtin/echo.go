package builtin

import (
	"fmt"
	"io"
	"strings"
)

func Echo(args []string, stdout, stderr io.Writer) error {
	_, _ = fmt.Fprintln(stdout, strings.Join(args, " "))
	return nil
}
