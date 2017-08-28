package config

import "os"

// DbCredentials from env or dev defaults
func DbCredentials() map[string]string {
	m := map[string]string{
		"host":     os.Getenv("CASE_DB_HOST"),
		"user":     os.Getenv("CASE_DB_USER"),
		"password": os.Getenv("CASE_DB_PASSWORD"),
	}
	if m["host"] == "" {
		m["host"] = "localhost"
	}
	if m["user"] == "" {
		m["user"] = "user"
	}
	if m["password"] == "" {
		m["password"] = "password"
	}
	return m
}

// JWTKey ... from env or insecure default
func JWTKey() string {
	key := os.Getenv("CASE_JWT_KEY")
	if key == "" {
		key = "InsecurePrivateKey"
	}
	return key
}
