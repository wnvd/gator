package main

import (
	"fmt"
	"github.com/wnvd/gator/internal/config"
)

const (
	USERNAME = "naveed"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("unable to read file %w", err)
	}

	if err := config.SetUser(USERNAME, cfg); err != nil {
		fmt.Println("unable to set username %w", err)
	}
}
