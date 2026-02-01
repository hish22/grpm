package packet

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
	Stars       string
	Owner       string
}

/* repo's page info struct */
type RepoPageInfo struct {
	ID        int
	RepoName  string
	Owner     string
	Link      string
	CreatedAt string
	Readme    string
}
