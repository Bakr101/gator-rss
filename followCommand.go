package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bakr101/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.handler) != 1 {
		return fmt.Errorf("the follow command expects a single argument, the url. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	url := cmd.handler[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url in handlerFollow, err: %v", err)
	}
	
	feedFollowParams:= database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: feed.ID,
		UserID: user.ID,
	}
	followFeed, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow in handlerFollow, err: %v", err)
	}
	fmt.Printf("%v followed %v\n", user.Name, followFeed.FeedName)
	return nil
}

func commandFollow() command {
	return command{
		name: "follow",
		handler: []string{},
	}
}