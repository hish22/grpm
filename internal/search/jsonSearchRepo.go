package search

import (

	// "hish22/grpm/internal/cache"

	"encoding/hex"
	"hish22/grpm/internal/cache"
	"hish22/grpm/internal/packet"
	"hish22/grpm/internal/serialization"
	"hish22/grpm/internal/util"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"lukechampine.com/blake3"
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

func fetchFromCache(jsonSearchResult *Response, link string) bool {
	hashed := blake3.Sum256([]byte(link))
	BlakeHexVersion := hex.EncodeToString(hashed[:])
	blob, exists := cache.FetchBlob([]byte(BlakeHexVersion))
	if exists {
		if blob.Expire.After(time.Now()) {
			buf := cache.ReadBlob(&blob.Location)
			serialization.JsonUnserialization(buf, &jsonSearchResult)
			return true
		}
	}
	return false
}

func JsonSearchRepo(repo *packet.RepoInfo) []packet.Srepo {
	link := searchLink(repo)
	var jsonSearchResult Response
	if !fetchFromCache(&jsonSearchResult, link) {
		req, _ := http.NewRequest("GET", link, nil)
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
		serialization.JsonUnserialization(buf, &jsonSearchResult)

		cache.NewCache(link, jsonSearchResult)
	}

	return convertToSearchRepo(&jsonSearchResult)
}
