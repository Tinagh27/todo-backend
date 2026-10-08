package main

import (
	"fmt"
	"todo-backend/internal/adapter/inbound/http"

	"todo-backend/internal/config"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		router := http.NewRouter()

		fmt.Printf("Starting HTTP server on :%s\n", cfg.AppPort)

		return router.Run(":" + cfg.AppPort)
	},
}
