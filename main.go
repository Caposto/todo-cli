package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(s string) error {
	trimmedInput := strings.TrimSuffix(s, "\n")
	command := exec.Command(trimmedInput)
	command.Stderr =  os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}

func main() {	
	reader := bufio.NewReader(os.Stdin)
	// Create input loop
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		execInput(line)
	}
}
