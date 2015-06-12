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

var log = map[string]string{}
var logmu sync.Mutex

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	var wg sync.WaitGroup

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	logmu.Lock()
	log[url] = body
	logmu.Unlock()

	for _, u := range urls {
		wg.Add(1)
		go Crawler(u, depth-1, fetcher, &wg)
	}
	wg.Wait()
	return
}

func Crawler(url string, depth int, fetcher Fetcher,
	w *sync.WaitGroup) {
	if depth <= 0 {
		w.Done()
		return
	}

	// previous crawl check
	logmu.Lock()
	_, ok := log[url]
	logmu.Unlock()
	if ok {
		w.Done()
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)

		//add not-found URL to log
		logmu.Lock()
		log[url] = body
		logmu.Unlock()

		w.Done()
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	logmu.Lock()
	log[url] = body
	logmu.Unlock()

	for _, u := range urls {
		w.Add(1)
		go Crawler(u, depth-1, fetcher, w)
	}
	w.Done()
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
