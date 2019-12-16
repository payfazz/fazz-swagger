package command

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

type fazzSwagger struct {
	basename string
}

// NewFazzSwagger create base command as fazzswagger
func NewFazzSwagger() *fazzSwagger {
	return &fazzSwagger{
		basename: "fazzswagger",
	}
}

// Execute is a function that executed first time upon calling base cmd
func (fs *fazzSwagger) Execute() {
	rootCmd := &cobra.Command{
		Use:   NewFazzSwagger().basename,
		Short: "FazzSwagger is a compiler for OpenAPI3 Files",
		Long:  `FAZZSWAGGER: A Compiler to merge every component files of OpenAPI3 to one giant Swagger Yaml File that is ready to be served.`,
	}

	rootCmd.AddCommand(NewCompileCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
