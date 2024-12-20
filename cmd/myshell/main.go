package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for i:=0; ;i++ {
		fmt.Fprint(os.Stdout, "$ ")
		
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		if command == "exit 0\n" {
			os.Exit(0)
		}
		
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
