package structures

import (
	"encoding/json"
	"testing"
)

func TestRepository_JSONMarshal(t *testing.T) {
	repo := Repository{
		ID:       12345,
		NodeID:   "MDM6UmVmMTIzNDU=",
		Name:     "myrepo",
		FullName: "owner/myrepo",
		Private:  false,
		Owner: Owner{
			ID:       111,
			Username: "owner",
			Url:      "https://api.github.com/users/owner",
			HtmlUrl:  "https://github.com/owner",
			Type:     "User",
		},
		HtmlUrl:             "https://github.com/owner/myrepo",
		Description:         "A cool repository",
		Stars:               100,
		Watchers:            50,
		Forks:               25,
		ProgrammingLanguage: "Go",
		License: license{
			Key:    "mit",
			Name:   "MIT License",
			SpdxID: "MIT",
			Url:    "https://opensource.org/licenses/MIT",
		},
		Topics: []string{"go", "cli", "tool"},
	}

	data, err := json.Marshal(repo)
	if err != nil {
		t.Fatalf("Failed to marshal Repository: %v", err)
	}

	if len(data) == 0 {
		t.Error("Marshaled data is empty")
	}

	var result Repository
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Repository: %v", err)
	}

	if result.Name != "myrepo" {
		t.Errorf("Expected Name myrepo, got %s", result.Name)
	}
	if result.FullName != "owner/myrepo" {
		t.Errorf("Expected FullName owner/myrepo, got %s", result.FullName)
	}
	if result.Owner.Username != "owner" {
		t.Errorf("Expected Owner.Username owner, got %s", result.Owner.Username)
	}
	if result.Stars != 100 {
		t.Errorf("Expected Stars 100, got %d", result.Stars)
	}
	if len(result.Topics) != 3 {
		t.Errorf("Expected 3 topics, got %d", len(result.Topics))
	}
}

func TestRepository_JSONUnmarshal(t *testing.T) {
	jsonData := `{
		"id": 12345,
		"node_id": "MDM6UmVmMTIzNDU=",
		"name": "myrepo",
		"full_name": "owner/myrepo",
		"private": false,
		"owner": {
			"id": 111,
			"login": "owner",
			"url": "https://api.github.com/users/owner",
			"html_url": "https://github.com/owner",
			"type": "User"
		},
		"html_url": "https://github.com/owner/myrepo",
		"description": "A cool repository",
		"stargazers_count": 100,
		"watchers_count": 50,
		"forks": 25,
		"language": "Go",
		"license": {
			"key": "mit",
			"name": "MIT License",
			"spdx_id": "MIT",
			"url": "https://opensource.org/licenses/MIT"
		},
		"topics": ["go", "cli", "tool"]
	}`

	var repo Repository
	err := json.Unmarshal([]byte(jsonData), &repo)
	if err != nil {
		t.Fatalf("Failed to unmarshal Repository: %v", err)
	}

	if repo.ID != 12345 {
		t.Errorf("Expected ID 12345, got %d", repo.ID)
	}
	if repo.Name != "myrepo" {
		t.Errorf("Expected Name myrepo, got %s", repo.Name)
	}
	if repo.Owner.Username != "owner" {
		t.Errorf("Expected Owner.Username owner, got %s", repo.Owner.Username)
	}
	if repo.ProgrammingLanguage != "Go" {
		t.Errorf("Expected ProgrammingLanguage Go, got %s", repo.ProgrammingLanguage)
	}
	if repo.License.SpdxID != "MIT" {
		t.Errorf("Expected License.SpdxID MIT, got %s", repo.License.SpdxID)
	}
}

func TestRepositories_JSONMarshal(t *testing.T) {
	repos := Repositories{
		TotalCount: 2,
		Repositories: []Repository{
			{
				ID:       1,
				Name:     "repo1",
				FullName: "owner/repo1",
				Owner:    Owner{Username: "owner"},
			},
			{
				ID:       2,
				Name:     "repo2",
				FullName: "owner/repo2",
				Owner:    Owner{Username: "owner"},
			},
		},
	}

	data, err := json.Marshal(repos)
	if err != nil {
		t.Fatalf("Failed to marshal Repositories: %v", err)
	}

	var result Repositories
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Repositories: %v", err)
	}

	if result.TotalCount != 2 {
		t.Errorf("Expected TotalCount 2, got %d", result.TotalCount)
	}
	if len(result.Repositories) != 2 {
		t.Errorf("Expected 2 repositories, got %d", len(result.Repositories))
	}
}

func TestOwner_Fields(t *testing.T) {
	owner := Owner{
		ID:       111,
		Username: "testuser",
		Url:      "https://api.github.com/users/testuser",
		HtmlUrl:  "https://github.com/testuser",
		Type:     "User",
	}

	if owner.ID != 111 {
		t.Errorf("Expected ID 111, got %d", owner.ID)
	}
	if owner.Username != "testuser" {
		t.Errorf("Expected Username testuser, got %s", owner.Username)
	}
	if owner.Type != "User" {
		t.Errorf("Expected Type User, got %s", owner.Type)
	}
}

func TestLicense_Fields(t *testing.T) {
	lic := license{
		Key:    "apache-2.0",
		Name:   "Apache License 2.0",
		SpdxID: "Apache-2.0",
		Url:    "https://www.apache.org/licenses/LICENSE-2.0",
	}

	if lic.Key != "apache-2.0" {
		t.Errorf("Expected Key apache-2.0, got %s", lic.Key)
	}
	if lic.SpdxID != "Apache-2.0" {
		t.Errorf("Expected SpdxID Apache-2.0, got %s", lic.SpdxID)
	}
}
