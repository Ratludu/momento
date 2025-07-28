package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed templates/momento.db
var embeddedDBFS embed.FS

func getUserDBPath(dbName string) (string, error) {

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config directory: %w", err)
	}

	appDir := filepath.Join(configDir, "momento") // Create a subdirectory for your app
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create app config directory: %w", err)
	}

	return filepath.Join(appDir, dbName), nil
}

// createDbCmd represents the createDb command
var createDbCmd = &cobra.Command{
	Use:   "createDb",
	Short: "Initiates a database in your config directory.",
	Run: func(cmd *cobra.Command, args []string) {

		localDBPath, err := getUserDBPath("momento.db")
		if err != nil {
			log.Fatalf("Error getting user DB path: %v", err)
		}

		fmt.Printf("Local database path: %s\n", localDBPath)

		_, err = os.Stat(localDBPath)
		if os.IsNotExist(err) {

			fmt.Println("Local database not found. Copying from embedded template...")

			embeddedDBData, err := fs.ReadFile(embeddedDBFS, "templates/momento.db")
			if err != nil {
				log.Fatalf("Error reading embedded database: %v", err)
			}

			err = os.WriteFile(localDBPath, embeddedDBData, 0644) // 0644 for rw-r--r--
			if err != nil {
				log.Fatalf("Error writing local database file: %v", err)
			}
			fmt.Println("Local database copied successfully.")
			fmt.Println("Please Export the path as MOMENTO_DB in your .bashrc file or equivalent. Below is a template:")
			fmt.Printf("export MOMENTO_DB=%s\n", localDBPath)
		} else if err != nil {
			log.Fatalf("Error checking local database existence: %v", err)
		} else {
			fmt.Println("Local database already exists. Using existing file.")
		}
	},
}

func init() {
	rootCmd.AddCommand(createDbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
