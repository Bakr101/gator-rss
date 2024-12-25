package main

import (
	"context"
	"fmt"
	"html"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.handler) == 0 {
		return fmt.Errorf("the aggregator command expects a single argument, the url. argsLen: %v args:%v", len(cmd.handler), cmd.handler)
	}
	
	feed, err := s.Client.FetchFeed(context.Background(), cmd.handler[0])
	if err != nil {
		fmt.Println(err)
	}
	html.UnescapeString(feed.Channel.Title)
	html.UnescapeString(feed.Channel.Description)
	
	for idx:= range feed.Channel.Item{
		feed.Channel.Item[idx].Title = html.UnescapeString(feed.Channel.Item[idx].Title)
		feed.Channel.Item[idx].Description = html.UnescapeString(feed.Channel.Item[idx].Description)
		
	}
	fmt.Println(feed)
	return nil 
}

func commandAgg()command{
	return command{
		name: "agg",
		handler: []string{},
	}
}