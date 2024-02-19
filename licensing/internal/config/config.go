package config

import (
    "github.com/joho/godotenv"
    "fmt"
    "os"
)

func verifyEnv(key string) {
    if os.Getenv(key) == "" {
        fmt.Println("Environment variable " + key + " is not set: exiting")
        os.Exit(1)
    }
}

func Init() {
    err := godotenv.Load()
    verifyEnv("ADMIN_KEY")
    if err != nil {
        fmt.Println("Error loading .env file")
        os.Exit(1)
    }
}
