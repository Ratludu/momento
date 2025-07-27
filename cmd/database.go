package cmd

import (
	"log"
	"time"

	"github.com/spf13/viper"
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
	sqliteLayout := "2006-01-02 15:04:05.000"
	newFormat := t.Format(sqliteLayout)
	return newFormat
}
