package info

import (
	"hish22/grpm/internal/cache"
	"hish22/grpm/internal/core"
	"hish22/grpm/internal/serialization"
	"io"
	"log"
)

/* repo's page info struct */
type RepoPageInfo struct {
	ID        int
	RepoName  string
	Owner     string
	Link      string
	CreatedAt string
	Readme    string
}

type Readme struct {
	DisplayName string `json:"displayName"`
	RichText    string `json:"richText"`
}

type Tree struct {
	Readme Readme `jsno:"readme"`
}

type Repo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	OwnerLogin string `json:"ownerLogin"`
	CreatedAt  string `json:"createdAt"`
}

type Payload struct {
	Repo Repo `json:"repo"`
	Tree Tree `json:"tree"`
}

type Response struct {
	Payload Payload `json:"payload"`
}

func convertIntoInfoRepo(response *Response, link *string) RepoPageInfo {
	return RepoPageInfo{
		ID:        response.Payload.Repo.ID,
		RepoName:  response.Payload.Repo.Name,
		Owner:     response.Payload.Repo.OwnerLogin,
		Link:      *link,
		CreatedAt: response.Payload.Repo.CreatedAt,
		Readme:    response.Payload.Tree.Readme.RichText,
	}
}

func JsonInfoRepo(owner *string, repo *string) (RepoPageInfo, error) {
	link := core.InfoLink(owner, repo)
	var jsonInfoResult Response
	if !cache.FetchFromCache(&jsonInfoResult, link) {
		resp, err := core.NewJsonReq(&link)
		// Handle http request error
		if err != nil {
			return RepoPageInfo{}, err
		}

		defer resp.Body.Close()
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Can't fetch repo information, ", err)
		}
		serialization.JsonUnserialization(buf, &jsonInfoResult)

		cache.NewCache(link, &jsonInfoResult)
	}

	return convertIntoInfoRepo(&jsonInfoResult, &link), nil
}
