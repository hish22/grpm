package gitlab

import (
	"hish22/grpm/internal/services"
	"time"
)

/*
 * Project Information struct
 *
 * A struct that provides basic a building
 * information for any project contained
 * on the gitlab platform.
 */

type project struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
}

func RequestProjectById(p *project, id string) error {
	link, err := services.CreateRequestLink(
		services.GitlabApiBase,
		[]string{"projects", id},
		[]string{},
	)

	if err != nil {
		return err
	}

	apiReq, err := services.CreateApiRequest(
		link,
		map[string]string{
			"Accept":     "application/json",
			"User-Agent": "grpm/0.0.1",
		},
		"GET",
		5*time.Second,
	)

	if err != nil {
		return err
	}

	err = apiReq.RequestWithDecode(&p)

	if err != nil {
		return err
	}

	return nil
}

// func RequestProjectBySearch(p []project, name string, path string) error {

// 	if name == "" || path == "" {
// 		return errors.New("Value of name or path can't be empty")
// 	}

// 	link, err := services.CreateRequestLink(
// 		services.GitlabApiBase,
// 		[]string{"projects"},
// 		[]string{"search=" + strings.Join([]string{name, path}, "/"), "search_namespaces=true"},
// 	)

// 	fmt.Println(link.String())

// 	if err != nil {
// 		return err
// 	}

// 	apiReq, err := services.CreateApiRequest(
// 		link,
// 		map[string]string{
// 			"Accept":     "application/json",
// 			"User-Agent": "grpm/0.0.1",
// 		},
// 		"GET",
// 		5*time.Second,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	err = apiReq.RequestWithDecode(&p)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
