package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
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
	// split the input to command and arguments
	args := strings.Split(input, " ")
	// prepare the input command
	cmd := exec.Command(args[0], args[1:]...)
	// assign error device
	cmd.Stderr = os.Stderr
	// assign output device
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

