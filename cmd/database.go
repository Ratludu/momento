package cmd

import (
	"log"
	"time"

	"fmt"
	"github.com/spf13/viper"
)

const (
	sqliteLayoutSimple      = "2006-01-02 15:04:05"     // Example: "2025-07-27 18:35:45"
	sqliteLayoutMillis      = "2006-01-02 15:04:05.000" // Example: "2025-07-27 18:35:45.123"
	sqliteLayoutRFC3339     = time.RFC3339              // Example: "2025-07-27T18:35:45+10:00"
	sqliteLayoutRFC3339Nano = time.RFC3339Nano          // Example: "2025-07-27T18:35:45.123456789+10:00"
)

func GetDbPath() string {

	viper.AutomaticEnv()
	dbPath := viper.GetString("MOMENTO_DB")
	if dbPath == "" {
		log.Fatal("No database found. Please export to a location where you would like the database e.g. export MOMENTO_DB=~/.local/share/momento/momemto.db")
	}
	return dbPath
}

func ConvertTime(t time.Time) string {
	newFormat := t.Format(sqliteLayoutMillis)
	return newFormat
}

func ParseTime(timeStr string) (time.Time, error) {

	layoutsToTry := []string{
		sqliteLayoutRFC3339Nano,
		sqliteLayoutRFC3339,
		sqliteLayoutMillis,
		sqliteLayoutSimple,
	}

	for _, layout := range layoutsToTry {
		t, err := time.Parse(layout, timeStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse time string '%s' with any known layout", timeStr)
}
