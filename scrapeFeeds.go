package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Bakr101/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error{
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markedFeedParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		ID: nextFeed.ID,
	}
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), markedFeedParams)
	if err != nil {
		return err
	}
	
	feedData, err := s.Client.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}
	
	for _, post := range feedData.Channel.Item{
		publicationTime := sql.NullTime{}
		if time, err := time.Parse(time.RFC1123Z, post.PubDate); err == nil{
			publicationTime = sql.NullTime{
				Time: time,
				Valid: true,
			}
		}
		params := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			Title: post.Title,
			Url: post.Link,
			Description: post.Description,
			PublishedAt: publicationTime,
			FeedID: markedFeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Error creating post, err: %v\n", err)
		}
	}
	
	return nil
}