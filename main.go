package main

import (
	"flag"

	"github.com/LaurenceUsas/riverisland-webcrawler/crawler"
)

var (
	fURL      = flag.String("u", "https://en.wikipedia.org/wiki/Main_Page", "URL of website to crawl") // Start URL
	fLimit    = flag.Int("l", 100, "Page limit to crawl")                                              // Limit of pages to crawl. Default - 100
	fExternal = flag.Bool("e", false, "Allow crawling external links")                                 // Crawl external links. Default - False
	// fDirty    = flag.Bool("dirty", false, "Enabled turns off URL verification before adding.")
)

func main() {
	flag.Parse()
	c := ricrawler.NewCrawler(uint64(*fLimit), *fExternal)
	c.Start(*fURL)
}
