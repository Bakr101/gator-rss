package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bakr101/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.handler) != 2 {
		return fmt.Errorf("the addFeed handler expects two arguments, the feed name and the feed url. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	
	newFeed := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.handler[0],
		Url: cmd.handler[1],
		UserID: user.ID,
	}
	userFeed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return fmt.Errorf("error creating feed in s.db.CreateFeed func, err: %v", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: userFeed.ID,
		UserID: user.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow in s.db.CreateFeedFollow func, err: %v", err)
	}
	fmt.Printf("New Feed created & followed: %v\n", userFeed)
	
	return nil
}

func commandAddFeed() command{
	return command{
		name: "addfeed",
		handler: []string{},
	}
}