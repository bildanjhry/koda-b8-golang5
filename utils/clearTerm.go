package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearTerm(status int, message string) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	if status == 1 {
		fmt.Printf("%s\n\n", message)
	}
}
