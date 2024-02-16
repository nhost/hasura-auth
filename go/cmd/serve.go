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
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"
	"github.com/urfave/cli/v2"
)

const (
	flagAPIPrefix           = "api-prefix"
	flagPort                = "port"
	flagDebug               = "debug"
	flagLogFormatTEXT       = "log-format-text"
	flagTrustedProxies      = "trusted-proxies"
	flagPostgresConnection  = "postgres"
	flagNodeServerPath      = "node-server-path"
	flagDatabseURL          = "database-url"
	flagDisableSignup       = "disable-signup"
	flagConcealErrors       = "conceal-errors"
	flagDefaultAllowedRoles = "default-allowed-roles"
	flagDefaultRole         = "default-role"
	flagDefaultLocale       = "default-locale"
	flagDisableNewUsers     = "disable-new-users"
	flagGravatarEnabled     = "gravatar-enabled"
	flagGravatarDefault     = "gravatar-default"
	flagGravatarRating      = "gravatar-rating"
)

func CommandServe() *cli.Command { //nolint:funlen
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
				Name:     flagLogFormatTEXT,
				Usage:    "format logs in plain text",
				Category: "general",
				EnvVars:  []string{"LOG_FORMAT_TEXT"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagPostgresConnection,
				Usage:    "Postgres connection string",
				Value:    "postgres://postgres:postgres@localhost:5432/local?sslmode=disable",
				Category: "postgres",
				EnvVars:  []string{"POSTGRES_CONNECTION", "HASURA_GRAPHQL_DATABASE_URL"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagNodeServerPath,
				Usage:    "Path to the node server",
				Value:    ".",
				Category: "node",
				EnvVars:  []string{"NODE_SERVER_PATH"},
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagDisableSignup,
				Usage:    "Disable signup",
				Value:    false,
				Category: "signup",
				EnvVars:  []string{"AUTH_DISABLE_SIGNUP"},
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagConcealErrors,
				Usage:    "Conceal errors",
				Value:    false,
				Category: "server",
				EnvVars:  []string{"AUTH_CONCEAL_ERRORS"},
			},
			&cli.StringSliceFlag{ //nolint: exhaustruct
				Name:     flagDefaultAllowedRoles,
				Usage:    "Default allowed roles",
				Category: "signup",
				Value:    cli.NewStringSlice("user", "me"),
				EnvVars:  []string{"AUTH_USER_DEFAULT_ALLOWED_ROLES"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagDefaultRole,
				Usage:    "Default role",
				Category: "signup",
				Value:    "user",
				EnvVars:  []string{"AUTH_USER_DEFAULT_ROLE"},
			},
			&cli.StringFlag{ //nolint: exhaustruct
				Name:     flagDefaultLocale,
				Usage:    "Default locale",
				Category: "signup",
				Value:    "en",
				EnvVars:  []string{"AUTH_LOCALE_DEFAULT"},
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagDisableNewUsers,
				Usage:    "Disable new users",
				Category: "signup",
				EnvVars:  []string{"AUTH_DISABLE_NEW_USERS"},
			},
			&cli.BoolFlag{ //nolint: exhaustruct
				Name:     flagGravatarEnabled,
				Usage:    "Enable gravatar",
				Category: "gravatar",
				Value:    true,
				EnvVars:  []string{"AUTH_GRAVATAR_ENABLED"},
			},
			&cli.GenericFlag{ //nolint: exhaustruct
				Name: flagGravatarDefault,
				Value: &EnumValue{ //nolint: exhaustruct
					Enum: []string{
						"blank",
						"identicon",
						"monsterid",
						"wavatar",
						"retro",
						"robohash",
						"mp",
						"404",
					},
					Default: "blank",
				},
				Usage:    "Gravatar default",
				Category: "gravatar",
				EnvVars:  []string{"AUTH_GRAVATAR_DEFAULT"},
			},
			&cli.GenericFlag{ //nolint: exhaustruct
				Name: flagGravatarRating,
				Value: &EnumValue{ //nolint: exhaustruct
					Enum: []string{
						"g",
						"pg",
						"r",
						"x",
					},
					Default: "g",
				},
				Usage:    "Gravatar rating",
				Category: "gravatar",
				EnvVars:  []string{"AUTH_GRAVATAR_RATING"},
			},
		},
		Action: serve,
	}
}

func getNodeServer(cCtx *cli.Context) *exec.Cmd {
	env := os.Environ()
	found := false
	authPort := strconv.Itoa(cCtx.Int(flagPort) + 1)
	for i, v := range env {
		if strings.HasPrefix(v, "AUTH_PORT=") {
			found = true
			env[i] = "AUTH_PORT=" + authPort
		}
	}
	if !found {
		env = append(env, "AUTH_PORT="+authPort)
	}
	env = append(env, "NODE_EXTRA_CA_CERTS=/etc/ssl/certs/ca-bundle.crt")
	env = append(env, "PWD="+cCtx.String(flagNodeServerPath))
	env = append(env, "AUTH_VERSION="+cCtx.App.Version)

	cmd := exec.CommandContext(cCtx.Context, "node", "./dist/start.js")
	cmd.Dir = cCtx.String(flagNodeServerPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	return cmd
}

func getGoServer(cCtx *cli.Context, db *sql.Queries, logger *slog.Logger) (*http.Server, error) {
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
	)

	auth := controller.New(
		db,
		controller.Config{
			ConcealErrors:       cCtx.Bool(flagConcealErrors),
			DisableSignup:       cCtx.Bool(flagDisableSignup),
			DisableNewUsers:     cCtx.Bool(flagDisableNewUsers),
			DefaultAllowedRoles: cCtx.StringSlice(flagDefaultAllowedRoles),
			DefaultRole:         cCtx.String(flagDefaultRole),
			DefaultLocale:       cCtx.String(flagDefaultLocale),
			GravatarEnabled:     cCtx.Bool(flagGravatarEnabled),
			GravatarDefault:     cCtx.String(flagGravatarDefault),
			GravatarRating:      cCtx.String(flagGravatarRating),
		},
	)
	handler := api.NewStrictHandler(auth, []api.StrictMiddlewareFunc{})
	mw := api.MiddlewareFunc(ginmiddleware.OapiRequestValidator(doc))
	api.RegisterHandlersWithOptions(
		router,
		handler,
		api.GinServerOptions{
			BaseURL:      cCtx.String(flagAPIPrefix),
			Middlewares:  []api.MiddlewareFunc{mw},
			ErrorHandler: nil,
		},
	)

	nodejsHandler, err := nodejsHandler()
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
	logger := getLogger(cCtx.Bool(flagDebug), cCtx.Bool(flagLogFormatTEXT))
	logger.Info(cCtx.App.Name + " v" + cCtx.App.Version)
	logFlags(logger, cCtx)

	ctx, cancel := context.WithCancel(cCtx.Context)
	defer cancel()

	nodeServer := getNodeServer(cCtx)
	go func() {
		defer cancel()
		if err := nodeServer.Run(); err != nil {
			logger.Error("node server failed", slog.String("error", err.Error()))
		}
	}()

	conn, err := pgx.Connect(ctx, cCtx.String(flagPostgresConnection))
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer conn.Close(ctx)

	server, err := getGoServer(cCtx, sql.New(conn), logger)
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
