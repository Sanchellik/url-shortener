package main

import (
	"urlshortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	_ = cfg

	// todo init logger (slog)

	// todo init storage (sqlite)

	// todo init router (chi)

	// todo run server
}
