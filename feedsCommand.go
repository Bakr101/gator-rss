package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error{
	if len(cmd.handler) != 0 {
		return fmt.Errorf("the feeds command expects no arguments, argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds in handlerFeeds, err: %v", err)
	}
	for _, feed := range feeds{
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user name, err: %v", err)
		}
		fmt.Println(user.Name)
	}
	
	return nil
}

func commandFeeds() command{
	return command{
		name: "feeds",
		handler: []string{},
	}
}