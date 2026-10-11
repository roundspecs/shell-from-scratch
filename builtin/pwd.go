package builtin

import (
	"fmt"
	"os"
)

func Pwd(args []string) error {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	return nil
}
