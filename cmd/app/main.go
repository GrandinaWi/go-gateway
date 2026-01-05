package main

import (
	"log"
	"net/http"

	"gateway/internal/config"
	"gateway/internal/router"
)

func main() {
	cfg := config.Load()

	handler := router.New(
		cfg.JWTSecret,
		cfg.UserAPIURL,
		cfg.CatalogURL,
		cfg.OrdersURL,
	)

	log.Println("gateway listening on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
