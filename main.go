package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Bakr101/gator/internal/config"
	"github.com/Bakr101/gator/internal/database"
	_ "github.com/lib/pq"
)



func main(){
	const db_URL = "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
	//Read gator configuration
	gatorConfig, err := config.Read()
	if err != nil{
		fmt.Printf("error getting gatorConfig error: %v", err)
	}

	//write DB URL & Load DB
	err = gatorConfig.SetUrl(db_URL)
	if err != nil {
		fmt.Printf("error setting URL, error: %v \n", err)
	}

	//DB open a connection & app state

	db, errorr := sql.Open("postgres", gatorConfig.Db_url)
	if err != nil{
		fmt.Printf("error connecting to DB, err: %v", errorr)
	}
	dbQueries := database.New(db)
	configState := state{
		db:  dbQueries,
		cfg: &gatorConfig,
	}
	

	//initialize commands handlers struct & resgister login command
	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	//Registering commands
	login := commandLogin()
	register := commandRegister()
	reset := commandReset()
	users := commandUsers()
	commands.register(login.name, handlerLogin)
	commands.register(register.name, handlerRegister)
	commands.register(reset.name, handlerReset)
	commands.register(users.name, handlerUsers)

	//Get agruments from Cli & splitting them 
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
	
	if commandName == "register"{
		if len(args) < 1{
			fmt.Println("no name provided for user")
			os.Exit(1)
		}
		register.handler = append(register.handler, args...)
		err := commands.run(&configState, register)
		if err != nil {
			fmt.Printf("error running register command, error: %v\n", err)
			os.Exit(1)
		}
	}

	if commandName == "reset"{
		if len(args) > 0 {
			fmt.Println("reset expects no arguments")
			os.Exit(1)
		}
		err := commands.run(&configState, reset)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}

	if commandName == "users"{
		if len(args) > 0 {
			fmt.Println("users expects no arguments")
			os.Exit(1)
		}
		err := commands.run(&configState, users)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
}