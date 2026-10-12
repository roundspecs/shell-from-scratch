package builtin

import (
	"fmt"
	"io"
	"strings"
)

func Echo(args []string, w io.Writer) error {
	_, _ = fmt.Fprintln(w, strings.Join(args, " "))
	return nil
}
