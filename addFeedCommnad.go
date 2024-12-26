package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bakr101/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.handler) < 2 {
		return fmt.Errorf("the addFeed handler expects two arguments, the feed name and the feed url. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("user not registerd, please register first")
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
	fmt.Printf("New Feed created: %v\n", userFeed)
	
	return nil
}

func commandAddFeed() command{
	return command{
		name: "addfeed",
		handler: []string{},
	}
}