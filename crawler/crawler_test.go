package ricrawler

import (
	"log"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawler_extractLinks(t *testing.T) {
	assert := assert.New(t)

	type input struct {
		URL  string
		data string
	}

	tests := []struct {
		name    string
		input   *input
		want    *Page
		wantErr bool
	}{
		{
			name: "Generic",
			input: &input{
				URL: "https://www.test.com",
				data: `<!DOCTYPE html>
				<html>
				  <title>Website Title</title>
				<h3><a href="/doc/install">Getting Started</a></h3>
				<h3 id="code"><a href="code.html">How to write Go code</a></h3>
				Also available as a <a href="//www.youtube.com/watch?v=XCsL89YtqCs">screencast</a>, this
				<li><a href="/doc/codewalk/functions">First-Class Functions in Go</a></li>
				<img class="gopher" src="/doc/gopher/talks.png"/>
				<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
				</html>`,
			},
			want: &Page{
				URL:   nil,
				Title: "Website Title",
				Interal: map[string]struct{}{
					"https://www.test.com/doc/install":            struct{}{},
					"https://www.test.com/code.html":              struct{}{},
					"https://www.test.com/doc/codewalk/functions": struct{}{},
				},
				External: map[string]struct{}{
					"https://www.youtube.com/watch?v=XCsL89YtqCs":     struct{}{},
					"http://www.google.com/intl/en/policies/privacy/": struct{}{},
				},
				Static: map[string]struct{}{
					"": struct{}{},
				},
			},
			wantErr: false,
		},
		// TODO add more cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCrawler(1, false)

			pageURL, err := url.Parse(tt.input.URL)
			if err != nil {
				log.Println(err)
			}
			data := strings.NewReader(tt.input.data)

			gotPage, err := c.extractLinks(pageURL, data)
			assert.Equal(tt.want.Title, gotPage.Title, "Title should match")
			assert.Equal(tt.want.Interal, gotPage.Interal, "Internal Links should match")
			assert.Equal(tt.want.External, gotPage.External, "External Links should match")
			// assert.Equal(tt.want.Static, gotPage.Static, "Static Links should match")
		})
	}
}
