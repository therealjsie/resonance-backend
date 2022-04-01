/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
	"github.com/therealjsie/resonance-backend/internals"
)

var steps int

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migration all up to the latest version.",
	Long:  `Runs migration all up to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calling migrate")

		config := internals.ReadConfig()

		kodataPath := os.Getenv("KO_DATA_PATH")
		migrationFiles := fmt.Sprintf("file://%s/migrations", kodataPath)

		m, err := migrate.New(
			migrationFiles,
			config.DatabaseConnectionString(),
		)

		if err != nil {
			panic(fmt.Errorf("Fatal error while creating migration object: %w \n", err))
		}

		if steps == 0 {
			fmt.Print("Migrating to the last version. \n")
			err = m.Up()
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Migrating %d steps. \n", steps)
			err = m.Steps(steps)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().IntVarP(&steps, "steps", "s", 0, "number of steps to migrate")
}
