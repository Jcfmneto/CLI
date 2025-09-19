package main

import (
	"bufio"
	"fmt"
	"github.com/Jcfmneto/CLI/commands"
	"os"
	"strings"
)

func startRepl() {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		commandsAvailable := commands.GetCommands()

		if cmd, ok := commandsAvailable[command]; ok {
			cmd.Callback()
		} else {
			fmt.Printf("'%s' not found \n", command)
		}

	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
