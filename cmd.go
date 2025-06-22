package main

import (
	"errors"

	"github.com/wnvd/gator/internal/config"
	"github.com/wnvd/gator/internal/database"
)

type state struct {
	db		*database.Queries
	cfg		*config.Config
}

type command struct {
	name	string
	args	[]string
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
