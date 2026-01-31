package search

import (

	// "hish22/grpm/internal/cache"

	"hish22/grpm/internal/cache"
	"hish22/grpm/internal/core"
	"hish22/grpm/internal/packet"
	"hish22/grpm/internal/serialization"
	"hish22/grpm/internal/util"
	"io"
	"log"
	"strconv"
)

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

func convertToSearchRepo(jsonRepo *Response) []packet.Srepo {
	var listOfSrepo []packet.Srepo
	for _, r := range jsonRepo.Payload.Results {
		listOfSrepo = append(listOfSrepo, packet.Srepo{
			Name:        r.Repo.Repository.Name,
			Description: util.CleanHtmlTags(r.Description),
			Stars:       strconv.Itoa(r.Stars),
			Owner:       r.Repo.Repository.OwnerLogin,
		})
	}
	return listOfSrepo
}

func JsonSearchRepo(repo *packet.RepoInfo) ([]packet.Srepo, error) {
	link := core.SearchLink(repo)
	var jsonSearchResult Response
	if !cache.FetchFromCache(&jsonSearchResult, link) {
		resp, err := core.NewJsonReq(&link)
		// Handle http request error
		if err != nil {
			return []packet.Srepo{}, err
		}

		defer resp.Body.Close()

		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Can't fetch repos JSON data, ", err)
		}
		serialization.JsonUnserialization(buf, &jsonSearchResult)

		cache.NewCache(link, &jsonSearchResult)
	}

	return convertToSearchRepo(&jsonSearchResult), nil
}
