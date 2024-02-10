package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/middleware"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"
	"github.com/urfave/cli/v2"
)

const (
	flagPathPrefix         = "path-prefix"
	flagBind               = "bind"
	flagDebug              = "debug"
	flagLogFormatJSON      = "log-format-json"
	flagTrustedProxies     = "trusted-proxies"
	flagAllowCORSOrigin    = "allow-cors-origin"
	flagPostgresConnection = "postgres"
	flagNodeServerPath     = "node-server-path"
)

func CommandServe() *cli.Command {
	return &cli.Command{ //nolint: exhaustruct
		Name:  "serve",
		Usage: "Serve the application",
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagPathPrefix,
				Usage:    "prefix for all routes",
				Value:    "/v1",
				Category: "server",
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagBind,
				Usage:    "bind address",
				Value:    ":8091",
				Category: "server",
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagDebug,
				Usage:    "enable debug logging",
				Category: "general",
				EnvVars:  []string{"DEBUG"},
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagLogFormatJSON,
				Usage:    "format logs in JSON",
				Category: "general",
			},
			&cli.StringSliceFlag{ //nolint: exhaustruct
				Name:     flagAllowCORSOrigin,
				Usage:    "Allow CORS from these origins",
				Value:    cli.NewStringSlice("*"),
				Category: "server",
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagPostgresConnection,
				Usage:    "Postgres connection string",
				Value:    "postgres://postgres:postgres@localhost:5432/local?sslmode=disable",
				Category: "postgres",
				EnvVars:  []string{"POSTGRES_CONNECTION"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagNodeServerPath,
				Usage:    "Path to the node server",
				Value:    ".",
				Category: "node",
				EnvVars:  []string{"NODE_SERVER_PATH"},
			},
		},
		Action: serve,
	}
}

func getNodeServer(ctx context.Context, nodeServerPath string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "pnpm", "start")
	cmd.Dir = nodeServerPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func getGoServer(cCtx *cli.Context, logger *slog.Logger) (*http.Server, error) {
	r := gin.New()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(api.OpenAPISchema)
	if err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI schema: %w", err)
	}
	doc.AddServer(&openapi3.Server{ //nolint:exhaustruct
		URL: cCtx.String(flagPathPrefix),
	})

	auth := &controller.Auth{}

	r.Use(
		ginmiddleware.OapiRequestValidator(doc),
		gin.Recovery(),
		middleware.Logger(logger),
		cors.New(cors.Config{ //nolint: exhaustruct
			AllowOrigins:     cCtx.StringSlice(flagAllowCORSOrigin),
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
		}),
	)

	handler := api.NewStrictHandler(auth, nil)

	api.RegisterHandlersWithOptions(
		r,
		handler,
		api.GinServerOptions{
			BaseURL:      cCtx.String(flagPathPrefix),
			Middlewares:  []api.MiddlewareFunc{},
			ErrorHandler: nil,
		},
	)

	server := &http.Server{ //nolint:exhaustruct
		Addr:              cCtx.String(flagBind),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second, //nolint:gomnd
	}

	return server, nil
}

func serve(cCtx *cli.Context) error {
	logger := getLogger(cCtx.Bool(flagDebug), cCtx.Bool(flagLogFormatJSON))
	logger.Info(cCtx.App.Name + " v" + cCtx.App.Version)
	logFlags(logger, cCtx)

	ctx, cancel := context.WithCancel(cCtx.Context)
	defer cancel()

	nodeServer := getNodeServer(ctx, cCtx.String(flagNodeServerPath))
	go func() {
		defer cancel()
		if err := nodeServer.Run(); err != nil {
			logger.Error("node server failed", slog.String("error", err.Error()))
		}
	}()

	server, err := getGoServer(cCtx, logger)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	go func() {
		defer cancel()
		if err := server.ListenAndServe(); err != nil {
			logger.Error("server failed", slog.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	logger.Info("shutting down server")
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	return nil
}
