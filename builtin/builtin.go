// Package builtin
package builtin

import "io"

var BuiltinCommands map[string]CommandFunc

type CommandFunc func(args []string, w io.Writer) error

func init() {
	BuiltinCommands = map[string]CommandFunc{
		"cd":   Cd,
		"echo": Echo,
		"exit": Exit,
		"pwd":  Pwd,
		"type": Type,
	}
}
