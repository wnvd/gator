package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wnvd/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markedfeed := database.MarkFeedFetchedParams {
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: time.Now().UTC(),
		ID: feed.ID,
	}

	if err = s.db.MarkFeedFetched(context.Background(), markedfeed); err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("unable to fetch rss feed from the given url")
		return err
	}

	fmt.Println("Channel Title :",rssFeed.Channel.Title)
	fmt.Println("Channel Link :",rssFeed.Channel.Link)
	fmt.Println("Channel Description :",rssFeed.Channel.Description)

	for _, val := range rssFeed.Channel.Item {
		fmt.Println("Item Title: ", val.Title)
		fmt.Println("Item Link: ", val.Link)
		fmt.Println("Item PubDate: ", val.PubDate)
		fmt.Println("Item Description: ", val.Description)
		fmt.Println()
	}

	return nil 
}
