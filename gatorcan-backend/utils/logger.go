package utils

import (
	"log"
	"os"
	"path/filepath"
)

var (
	logger *log.Logger
)

const LogPath = "logs\\app.log"

func init() {

	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(LogPath), os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to create log directory:", err)
	}

	// Create a new log file
	file, err := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	// Create a logger with the log file
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Log() *log.Logger {
	return logger
}
