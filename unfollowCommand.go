package main

import (
	"context"
	"fmt"

	"github.com/Bakr101/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.handler) != 1 {
		return fmt.Errorf("the unfollow command expects one argument, the url. ArgsLen: %v Args: %v", len(cmd.handler), cmd.handler)
	}
	url := cmd.handler[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed not found, err: %v", err)
	}
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error deleting feed follow, err: %v", err)
	}
	fmt.Printf("%v unfollowed %v\n", user.Name, feed.Name)
	return nil
}

func commandUnfollow() command {
	return command{
		name:    "unfollow",
		handler: []string{},
	}
}