package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/ratludu/momento/internal/database"
	"github.com/spf13/cobra"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		queries := database.New(db)

		name, _ := cmd.Flags().GetString("add")
		if name == "" {

			profiles, err := queries.GetAllProfiles(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			for i := range profiles {
				fmt.Println(profiles[i].ProfileName)
			}
		} else {
			profile, err := queries.AddProfile(context.Background(), name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Profile Created")
			fmt.Println(profile)
		}
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
	profilesCmd.Flags().StringP("add", "a", "", "Add a new profile with name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
