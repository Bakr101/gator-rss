package main

import (
	"context"
	"fmt"

	"github.com/Bakr101/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.handler) != 0 {
		return fmt.Errorf("the following command expects no arguments, argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows by user id in handlerFollowing, err: %v", err)
	}
	fmt.Printf("%v is following:\n", user.Name)
	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}

func commandFollowing() command {
	return command{
		name:    "following",
		handler: []string{},
	}
}