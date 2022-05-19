package main

import (
	"log"

	"github.com/alice890308/blog-server/cmd/api"
	"github.com/alice890308/blog-server/cmd/file"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "blog-server [module]",
		Short: "Blog Server module entrypoints",
	}

	cmd.AddCommand(api.NewAPICommand())
	cmd.AddCommand(file.NewFileCommand())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
