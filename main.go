package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
	"errors"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = executeInput(input); err != nil {
			fmt.Println(os.Stderr, err)
		}
	}
}

func executeInput(input string) error {
	// remove the new line character at the end of the input 
	input = strings.TrimSuffix(input, "\n")
	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0]{
	case "cd":
		// path cannot be empty when using 'cd' as for now
		if len(args) < 2{
			return errors.New("path is required for execution")
		}
		return os.Chdir(args[1])
		
	case "exit":
		os.Exit(0)
	}

	// prepare the input command
	cmd := exec.Command(args[0], args[1:]...)
	// assign error device
	cmd.Stderr = os.Stderr
	// assign output device
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

