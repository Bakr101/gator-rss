package fetch

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

func (c *Client) FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	req, err := http.NewRequestWithContext(context.Background(), "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssResponse := &RSSFeed{}
	err = xml.Unmarshal(dat, &rssResponse)
	if err != nil {
		return &RSSFeed{}, err
	}

	return rssResponse, nil
	
}