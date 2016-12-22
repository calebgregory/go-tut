package main

import (
	"fmt"
	"sync"
)

type VisitationHistory struct {
	Urls map[string]bool
	M    sync.Mutex
}

func (v *VisitationHistory) HasBeenVisited(url string) bool {
	v.M.Lock()
	defer v.M.Unlock()

	if v.Urls[url] {
		return true
	} else {
		return false
	}
}

func (v *VisitationHistory) MarkAsVisited(url string) {
	v.M.Lock()
	defer v.M.Unlock()

	v.Urls[url] = true
}

var v = VisitationHistory{make(map[string]bool), sync.Mutex{}}
var wg sync.WaitGroup

func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Wait()
	defer wg.Done()

	if depth == 0 {
		return
	}

	if !v.HasBeenVisited(url) {
		body, urls, err := fetcher.Fetch(url)
		v.MarkAsVisited(url)

		if err != nil {
			fmt.Printf("°n° - %v\n", err)
			return
		}

		fmt.Printf("'u' - %v : %v\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher)
		}
	}

}

func main() {
	wg.Add(1)
	Crawl("http://golang.org/", 4, fetcher)
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %v", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
