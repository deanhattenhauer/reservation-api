package main

import "github.com/joho/godotenv"

func main() {
	// Load environment variables from .env file before reading any config.
	// Must be called before os.Getenv to ensure variables are available.
	godotenv.Load()
}