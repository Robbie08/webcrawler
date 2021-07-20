package crawler

import (
    "fmt"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)


// Basic layout of object that will
// contain our scraped data
type ScrapedData {
    URL     string
    Title   string
    H1      string
}


// let's create an interface to help multi-file data transfer
type Scraper interface {
    ScrapePage(*goquery.Document) ScrapedData
}

// Dummy function to help test if we are still connected to main.go
func PrintFromCrawler(val string){
    fmt.Println("Hello ", val)
}

// Grab package using Get request.
// This is the start of the scraping process and will help us
// gain access to the HTML code from the site.
func getRequest(url string) (*http.Response, error){
    target := &http.Client{}
    req, _ := http.NewRequest("Get", url, nil)

    // we need to wrap our request with a googlebot agent 
    // to avoid being detected by client.
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://.google.com/bot.html)")

    res, err := target.Do(req) // perform the request on target 

    if err != nil {
        return nil, err
    }

    return res, nil
}
