package search

import (
	"encoding/json"
	"hish22/grpm/internal/packet"
	"hish22/grpm/internal/util"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Repository struct {
	Name        string `json:"name"`
	Owner_login string `json:"owner_login"`
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

func unMarshalJsonSearch(buf []byte, structure any) {
	if err := json.Unmarshal(buf, &structure); err != nil {
		log.Fatal("Can't decode repos JSON data, ", err)
	}
}

func convertToSearchRepo(jsonRepo *Response) []packet.Srepo {
	var listOfSrepo []packet.Srepo
	for _, r := range jsonRepo.Payload.Results {
		listOfSrepo = append(listOfSrepo, packet.Srepo{
			Name:        r.Repo.Repository.Owner_login + "/" + r.Repo.Repository.Name,
			Description: util.CleanHtmlTags(r.Description),
			Stars:       strconv.Itoa(r.Stars),
		})
	}
	return listOfSrepo
}

func JsonSearchRepo(repo *packet.RepoInfo) []packet.Srepo {
	req, _ := http.NewRequest("GET", searchLink(repo), nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Can't bring data,", err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Can't fetch repos JSON data, ", err)
	}
	var jsonSearchResult Response
	unMarshalJsonSearch(buf, &jsonSearchResult)

	return convertToSearchRepo(&jsonSearchResult)
}
