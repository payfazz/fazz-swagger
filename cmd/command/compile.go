package command

import (
	"github.com/payfazz/fazz-swagger/internal/compile"
	"github.com/spf13/cobra"
)

type compileCommand struct{}

// NewCompile create compile as sub command
func NewCompileCommand() *cobra.Command {
	c := compileCommand{}
	cmd := &cobra.Command{
		Use:   "compile [directory]",
		Short: "Compile files to giant Swagger File",
		Args:  cobra.ExactArgs(1),
		Run:   c.Run,
	}

	return cmd
}

// Run the executes command logic
func (c *compileCommand) Run(cmd *cobra.Command, args []string) {
	compile.Compile(args[0])
}
