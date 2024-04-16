package app

import (
	"flag"
	"log"
	"os"

	"github.com/Muhless/GO-TOKO/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env.example file")
	}

	// port
	// todo : parameter 1=env, parameter 2=defaultnya
	appConfig.AppName = getEnv("APP_NAME", "GoTokoApp")
	appConfig.AppEnv = getEnv("APP_ENV", "Developments")
	appConfig.AppPort = getEnv("APP_PORT", "9000")
	appConfig.AppUrl = getEnv("APP_URL", "http://localhost:9000")

	// database
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "admin")
	dbConfig.DBName = getEnv("DB_NAME", "db_gotoko")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	// cli
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		// server
		server.Initialize(appConfig, dbConfig)
		server.Run(": " + appConfig.AppPort)
	}
}
