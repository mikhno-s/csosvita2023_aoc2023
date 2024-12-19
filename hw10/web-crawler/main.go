package main

import (
	"fmt"
	"strings"
)

type HtmlParser struct {
	urls map[string][]string
}

func (p *HtmlParser) GetUrls(url string) []string {
	return p.urls[url]
}

func (p *HtmlParser) Build(urls []string, edges [][]int) {
	p.urls = map[string][]string{}
	for i := range urls {
		p.urls[urls[i]] = []string{urls[i]}
		for _, edge := range edges {
			if edge[0] == i {
				p.urls[urls[i]] = append(p.urls[urls[i]], urls[edge[1]])
			}
		}
	}
}

// BFS
func crawl(startUrl string, htmlParser HtmlParser) []string {
	result := []string{}
	visitedURLs := map[string]bool{startUrl: true}

	// queue
	crawlingQ := []string{startUrl}
	result = append(result, startUrl)

	startHostname := strings.Split(startUrl, "/")[2]

	// itterate over queue
	for len(crawlingQ) > 0 {
		// get all adjustments of the current V (first from queue)
		for _, url := range htmlParser.GetUrls(crawlingQ[0]) {
			hostname := strings.Split(url, "/")[2]
			if hostname == startHostname && !visitedURLs[url] {
				visitedURLs[url] = true
				result = append(result, url)
				// push to the queue
				crawlingQ = append(crawlingQ, url)
			}
		}
		// pop from queue
		crawlingQ = crawlingQ[1:]

	}
	return result
}

func main() {
	htmlParser := &HtmlParser{}
	urls := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
		"http://news.yahoo.com/us",
	}
	edges := [][]int{{2, 0}, {2, 1}, {3, 2}, {3, 1}, {0, 4}}
	htmlParser.Build(urls, edges)
	fmt.Println(crawl("http://news.yahoo.com/news/topics/", *htmlParser))

	urls = []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
	}
	edges = [][]int{{0, 2}, {2, 1}, {3, 2}, {3, 1}, {3, 0}}
	htmlParser.Build(urls, edges)
	fmt.Println(crawl("http://news.google.com", *htmlParser))
}
