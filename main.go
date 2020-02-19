package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	file, err := os.OpenFile("tmp/a.s", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	code := `
.intel_syntax noprefix
.global main

main:
    mov rax, 42
    ret
`

	file.WriteString(code)

	cmd := exec.Command("gcc", "tmp/a.s")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		exitError := err.(*exec.ExitError)
		fmt.Fprintln(os.Stderr, "Assembler returns exit code", exitError.ExitCode())
		os.Exit(1)
	}
}
