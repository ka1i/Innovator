package main

import (
	"os"

	"github.com/ka1i/innovator/internal/app"
)

func main() {
	innovator := app.App
	os.Exit(innovator.Innovator())
}
