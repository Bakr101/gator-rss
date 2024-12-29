package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.handler) == 0 {
		return fmt.Errorf("the aggregator command expects a single argument, the time between requests. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	
	timeBetweenReqs, err := time.ParseDuration(cmd.handler[0])
	if err != nil {
		return fmt.Errorf("incorrect duration: %v", err)
	}

	fmt.Printf("Collecting feeds every: %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for range ticker.C {

	scrapeFeeds(s)

	}
	
	return nil 
}

func commandAgg()command{
	return command{
		name: "agg",
		handler: []string{},
	}
}