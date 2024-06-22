package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kishor82/lexer/src/lexer"
	"github.com/kishor82/lexer/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	// filePath relative to the current working directory of the program's execution
	examplesDir := "./examples"

	err := filepath.Walk(examplesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Printf("------------------------ %s -----------------------------\n", info.Name())

			bytes, _ := os.ReadFile(path)
			tokens := lexer.Tokenize(string(bytes))

			fmt.Printf(":- Tokens -:\n")
			fmt.Printf("------------\n")

			for _, token := range tokens {
				token.Debug()
			}

			fmt.Printf(":- AST -:\n")
			fmt.Printf("---------\n")

			if info.Name() == "02.lang" || info.Name() == "03.lang" {
				// TODO: parser doesn't handle minus sign ("-") or DASH so conditionally parsing here
				ast := parser.Parse(tokens)
				litter.Dump(ast)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", examplesDir, err)
	}
}
