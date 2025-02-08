package main

import (
	"os"
	"thorston-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
