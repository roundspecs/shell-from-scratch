// Package builtin
package builtin

var BuiltinCommands map[string]CommandFunc

type CommandFunc func(args []string) error

func init() {
	BuiltinCommands = map[string]CommandFunc{
		"exit": Exit,
		"echo": Echo,
		"type": Type,
		"pwd":  Pwd,
		"cd":   Cd,
	}
}
