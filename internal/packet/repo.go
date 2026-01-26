package packet

import "net/url"

/* Repo command info */
type RepoInfo struct {
	Name      string
	Page      string
	MostStars bool
	FewStars  bool
}

/* Search repo struct */
type Srepo struct {
	Name        string
	Description string
}

/* repo's page struct */
type RepoPage struct {
	Name   string
	Url    url.URL
	Readme string
}
