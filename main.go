package main

import (
	"fmt"
	"os"

	"github.com/wnvd/gator/internal/config"
)

const (
	USERNAME = "naveed"
)

func main() {

	var st state
	var err error
	st.cfg, err = config.Read()
	if err != nil {
		fmt.Println("unable to read file %w", err)
	}

	var c commands
	c.cmds = make(map[string]func(*state, command) error)

	// Registering commands
	c.register("login", handlerLogin)


	args := os.Args

	if len(args) < 2 {
		fmt.Println("bad command")
		os.Exit(1)
	}

	cmd := command{
		name:	args[1],
		args:	args[2:],
	}

	if err := c.run(&st, cmd); err != nil {
		fmt.Printf("Unable to run the command %v\n", err)
	}
}
