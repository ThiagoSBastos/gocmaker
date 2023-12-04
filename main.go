package main

import (
	"encoding/json"
	"log"
	"os"
)

// Based on https://clang.llvm.org/docs/JSONCompilationDatabase.html
type CompileCommandsJSON struct {
	Directory string
	Arguments []string `json:"arguments,omitempty"`
	Command   string   `json:"command,omitempty"`
	File      string
}

func main() {
	exampleData, err := os.ReadFile("compile_commands.json")
	if err != nil {
		log.Fatal(err)
	}

	var compile_commands CompileCommandsJSON
	err = json.Unmarshal(exampleData, &compile_commands)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"Struct is:\n Directory: %v\n Arguments: %v\n Command: %v\n File: %v\n",
		compile_commands.Directory,
		compile_commands.Arguments,
		compile_commands.Command,
		compile_commands.File)
}
