package config

import "os"

func Load() {
	if os.Getenv("MONGO_URI") == "" {
		panic("MONGO_URI is not set")
	}
	if os.Getenv("DB_NAME") == "" {
		panic("DB_NAME is not set")
	}
	if os.Getenv("JWT_SECRET") == "" {
		panic("JWT_SECRET is not set")
	}
}
