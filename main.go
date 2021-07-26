package main

import (
	"fmt"
	"github.com/Robbie08/webcrawler/pkg/crawler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Println("Server is running!")
	crawler.Dummy()
	http.HandleFunc("/shutdown", shutDownServer)
	http.HandleFunc("/run", runWebcrawler)
	http.HandleFunc("/", startServer)
	http.ListenAndServe(":8080", nil)
}

func runWebcrawler(res http.ResponseWriter, req *http.Request) {
	runCrawler() // start the webcrawler
}

// This function is in charge of gracefully shutting down the server
func shutDownServer(res http.ResponseWriter, req *http.Request) {
	log.Println("Server shutting down...")
	os.Exit(0)
}

// This function is in charge of handling requests to the homepage
func startServer(res http.ResponseWriter, req *http.Request) {
	log.Println("Someone hit the homepage")
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "You are home")
}

// This function will scrape the URLs from the website passed in as command line arguments
func runCrawler() {
	printBanner()
	fmt.Println("\nStarting up crawler ...")
	time.Sleep(3 * time.Second)

	urlsFound := make(map[string]bool)
	seedUrls := os.Args[1:]

	// Define and init our channels that we will use to communicate
	msg := make(chan string)
	done := make(chan bool)

	fmt.Println("\n----------------- Scraping from these URLs ------------------")
	// Spin up the crawler as a go routine
	for _, url := range seedUrls {
		fmt.Println("\nURL -> " + url)
		go crawler.Crawl(url, msg, done)
	}

	// Need to subcribe to both the msg and done channel to enable communication
	for stop := 0; stop < len(seedUrls); {
		select {
		case url := <-msg:
			urlsFound[url] = true
		case <-done:
			stop++
		}
	}

	fmt.Println("\n****************************************")
	// Print out the scraped results here
	fmt.Println("*      Crawler found [", len(urlsFound), "] urls      *")
	fmt.Println("****************************************")

	for url, _ := range urlsFound {
		fmt.Println(" * " + url)
	}
	close(msg) // close out the msg chanel
}

func printBanner() {
	fmt.Println(" _    ___       __           __     ______                    __         ")
	fmt.Println("| |  / (_)___  / /__  ____  / /_   / ____/________ __      __/ /__  _____")
	fmt.Println("| | / / / __ \\/ / _ \\/ __ \\/ __/  / /   / ___/ __ `/ | /| / / / _ \\/ ___/")
	fmt.Println("| |/ / / /_/ / /  __/ / / / /_   / /___/ /  / /_/ /| |/ |/ / /  __/ /    ")
	fmt.Println("|___/_/\\____/_/\\___/_/ /_/\\__/   \\____/_/   \\__,_/ |__/|__/_/\\___/_/     ")
}
