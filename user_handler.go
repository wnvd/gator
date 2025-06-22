package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wnvd/gator/internal/config"
	"github.com/wnvd/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("the login handler expects a single argument, the username.")
		os.Exit(1)
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := config.SetUser(cmd.args[0], s.cfg); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("the register handler expects a single argument, the username.")
		os.Exit(1)
	}

	data := database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), data)
	if err != nil {
		fmt.Println("Unable to create user of name ", cmd.args[0], ": ", err)
		os.Exit(1)
	}

	// setting current user to the provided name
	if err := config.SetUser(user.Name, s.cfg); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("%v has been registered and set as the current user\n", user.Name)

	printUser(user)

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:		%v\n", user.Name)
}
