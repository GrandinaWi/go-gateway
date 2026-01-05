package router

import (
	"net/http"

	"gateway/internal/auth"
	"gateway/internal/proxy"
)

func New(secret []byte, userURL, catalogURL, ordersURL string) http.Handler {
	mux := http.NewServeMux()

	// публичные маршруты
	mux.Handle("/login", proxy.New(userURL))
	mux.Handle("/register", proxy.New(userURL))

	// защищённые
	mux.Handle("/products", proxy.New(catalogURL))
	mux.Handle("/orders", auth.Middleware(secret, proxy.New(ordersURL)))

	return mux
}
