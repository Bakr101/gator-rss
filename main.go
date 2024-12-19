package main

import (
	"fmt"
	"os"

	"github.com/Bakr101/gator/internal/config"
)



func main(){
	gatorConfig, err := config.Read()
	if err != nil{
		fmt.Printf("error getting gatorConfig error: %v", err)
	}
	configState := state{
		cfg: &gatorConfig,
	}
	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	login := commandLogin()
	commands.register(login.name, handlerLogin)
	fullArgs := os.Args
	if len(fullArgs) < 2 {
		fmt.Println("not enough arguments provided.")
		os.Exit(1)
	}
	//fmt.Println(fullArgs)
	commandName := fullArgs[1]
	args := fullArgs[2:]
	//fmt.Println(commandName)

	if commandName == "login"{
		if len(args) < 1{
			fmt.Println("username required.")
			os.Exit(1)
		}
		login.handler = append(login.handler, args...)
		err:= commands.run(&configState, login)
		if err != nil {
			fmt.Printf("error running login command, error: %v\n", err)
			os.Exit(1)
		}
	}
	

	
}