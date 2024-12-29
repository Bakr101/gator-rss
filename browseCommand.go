package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Bakr101/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.handler) == 1{
		limitArg, err := strconv.Atoi(cmd.handler[0])
		if err != nil{
			return err
		}
		limit = limitArg
	}
	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return err
	} 
	
	for _, post := range posts{
		fmt.Printf("Title: %v\n", post.Title)
		fmt.Printf("Description: %v\n", post.Description)
		fmt.Printf("Published At: %v\n", post.PublishedAt.Time)
		fmt.Printf("URL: %v\n", post.Url)
		fmt.Println()
	}
	return nil
}

func commandBrowse() command{
	return command{
		name: "browse",
		handler: []string{},
	}
}