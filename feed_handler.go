package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wnvd/gator/internal/database"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		fmt.Println("addfeed command needs <user-name> and <url>")
		os.Exit(1)
	}

	userName := s.cfg.CurrentUserName
	feedName := cmd.args[0]
	url := cmd.args[1]


	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("unable to get the user with the provided name %w", err)
	}
	
	feed := database.CreateFeedParams{
		ID: uuid.New(),
		Name: feedName,
		Url: url,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err = s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("unable to create feed for the given user %w", err)
	}

	fmt.Printf("Feed Added Successfully\n")
	fmt.Println("--------------------------")
	fmt.Printf("Feed name: (%v)\n", feedName)
	fmt.Printf("Feed URL: (%v)\n", url)
	fmt.Printf("Feed user: (%v)\n", userName)

	return nil
}
