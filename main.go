package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(s string) error {
	trimmedInput := strings.TrimSuffix(s, "\n")
	args := strings.Split(trimmedInput, " ")

	switch args[0] {
		case "cd" :
			if len(args) < 2 {
				return errors.New("Path Required")
			}
			return os.Chdir(args[1])
		case "exit":
			os.Exit(0)
	}

	command := exec.Command(args[0], args[1:]...)
	command.Stderr =  os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}

func getInputIndicator() string {
	uname := exec.Command("uname")
	unameOutput, err := uname.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	hostname := exec.Command("hostname")
	hostnameOutput, err := hostname.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	pwd := exec.Command("pwd")
	pwdOutput, err := pwd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	// <uname> + <hostname> in <pwd>
	indicator := string(unameOutput) + " at " + string(hostnameOutput) + " in " + string(pwdOutput)
	trimmedIndicator := strings.ReplaceAll(indicator, "\n", "")
	return trimmedIndicator
}

func main() {	
	reader := bufio.NewReader(os.Stdin)
	// Create input loop
	for {
		fmt.Print(getInputIndicator())
		fmt.Print(" > ")
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		execInput(line)
	}
}
