
package main

import (
	"io"
	"encoding/xml"
	"time"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
		Description string		`xml:"description"`
		Language	string		`xml:"language"`
		Item		[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title		string	`xml:"title"`
	Link		string	`xml:"link"`
	Description string	`xml:"description"`
	PubDate 	string	`xml:"pubDate"`	
}

func urlToFeed(url string) (RSSFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}

