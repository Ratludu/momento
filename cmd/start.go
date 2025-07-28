/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a session in your current profile.",
	Long:  `Starts a session in your current profile. It accepts a optional flag --tag or -t, tagging a specifc task to the profile.`,
	Run: func(cmd *cobra.Command, args []string) {

		dbPath := GetDbPath()

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		queries := database.New(db)

		session, err := queries.GetSessions(context.Background())
		if len(session) != 0 {
			fmt.Printf("Session: %s is still open! Please close it before starting a new one :)\n", session[0].Note)
			return
		}

		profile, err := queries.GetCurrentProfile(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		startTime := ConvertTime(time.Now())
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			log.Fatal(err)
		}

		if tag == "" {
			tag = "misc"
		}

		result, err := queries.CreateSession(context.Background(), database.CreateSessionParams{
			ProfileID: profile.ID,
			Start:     startTime,
			Note:      tag,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Starting", tag, "...")
		fmt.Println("	- Profile:", profile.ProfileName)
		fmt.Println("	- Start time:", result.Start)
		fmt.Println("Lock In!")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("tag", "t", "", "Add a tag to track e.g. course name. Default is misc")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
