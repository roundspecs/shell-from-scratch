package builtin

import (
	"os"
	"strings"
)

func GetPath(command string) string {
	paths := strings.SplitSeq(os.Getenv("PATH"), ":")
	for path := range paths {
		filePath := path + "/" + command
		info, err := os.Stat(filePath)
		if err != nil {
			continue
		}
		mode := info.Mode()
		if mode.IsRegular() && mode&0o111 != 0 {
			return filePath
		}
	}
	return ""
}
