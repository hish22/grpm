package packet

import "net/url"

/* Seach repo struct */
type Srepo struct {
	Name        string
	Description string
	Url         url.URL
}

/* repo's page struct */
type Repo struct {
	Name   string
	Url    url.URL
	Readme string
}
