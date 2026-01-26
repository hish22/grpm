package core

import (
	"time"

	"github.com/gocolly/colly/v2"
)

/* Build a new Collector to start scraping github */
func Collector(cacheName string) *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains("github.com"),
		colly.CacheDir("./github_"+cacheName),
		colly.CacheExpiration(24*time.Hour),
	)
}
