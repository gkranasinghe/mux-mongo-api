package configs

import (
	"fmt"
	"os"
	// "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }
    fmt.Printf("MONGOURI is : %s, \n", os.Getenv("MONGOURI"))
    return os.Getenv("MONGOURI")
}