package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alcb1310/go-rest-api/internal/api"
	"github.com/joho/godotenv"
	"gitlab.com/0x4149/logz"
)

const local bool = true

var srv *http.Server

func init() {
	if local {
		logz.VerbosMode()
	}

	logz.Run()
}

func main() {
	// Load Environment Variables
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

	router := api.NewRouter()
	router.Prefix("/auth").AuthEndpoints()

	if local {
		srv = &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      router,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}
	}

	logz.Info("Server is running")
	logz.Fatal(srv.ListenAndServe())
}
