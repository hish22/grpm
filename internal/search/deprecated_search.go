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
