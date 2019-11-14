package command

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

type fazzSwagger struct {
	basename string
}

func NewFazzSwagger() *fazzSwagger {
	return &fazzSwagger{
		basename: "fazzswagger",
	}
}

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
