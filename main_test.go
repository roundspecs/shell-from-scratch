package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	f, err := os.Open("does-not-exist.txt")
	fmt.Printf("f: %v\n", f)
	fmt.Printf("err: %v\n", err)
}
