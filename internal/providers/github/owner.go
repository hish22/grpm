package github

/*
 * Owner Information struct
 *
 * A struct that provides a basic author
 * information for any owner belong
 * to the github platform.
 */

type owner struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Url   string `json:"html_url"`
}
