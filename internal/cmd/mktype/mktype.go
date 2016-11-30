package main

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/kbinani/win/internal"
)

func main() {
	typename := os.Args[1]
	filename := TemplateFilePath(typename, "")
	if !IsFileExist(filename) {
		f, _ := os.Create(filename)
		if f != nil {
			fmt.Fprintf(f, "%s\n", typename)
			f.Close()
		}
	}
	fmt.Printf("%s\n", filename)
	cmd := exec.Command("vim", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
