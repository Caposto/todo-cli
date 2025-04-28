package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	fmt.Println("LS COMMAND")
	ls := exec.Command("ls", ".")

	output, err := ls.CombinedOutput()

	if  err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
