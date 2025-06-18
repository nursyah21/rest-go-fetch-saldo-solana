package helper

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	RPC_URI    string
	MONGO_URI  string
	SECRET_KEY string
)

func LoadEnv() {
	file, err := os.Open(".env")
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			os.Setenv(key, value)
		}
	} else {
		log.Println("Warning: .env file not found, using system env vars")
	}

	RPC_URI = os.Getenv("RPC_URI")
	if RPC_URI == "" {
		log.Fatal("RPC_URI is not set in environment or .env file")
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI is not set in environment or .env file")
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY is not set in environment or .env file")
	}
}
