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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := GetDbPath()

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		queries := database.New(db)

		endTime := ConvertTime(time.Now())
		session, err := queries.CloseSession(context.Background(), endTime)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Stopped Session:", session.Note)
		fmt.Println("	- Start:", session.Start)
		fmt.Println("	- End:", session.End)

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
