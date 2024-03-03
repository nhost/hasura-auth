package cmd

import (
	"time"

	"github.com/gin-contrib/cors"
)

func getCORSConfig() cors.Config {
	return cors.Config{ //nolint:exhaustruct
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "HEAD", "DELETE"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //nolint: gomnd
	}
}
