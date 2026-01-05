package config

import "os"

type Config struct {
	Port       string
	JWTSecret  []byte
	UserAPIURL string
	CatalogURL string
	OrdersURL  string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("HTTP_PORT", "8000"),
		JWTSecret:  []byte(getEnv("JWT_SECRET", "supersecret")),
		UserAPIURL: getEnv("USER_API_URL", "http://user-api:8080"),
		CatalogURL: getEnv("CATALOG_API_URL", "http://catalog-api:8080"),
		OrdersURL:  getEnv("ORDERS_API_URL", "http://orders-api:8080"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
