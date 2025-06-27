package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

	for _, val := range rssFeed.Channel.Item {

		// pubdate parsing
		var parsedPubTime sql.NullTime
		parsedPubTime.Time, err = time.Parse(time.RFC3339, val.PubDate)
		parsedPubTime.Valid = true
		if err != nil {
			parsedPubTime.Valid = false
		}

		post := database.CreatePostParams {
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: val.Title,
			Url: val.Link,
			Description: val.Description,
			PublishedAt: parsedPubTime,
			FeedID: markedfeed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), post)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				continue
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	return nil 
}
