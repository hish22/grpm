package search

import (

	// "hish22/grpm/internal/cache"

	"hish22/grpm/internal/core"
	"hish22/grpm/internal/persistance"
	"hish22/grpm/internal/serialization"
	"hish22/grpm/internal/util"
	"io"
	"log"
	"strconv"
)

/* Search repo struct */
type Srepo struct {
	Name        string
	Description string
	Stars       string
	Owner       string
}

/* Repo command info */
type RepoInfo struct {
	Name      string
	Page      string
	MostStars bool
	FewStars  bool
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

func JsonSearchRepo(repo *RepoInfo) ([]Srepo, error) {
	link := core.SearchLink(&repo.Name, &repo.Page, &repo.MostStars, &repo.FewStars)
	var jsonSearchResult Response
	if !persistance.FetchFromCache(&jsonSearchResult, link) {
		resp, err := core.NewJsonReq(&link)
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
