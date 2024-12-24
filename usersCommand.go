package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("can't get all users, error: %v", err)
	}
	for _, user := range users{
		if user.Name == s.cfg.Current_user_name{
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil 
}

func commandUsers()command{
	return command{
		name: "users",
		handler: []string{},
	}
}