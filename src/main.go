package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kishor82/lexer/src/lexer"
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

			for _, token := range tokens {
				token.Debug()
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", examplesDir, err)
	}
}
