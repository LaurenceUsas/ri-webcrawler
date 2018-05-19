##  Brief
“We'd like you to write a simple web crawler in Go. The crawler should be limited to one domain - so when crawling it would crawl the domain, but not follow external links, for example to the Facebook and Twitter accounts.

Given a URL, your program should output a site map showing each page's url, title, static assets, internal links and external links.

The number of pages that are crawled should be configurable. We suggest crawling wikipedia and limiting the number of pages to 100.

Ideally, write it as you would a production piece of code. Bonus points for tests and making it as fast as possible!”

##  Approach
1. Scrape website.
2. As soon as internal link found - start scraping it concurently
3. Continue extracting internal/external/static(Tags - "img", "audio", "script", "video", "embed", "source")
4. Once all data is scrapped - Print out result of this page.

The way it is implemented is pretty wild and should be relatively fast but should be further tested to make it more stable.

To further optimise it:
Add workers
Reuse connection
Benchmark data extraction. Consider GoQuery/Regex([Not recommended](https://stackoverflow.com/questions/1732348/regex-match-open-tags-except-xhtml-self-contained-tags/1732454#1732454))


## How to run
```$ make runwiki```  