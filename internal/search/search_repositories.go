package search

import (
	corehttp "hish22/grpm/internal/coreHttp"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/serialization"
	"hish22/grpm/internal/util"
	"io"
	"log"
	"strconv"
)

type Srepo struct {
	Name        string
	Description string
	Stars       string
	Owner       string
}

type RepoInfo struct {
	Name  string
	Page  string
	Sort  string
	Order string
}

type Repository struct {
	Name       string `json:"name"`
	OwnerLogin string `json:"owner_login"`
}

type Repo struct {
	Repository Repository `json:"repository"`
}

type Results struct {
	Repo        Repo   `json:"repo"`
	Description string `json:"hl_trunc_description"`
	Stars       int    `json:"followers"`
}

type Payload struct {
	Results []Results `json:"results"`
}

type Response struct {
	Payload Payload `json:"payload"`
}

func convertToSearchRepo(jsonRepo *Response) []Srepo {
	var listOfSrepo []Srepo
	for _, r := range jsonRepo.Payload.Results {
		listOfSrepo = append(listOfSrepo, Srepo{
			Name:        r.Repo.Repository.Name,
			Description: util.CleanHtmlTags(r.Description),
			Stars:       strconv.Itoa(r.Stars),
			Owner:       r.Repo.Repository.OwnerLogin,
		})
	}
	return listOfSrepo
}

type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"login"`
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	Type     string `json:"type"`
}

type Items struct {
	ID                  int    `json:"id"`
	NodeID              string `json:"node_id"`
	Name                string `json:"name"`
	Owner               Owner  `json:"owner"`
	FullName            string `json:"full_name"`
	Private             bool   `json:"private"`
	HtmlUrl             string `json:"html_url"`
	Description         string `json:"description"`
	Stars               int    `json:"stargazers_count"`
	Watchers            int    `json:"watchers_count"`
	Forks               int    `json:"forks"`
	ProgrammingLanguage string `json:"language"`
}

type Repositories struct {
	TotalCount int     `json:"total_count"`
	TotalItems []Items `json:"items"`
}

/* Search Repositories by requesting api.github */
func SearchRepositories(metadata *RepoInfo) *Repositories {
	var repositories Repositories
	link := corehttp.RequestLink{
		Base:      corehttp.ApiLink,
		Endpoints: []string{"search", "repositories"},
		Queries: []string{"q=" + metadata.Name, "sort=" +
			metadata.Sort, "order=" + metadata.Order,
			"page=" + metadata.Page},
	}.Build()
	if !persistance.FetchFromCache(&repositories, *link) {
		corehttp.Request(link, &repositories)
		persistance.NewCache(*link, &repositories)
	}
	return &repositories
}

func JsonSearchRepo(repo *RepoInfo) ([]Srepo, error) {
	MostStars := false
	FewStars := false
	link := corehttp.SearchLink(&repo.Name, &repo.Page, &MostStars, &FewStars)
	var jsonSearchResult Response
	if !persistance.FetchFromCache(&jsonSearchResult, link) {
		resp, err := corehttp.NewJsonReq(&link)
		// Handle http request error
		if err != nil {
			return []Srepo{}, err
		}

		defer resp.Body.Close()

		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Can't fetch repos JSON data, ", err)
		}
		serialization.JsonUnserialization(buf, &jsonSearchResult)

		persistance.NewCache(link, &jsonSearchResult)
	}

	return convertToSearchRepo(&jsonSearchResult), nil
}
