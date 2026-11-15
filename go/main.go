package main

import (
	"fmt"
	"os"

	"dbms/engine"
)

type args struct {
	file  string
	query string
}

func parseArgs(argv []string) (args, bool) {
	var a args
	for i := 1; i < len(argv); i++ {
		switch {
		case argv[i] == "--file" && i+1 < len(argv):
			i++
			a.file = argv[i]
		case argv[i] == "--query" && i+1 < len(argv):
			i++
			a.query = argv[i]
		default:
			return a, false
		}
	}
	return a, a.file != "" && a.query != ""
}

func main() {
	a, ok := parseArgs(os.Args)
	if !ok {
		fmt.Fprintf(os.Stderr, "Usage: %s --file <path> --query '<command>'\n", os.Args[0])
		os.Exit(1)
	}

	var storage engine.Storage
	if err := storage.Load(a.file); err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot load %s: %v\n", a.file, err)
		os.Exit(1)
	}

	tokens, err := engine.Tokenize(a.query)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	result, changed, err := engine.Execute(&storage, tokens)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if changed {
		if err := storage.Save(a.file); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	fmt.Println("-> " + result)
}
