package commands

type CliCommand struct {
	Command     string
	Description string
	Callback    func()
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Command:     "help",
			Description: "Print all commands",
			Callback:    help,
		},
		"exit": {
			Command:     "exit",
			Description: "Kills the process",
			Callback:    exit,
		},
	}
}
