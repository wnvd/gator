package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wnvd/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		fmt.Printf("usage: %s <feed-name> <url>\n", cmd.name)
		os.Exit(1)
	}

	feedName := cmd.args[0]
	url := cmd.args[1]

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
	}

	_, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("unable to create feed for the given user %w\n", err)
	}

	followFeed := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followFeed)
	if err != nil {
		return fmt.Errorf("failed to follow added feed: %w\n", err)
	}

	fmt.Printf("Feed Added and Followed Successfully\n")
	fmt.Println("--------------------------")
	fmt.Printf("Feed name: (%v)\n", feedName)
	fmt.Printf("Feed URL: (%v)\n", url)
	fmt.Printf("Feed user: (%v)\n", user.Name)

	return nil
}

func handlerShowFeeds(s *state, _ command) error {

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get all the feeds records: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds available")
		os.Exit(1)
	}

	for _, feed := range feeds {
		fmt.Println("Feed: ", feed.Name)
		fmt.Println("URL: ", feed.Url)
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("unable to fetch user by ID provided from feed: %w\n", err)
		}
		fmt.Println("User: ", user.Name)
		fmt.Println()
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		fmt.Printf("usage: %s <url>\n", cmd.name)
		os.Exit(1)
	}
	
	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("unable to get feed details: %w", err)
	}

	followfeed := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followfeed)
	if err != nil {
		return fmt.Errorf("unable to link feed to the current user:  %w", err)
	}

	fmt.Printf("User: %s FOLLOWS feed: %s\n", user.Name, feed.Url)

	return nil
}

func handlerListUserFeeds(s *state, cmd command, user database.User) error {
	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get user feeds: %w", err)
	}

	for _, feed := range userFeeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <feed-name>\n", cmd.name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollow := database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err = s.db.DeleteFeedFollowByUserAndURL(context.Background(), feedFollow); err != nil {
		return err
	}

	fmt.Printf("%v has been unfollowed\n", feed.Url)

	return nil
}
