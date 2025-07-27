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
	Short: "Show all profiles that have been registered.",
	Long: `Default is to show all profiles that have been regisered for a user. 
For example:
mnt profiles
- work
- study (current)
- personal project
	`,
	Run: func(cmd *cobra.Command, args []string) {

		dbPath := GetDbPath()

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		queries := database.New(db)

		name, _ := cmd.Flags().GetString("add")
		set, _ := cmd.Flags().GetString("set")

		if name != "" {
			profile, err := queries.AddProfile(context.Background(), name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Profile Created")
			fmt.Println(profile)
		}

		if set != "" {
			err := queries.ResetCurrentProfile(context.Background())
			if err != nil {
				log.Fatal("Cannot reset current profile.")
			}

			current, err := queries.SetCurrentProfile(context.Background(), set)
			if err != nil {
				log.Fatal("Could not find profile, please make sure to register it first")
			}

			fmt.Printf("INFO: %s has been set as your current profile.\n", current.ProfileName)
		}

		profiles, err := queries.GetAllProfiles(context.Background())
		if err != nil {
			log.Fatal(err, dbPath)
		}

		fmt.Println("Profiles:")
		for i := range profiles {

			if profiles[i].CurrentProfile == 1 {
				fmt.Println("	-", profiles[i].ProfileName, "(current)")
			} else {
				fmt.Println("	-", profiles[i].ProfileName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
	profilesCmd.Flags().StringP("add", "a", "", "Add a new profile with name")
	profilesCmd.Flags().StringP("set", "s", "", "Set which profile you want to use")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
