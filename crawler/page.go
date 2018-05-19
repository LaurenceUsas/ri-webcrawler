package ricrawler

import (
	"fmt"
	"net/url"
)

type Page struct {
	URL      *url.URL
	Title    string
	Interal  map[string]struct{} // String instead of *url.URL to resolve duplicates.
	External map[string]struct{}
	Static   map[string]struct{}
}

func NewPage(pageURL *url.URL) *Page {
	p := &Page{
		URL:      pageURL,
		Interal:  make(map[string]struct{}, 0),
		External: make(map[string]struct{}, 0),
		Static:   make(map[string]struct{}, 0),
	}
	return p
}

func (p *Page) Print() {
	fmt.Println("========= START =========")
	fmt.Printf("URL: %s\n", p.URL.String())
	fmt.Printf("Title: %s\n", p.Title)
	fmt.Printf("Internal count: %d\n", len(p.Interal))
	for k, _ := range p.Interal {
		fmt.Println(k)
	}
	fmt.Printf("External count: %d\n", len(p.External))
	for k, _ := range p.External {
		fmt.Println(k)
	}
	fmt.Printf("Static count: %d\n", len(p.Static))
	for k, _ := range p.Static {
		fmt.Println(k)
	}
}
