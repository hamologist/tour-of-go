package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type safeCache struct {
	mu sync.Mutex
	v  map[string]struct{}
}

func (c *safeCache) add(key string) {
	c.mu.Lock()
	c.v[key] = struct{}{}
	c.mu.Unlock()
}

func (c *safeCache) check(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, in := c.v[key]

	return in
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// ---TODO: Fetch URLs in parallel.---
	// ---TODO: Don't fetch the same URL twice.---
	// This implementation doesn't do either:
	var wg sync.WaitGroup
	cache := safeCache{v: make(map[string]struct{})}

	var inner func(url string, depth int)
	inner = func(url string, depth int) {
		if depth <= 0 {
			return
		}
		if cache.check(url) {
			return
		}
		cache.add(url)

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				inner(u, depth-1)
			}(u)
		}
	}
	inner(url, depth)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}
