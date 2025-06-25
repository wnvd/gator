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
		fmt.Println("usage: %s <user-name>", cmd.name)
		os.Exit(1)
	}

	name := cmd.args[0]
	_, err := s.db.GetUserByName(context.Background(), name)
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
		fmt.Println("usage: %s <user-name>", cmd.name)
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

func handlerGetUsers(s *state, cmd command) error {

	names, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get user data: %w", err)
	}

	if len(names) == 0 {
		fmt.Println("no users are registered")
		os.Exit(1)
	}

	for _, name := range names {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("%v (current)\n", name)
			continue
		}
		fmt.Println(name) 
	}

	return nil
}

// NOTE: This handler is for development purpose only.
func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("unable to reset users state: %w", err)
	}

	fmt.Println("User table has been reset successfully")
	return nil
}

func handlerAggregate(s *state, cmd command) error {

	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to fetch rss feed: %w", err)
	}

	fmt.Println("Title :", rssFeed.Channel.Title)
	fmt.Println("Link :", rssFeed.Channel.Link)
	fmt.Println("Desc :", rssFeed.Channel.Description)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println("Title :", item.Title)
		fmt.Println("Link :", item.Link)
		fmt.Println("Desc :", item.Description)
		fmt.Println("PubDate :", item.PubDate)
		fmt.Println("-----")
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:		%v\n", user.Name)
}
