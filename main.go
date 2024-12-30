package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shravanasati/ananke/html2md"
)

const helpText = `
ananke is a simple command line tool to convert html to markdown. it can read input from stdin as well as from the given arguments.

visit "https://github.com/shravanasati/ananke".. for more information.
`

func main() {
	converter:= html2md.NewConverter()

	// Check if there is any input available in stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			input := scanner.Text()
			output, err := (converter.ConvertString(input))
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}
			fmt.Println(output)
		}

		if err := scanner.Err(); err != nil {
			return
		}
	} else {
		if len(os.Args) > 1 {
			args := os.Args[1:]
			text := strings.Join(args, " ")
			output, err := (converter.ConvertString(text))
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}
			fmt.Println(output)
		} else {
			fmt.Print(helpText)
		}
	}
}