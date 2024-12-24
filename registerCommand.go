package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bakr101/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.handler) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	
	
	dbUser := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.handler[0],
	}

	if user, err := s.db.GetUser(context.Background(), dbUser.Name); err != nil{
		user, err = s.db.CreateUser(context.Background(), dbUser)
		if err != nil {
			return fmt.Errorf("error creating user in s.db.CreateUser func, err: %v", err)
		}
		err = s.cfg.SetUser(user.Name)
		if err != nil{
			return fmt.Errorf("%v", err)
		}
		fmt.Printf("New User created: %v\n", user)
	}else{
		return fmt.Errorf("user already created, user: %v", dbUser)
	}


	return nil
}


func commandRegister() command{
	return command{
		name: "register",
		handler: []string{},
	}
}