package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key

//GoDotEnvVariable  . .
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

/*
//GetEnv . . .
func GetEnv() {
	// os package

	// godotenv package
	dotenv := GoDotEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}
*/
