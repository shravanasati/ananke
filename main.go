package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/shravanasati/ananke/html2md"
)

const helpText = `
ananke is a simple command line tool to convert html to markdown. it can read input from stdin as well as from the given arguments.

visit "https://github.com/shravanasati/ananke" for more information.
`

func main() {
	converter := html2md.NewConverter()

	// Check if there is any input available in stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read all input from stdin at once
		input, err := io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			fmt.Println("error reading input: ", err)
			os.Exit(1)
		}

		output, err := converter.ConvertString(string(input))
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		fmt.Println(output)
	} else {
		// Handle input from arguments
		if len(os.Args) > 1 {
			args := os.Args[1:]
			text := strings.Join(args, " ")
			output, err := converter.ConvertString(text)
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}
			fmt.Println(output)
		} else {
			// Print help text if no arguments are provided
			fmt.Print(helpText)
		}
	}
}
