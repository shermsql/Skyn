package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Help() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Skyn, Failed To Determine Own File Location!")
		return
	}

	dir := filepath.Dir(filename)

	helpPath := filepath.Join(dir, "Help.txt")

	content, err := os.ReadFile(helpPath)
	if err != nil {
		fmt.Printf("Skyn, Help File Could Not Be Read!\nPath, %s\nError, %v\n", helpPath, err)
		return
	}

	fmt.Println(string(content))
}
