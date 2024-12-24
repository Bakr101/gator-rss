package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("resetting table wasn't successful err: %v", err)
	}
	fmt.Println("Reseting was successful")
	return nil
}


func commandReset() command{
	return command{
		name: "reset",
		handler: []string{},
	}
}