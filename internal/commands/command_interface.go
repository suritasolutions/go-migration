package commands

type CommandInterface interface {
	Run(cmd *CommandInterface, args []string)
}
