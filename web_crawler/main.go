package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/net/html"
)

var (
	urlsProcessed = sync.Map{}
)

// recursive web-crawler
// sample usage: go run main.go https://xkcd.com comics/dest
func main() {
	// args: starting url and dest dir
	if len(os.Args) > 3 {
		log.Fatalf("usage: invalid number of arguments %v: required 2", len(os.Args))
	}

	validURL, err := url.Parse(os.Args[1])
	if err != nil || validURL.Scheme == "" || validURL.Host == "" {
		log.Fatal("invalid url entered")
	}

	terminateSignal := make(chan os.Signal, 1)
	signal.Notify(terminateSignal, syscall.SIGINT)
	go func() {
		<-terminateSignal
		log.Fatal("CTRL+C entered, quitting")
	}()

	if err := fetchAndDownloadPageContent(validURL); err != nil {
		log.Printf("error occurred: %v", err)
	}
	log.Println("successfully downloaded all files")
}

func downloadURL(validURL *url.URL) {
	// wait for 1 minute for response
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, os.Args[1], nil)
	client := http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	// have a re-try of 5 times
	if err != nil {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 5)
			resp, err = client.Do(req)
			if err == nil {
				break
			}
		}
	}
	// create dest dir if it doesn't exist
	dirPath := filepath.Join(".", os.Args[2])
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating destination directory %v\n", err)
	}
	if err != nil {
		log.Printf("error downloading %v\n", validURL.Path)
	}
	newFilePath := filepath.Join(dirPath, strings.TrimPrefix(validURL.Path, "/"))
	file, err := os.Create(newFilePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	io.Copy(file, resp.Body)
}

func fetchAndDownloadPageContent(validURL *url.URL) error {
	downloadURL(validURL)
	//mark it as visited
	urlsProcessed.Store(validURL.String(), struct{}{})
	// traverse this URL's embedded links
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, os.Args[1], nil)
	client := http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	urlsToTraverse := traverseURLs(validURL, resp.Body)
	wg := new(sync.WaitGroup)
	for _, nextLink := range urlsToTraverse {
		absoluteLink := validURL.ResolveReference(&nextLink)
		if _, ok := urlsProcessed.Load(absoluteLink); !ok && absoluteLink != nil {
			log.Printf("downloading %v", absoluteLink.String())
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := fetchAndDownloadPageContent(absoluteLink); err != nil {
					log.Printf("error fetching page %v", absoluteLink.String())
				}
			}()
		}
	}
	wg.Wait()
	return err
}

func traverseURLs(base *url.URL, resp io.Reader) []url.URL {
	token := html.NewTokenizer(resp)
	urls := make([]url.URL, 0)

	for {
		tokenType := token.Next()
		if tokenType == html.ErrorToken {
			return urls
		}
		if tokenType == html.StartTagToken {
			token := token.Token()
			// <a> <href> is the hyperlink we want to follow
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link, err := url.Parse(attr.Val)
						if err == nil && link.Host == "" {
							urls = append(urls, *base.ResolveReference(link))
						}
					}
				}
			}
		}
	}
}
