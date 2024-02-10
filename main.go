package main

import (
	"log"
	"os"

	"github.com/nhost/hasura-auth/go/cmd"
	"github.com/urfave/cli/v2"
)

var Version string

//go:generate oapi-codegen -config api/server.cfg.yaml api/openapi.yaml
//go:generate oapi-codegen -config api/types.cfg.yaml api/openapi.yaml
func main() {
	app := &cli.App{ //nolint:exhaustruct
		Name:    "auth",
		Version: Version,
		Usage:   "Nhost Auth API server",
		Commands: []*cli.Command{
			cmd.CommandServe(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
