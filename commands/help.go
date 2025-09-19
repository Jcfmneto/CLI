package commands

import "fmt"

func help() {
	commands := GetCommands()

	fmt.Println("All Commands Available:")
	for command, val := range commands {
		fmt.Printf("%s: %s \n", command, val.Description)
	}
}
