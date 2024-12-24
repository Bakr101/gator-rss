package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.handler) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	user, err := s.db.GetUser(context.Background(),cmd.handler[0])
	if err != nil {
		return fmt.Errorf("user not registerd, please register first")
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil{
		return fmt.Errorf("error setting user in handlerLogin, err: %v", err)
	}
	
	fmt.Printf("New Username: %v\n", s.cfg.Current_user_name)

	return nil
}

func commandLogin() command{
	return command{
		name: "login",
		handler: []string{},
	}
}