package router

import (
	"net/http"

	"gateway/internal/auth"
	"gateway/internal/proxy"
)

func New(secret []byte, userURL, catalogURL, ordersURL string) http.Handler {
	mux := http.NewServeMux()

	// USER
	mux.Handle("/login", proxy.New(userURL))
	mux.Handle("/register", proxy.New(userURL))
	mux.Handle("/user", auth.Middleware(secret, proxy.New(userURL)))

	// CATALOG
	mux.Handle("/products", proxy.New(catalogURL))
	mux.Handle("/products/", proxy.New(ordersURL))

	// ORDERS
	mux.Handle("/orders", auth.Middleware(secret, proxy.New(ordersURL)))
	mux.Handle("/orders/", auth.Middleware(secret, proxy.New(ordersURL)))
	mux.Handle("/health", proxy.New(ordersURL))

	return mux
}
