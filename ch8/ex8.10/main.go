package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

//cancel
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// Extract makes an HTTP GET request to the specified URL, parses the response as HTML, and returns the links in the HTML document
func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)

	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//tokens is a counting semaphore used to
//enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string,  *sync.WaitGroup) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	//Cancel traversal when input is detected
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	worklist := make(chan []string)
	var n sync.WaitGroup
	go func() {
		n.Add(1)
		worklist <- os.Args[1:]
	}()
	go func() {
		n.Wait()
		close(worklist)
	}()

	//Crawl the web concurrently
	seen := make(map[string]bool)
loop:
	for {
		select {
		case <-done:
			for range worklist {
				n.Done()
			}
			fmt.Println("canel all request")
		case list, ok := <-worklist:
			if !ok {
				break loop
			}
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}

}
