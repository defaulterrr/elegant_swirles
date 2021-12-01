package main

import (
	"fmt"
	"os"

	"github.com/defaulterrr/elegant_swirles/dht/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Printf("app.Run: %v", err)
		os.Exit(1)
	}
}
