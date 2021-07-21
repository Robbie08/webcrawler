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

    // this code will test our current progrss!
    resp, _ := getRequest("https://ortizrobert.herokuapp.com/")
    doc, _ := goquery.NewDocumentFromResponse(resp)
    links := scrapeLinks(doc)
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

func scrapeLinks(doc *goquery.Document) []string {
    urlsFound := []string{} // this array will keep the urls we have found

    // If we have a non nil document then we can begin searching through it
    if doc != nil {
        doc.Find("a").Each(func(i int, s *goquery.Selection){
            res, _ := s.Attr("href")
            urlsFound = append(urlsFound, res)
        })
    }

    // if we don't find any urls, let us log that
    if len(urlsFound) == 0 {
        log.Println("No URLs found on this webpage")
        return
    }else{ // this statement is for testing purposes
        for _, elem := range(urlsFound) {
            fmt.Println(elem)
        }
    }

    return urlsFound
}

