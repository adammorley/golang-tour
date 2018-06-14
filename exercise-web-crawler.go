package main

import (
    "fmt"
    "sync"
)

type UrlTracker struct {
    v   map[string]bool
    mux sync.Mutex
}

// returns false if already present
func (t *UrlTracker) add(s string) bool {
    t.mux.Lock()
    defer t.mux.Unlock()
    if t.v[s] == true {
        return false
    }
    t.v[s] = true
    return true
}

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, t UrlTracker) {
    // fetch in parallel, don't re-fetch
    // use go routine       use map w/ mutex
    if depth <= 0 {
        return
    }
    if t.add(url) {
        body, urls, err := fetcher.Fetch(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("found: %s %q\n", url, body)
        for _, u := range urls {
            Crawl(u, depth-1, fetcher, t)
        }
    }
    return
}

func main() {
    var t UrlTracker
    t.v = make(map[string]bool)
    Crawl("https://golang.org/", 4, fetcher, t)
}

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
