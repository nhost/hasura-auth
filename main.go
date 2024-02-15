package main

import (
	"log"
	"os"

	"github.com/nhost/hasura-auth/go/cmd"
	"github.com/urfave/cli/v2"
)

var Version string

//go:generate oapi-codegen -config go/api/server.cfg.yaml go/api/openapi.yaml
//go:generate oapi-codegen -config go/api/types.cfg.yaml go/api/openapi.yaml
func main() {
	serveCmd := cmd.CommandServe()
	app := &cli.App{ //nolint:exhaustruct
		Name:    "auth",
		Version: Version,
		Usage:   "Nhost Auth API server",
		Commands: []*cli.Command{
			serveCmd,
		},
		Action: serveCmd.Action,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
