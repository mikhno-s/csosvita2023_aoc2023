package main

import (
	"fmt"
	"strings"
)

type HtmlParser struct {
	urls map[string][]string
}

func (this *HtmlParser) GetUrls(url string) []string {
	return this.urls[url]
}

func (this *HtmlParser) Build(urls []string, edges [][]int) {
	this.urls = map[string][]string{}
	for i := range urls {
		this.urls[urls[i]] = []string{urls[i]}
		for _, edge := range edges {
			if edge[0] == i {
				this.urls[urls[i]] = append(this.urls[urls[i]], urls[edge[1]])
			}
		}
	}
	return
}

// BFS
func crawl(startUrl string, htmlParser HtmlParser) []string {
	result := []string{}
	visitedURLs := map[string]bool{startUrl: true}

	// queue
	crawlingQ := []string{startUrl}
	result = append(result, startUrl)

	// itterate over queue
	for len(crawlingQ) > 0 {
		// get all adjustments of the current V (first from queue)
		crawledUrls := getSameHostames(startUrl, htmlParser.GetUrls(crawlingQ[0]))
		// fmt.Println(crawledUrls)
		for _, url := range crawledUrls {
			if !visitedURLs[url] {
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

func getSameHostames(origin string, hostnames []string) []string {
	result := make([]string, 0)
	origin, _ = strings.CutPrefix(origin, "http://")
	origin, _ = strings.CutPrefix(origin, "https://")
	origin, _, _ = strings.Cut(origin, "/")
	for i := range hostnames {
		hostname := hostnames[i]
		hostname, _ = strings.CutPrefix(hostname, "http://")
		hostname, _ = strings.CutPrefix(hostname, "https://")
		hostname, _, _ = strings.Cut(hostname, "/")
		if origin == hostname {
			result = append(result, hostnames[i])
		}
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
}
