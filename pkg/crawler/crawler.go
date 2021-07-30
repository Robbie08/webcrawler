package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func Dummy() {
	fmt.Println("Hello Friend!")
}

// Find and return the href value from the given token
func scrapeHref(tkn html.Token) (ok bool, href string) {
	// we will iterate through attributes only returning the href value
	for _, attr := range tkn.Attr {
		if attr.Key == "href" {
			href = attr.Val
			ok = true
		}
	}
	return // we will make a naked return
}

// function in charge of crawling and scraping a given url
func Crawl(url string, msg chan string, done chan bool) {
	resp, err := http.Get(url) // make a Get request and returns the response from url

	defer func() {
		done <- true // we publish to the done chanel that we are done crawling
	}()

	if err != nil {
		log.Println("Oops! Error when getting response from: ", url)
		return
	}

	myBod := resp.Body  // get the body of the response
	defer myBod.Close() // make sure to close out body when we are done operating on it

	tokenizedBody := html.NewTokenizer(myBod) // tokenize the HTML code into TAGs

	// now we must iterate through the tokenized data and keep track of the <a> TAGs
	for {
		tk := tokenizedBody.Next() // move to the next item

		switch {
		case tk == html.ErrorToken:
			return // finish here if there is an error with our data
		case tk == html.StartTagToken:
			tkn := tokenizedBody.Token() // grab token

			// let's only keep the <a> Tokens

			if !(tkn.Data == "a") {
				continue // if the token we currently have is not "a" mv to next token
			}

			ok, fetchedUrl := scrapeHref(tkn) // we will use our created function to scrape the href

			// Guard againts an <a> tag that doesn't actually have an href inside
			if !ok {
				continue
			}

			hasProto := strings.Index(fetchedUrl, "http") == 0 // store as bool if the url contains http
			if hasProto {
				msg <- fetchedUrl // send the fetched url to the msg channel
			}
		}
	}
}
