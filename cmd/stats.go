package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/ratludu/momento/internal/database"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Provides simple statistics for your profiles",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := GetDbPath()

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		queries := database.New(db)

		data, err := queries.GetSessionsWithProfile(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		profiles, err := queries.GetAllProfiles(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var mins time.Duration
		for p := range profiles {
			mins = 0.0
			for d := range data {

				if profiles[p].ProfileName == data[d].ProfileName.String {
					start, _ := ParseTime(data[d].SessionStart)
					end, _ := ParseTime(data[d].SessionEnd)
					mins += end.Sub(start)
				}

			}
			fmt.Printf("	- Profile: %s Total: %.2f minutes\n", profiles[p].ProfileName, mins.Minutes())

		}

	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
