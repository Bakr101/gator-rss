package main

import (
	"fmt"

	"github.com/Bakr101/gator/internal/config"
)

type state struct{
	cfg	*config.Config
}

type command struct{
	name		string
	handler 	[]string
}

type commands struct{
	handlers	map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error){
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error{
	commandFunc := c.handlers[cmd.name]
	err := commandFunc(s, cmd)
	if err != nil {
		return fmt.Errorf("can't run function, error: %v", err)
	}
	return nil
}