package search

import (
	"fmt"
	"hish22/grpm/internal/core"
	"hish22/grpm/internal/packet"
	"strings"

	"github.com/gocolly/colly/v2"
)

func SearchRepo(query string, page string) []packet.Srepo {
	gc := core.Collector(query)

	/* Create a slice of Srepos */
	var matchedQueries []packet.Srepo
	var srp packet.Srepo // init a new search repo
	gc.OnHTML("span.search-match", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("class"), "hkFRpV") { // append the name of a repo
			srp.Name = e.Text
		}
		if strings.Contains(e.Attr("class"), "dVFwsC") { // append the description of a repo
			srp.Description = e.Text
		}
		if srp.Name != "" && srp.Description != "" {
			matchedQueries = append(matchedQueries, srp) // Append a clone of srp
		}
	})

	gc.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	gc.Visit("https://github.com/search?q=" + query + "&type=repositories" + "&p=" + page)

	return matchedQueries
}
