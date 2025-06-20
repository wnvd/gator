package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/wnvd/gator/internal/config"
)

type state struct {
	cfg		*config.Config
}

type command struct {
	name	string
	args	[]string
}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("the login handler expects a single argument, the username.")
		os.Exit(1)
	}

	if err := config.SetUser(cmd.args[0], s.cfg); err != nil {
		return err
	}

	fmt.Printf("User %v has been set\n", cmd.args[0])
	
	return nil
}

type commands struct {
	cmdsReg		map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, isPresent := c.cmdsReg[cmd.name]
	if !isPresent {
		return errors.New("command does not exist.")
	}

	if err := fn(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdsReg[name] = f
}
