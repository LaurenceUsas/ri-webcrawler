package ricrawler

import (
	"io"
	"log"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Crawler struct {
	sync.Mutex
	sync.WaitGroup
	pageLimit     uint64 // up to 18446744073709551615
	pagesVisited  uint64
	allowExternal bool
	visited       map[string]bool
}

func NewCrawler(pageLimit uint64, allowExternal bool) *Crawler {
	if pageLimit < 1 {
		log.Println("Minimal allowed page limit is 1. Limit set to 1")
		pageLimit = 1
	}
	c := &Crawler{
		pageLimit:     pageLimit,
		allowExternal: allowExternal,
		visited:       make(map[string]bool),
	}
	return c
}

func (c *Crawler) Start(pageURL string) {
	// Verify link.
	if _, err := url.Parse(pageURL); err != nil {
		log.Fatal(err)
		return
	}

	// Start Crawling
	c.Add(1)
	c.pagesVisited++
	c.visited[pageURL] = true
	go c.Crawl(pageURL)

	c.Wait()
}

func (c *Crawler) Crawl(crawlURL string) error {
	client := NewHTTPClient() // Reuse connection. If cant, create new for that domain.
	resp, err := client.Get(crawlURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	p, err := c.extractLinks(resp.Request.URL, resp.Body)
	if err != nil {
		return err
	}
	p.Print()
	c.Done()
	return nil
}

// Ignoring <link> tag for now
// <img> x/net/html has problems extracting
//
func (c *Crawler) extractLinks(requestURL *url.URL, data io.Reader) (*Page, error) {
	p := NewPage(requestURL)
	ts := html.NewTokenizer(data)

	for {
		tt := ts.Next()
		if tt == html.ErrorToken { //an EOF
			return p, nil
		}

		if tt == html.StartTagToken {
			t := ts.Token()

			switch t.DataAtom.String() {
			case "title":
				tt := ts.Next()
				if tt == html.TextToken {
					t := ts.Token()
					if p.Title == "" {
						p.Title = strings.TrimSpace(t.Data)
					}
				}
			case "a": // LINKS
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						link, err := url.Parse(attr.Val)
						if err != nil {
							log.Println(err)
							continue
						}
						resolved := p.URL.ResolveReference(link)
						if p.URL.Host == resolved.Host { // Internal
							p.Interal[resolved.String()] = struct{}{} // String instead of *url.URL to resolve duplicates.
							c.tryCrawl(resolved)
						} else { // External
							p.External[resolved.String()] = struct{}{}
							if c.allowExternal {
								c.tryCrawl(resolved)
							}
						}
					}
				}
			case "img", "audio", "script", "video", "embed", "source", "input": // STATIC
				for _, attr := range t.Attr {
					if attr.Key == "src" {
						link, err := url.Parse(attr.Val)
						if err != nil {
							log.Println(err)
							continue
						}
						resolved := p.URL.ResolveReference(link)
						p.Static[resolved.String()] = struct{}{}
					}
				}
			}
		}
	}
}

func (c *Crawler) tryCrawl(link *url.URL) {
	linkStr := link.String()

	c.Lock()
	if ok := c.visited[linkStr]; !ok && c.pagesVisited < c.pageLimit {
		c.visited[linkStr] = true
		c.pagesVisited++
		c.Add(1)
		go c.Crawl(linkStr)
	}
	c.Unlock()
}
