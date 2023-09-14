package main

import (
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/0x4149/logz"
)

const local bool = true

func init() {
	if local {
		logz.VerbosMode()
	}

	logz.Run()
}

func main() {
	mode := os.Getenv("GO_ENV")
	if mode != "production" {
		err := godotenv.Load()
		if err != nil {
			logz.Fatal("Environment variables couldn'g be loaded")
		}
	}

	port, hasPort := os.LookupEnv("PORT")
	if !hasPort {
		logz.Fatal("PORT environment variable couldn't be found")
	}

	logz.Info("Everything is working")
	logz.Debug(port)
}
