package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/urfave/cli/v2"
)

const (
	flagAPIPrefix          = "api-prefix"
	flagPort               = "port"
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
				Name:     flagAPIPrefix,
				Usage:    "prefix for all routes",
				Value:    "/v1",
				Category: "server",
				EnvVars:  []string{"AUTH_API_PREFIX"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagPort,
				Usage:    "Port to bind to",
				Value:    "4000",
				Category: "server",
				EnvVars:  []string{"AUTH_PORT"},
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

func getNodeServer(ctx context.Context, nodeServerPath string, port int) *exec.Cmd {
	env := os.Environ()
	found := false
	for i, v := range env {
		if strings.HasPrefix(v, "AUTH_PORT=") {
			found = true
			env[i] = "AUTH_PORT=" + strconv.Itoa(port+1)
		}
	}
	if !found {
		env = append(env, "AUTH_PORT="+strconv.Itoa(port+1))
	}

	cmd := exec.CommandContext(ctx, "pnpm", "start")
	cmd.Dir = nodeServerPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	return cmd
}

func getGoServer(cCtx *cli.Context, logger *slog.Logger) (*http.Server, error) {
	router := gin.New()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(api.OpenAPISchema)
	if err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI schema: %w", err)
	}
	doc.AddServer(&openapi3.Server{ //nolint:exhaustruct
		URL: cCtx.String(flagAPIPrefix),
	})

	router.Use(
		// ginmiddleware.OapiRequestValidator(doc),
		gin.Recovery(),
		middleware.Logger(logger),
		cors.New(cors.Config{ //nolint: exhaustruct
			AllowOrigins:     cCtx.StringSlice(flagAllowCORSOrigin),
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
		}),
	)

	// auth := &controller.Auth{}
	// handler := api.NewStrictHandler(auth, nil)
	// mw := api.MiddlewareFunc(ginmiddleware.OapiRequestValidator(doc)),

	// api.RegisterHandlersWithOptions(
	// 	router,
	// 	handler,
	// 	api.GinServerOptions{
	// 		BaseURL: cCtx.String(flagAPIPrefix),
	// 		Middlewares: []api.MiddlewareFunc{mw},
	// 		ErrorHandler: nil,
	// 	},
	// )

	nodejsHandler, err := nodejsHandler(cCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to create nodejs handler: %w", err)
	}
	router.NoRoute(nodejsHandler)

	server := &http.Server{ //nolint:exhaustruct
		Addr:              ":" + cCtx.String(flagPort),
		Handler:           router,
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

	nodeServer := getNodeServer(ctx, cCtx.String(flagNodeServerPath), cCtx.Int(flagPort))
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
