package main

import (
	"fmt"
	"github.com/Robbie08/webcrawler/pkg/crawler"
	"log"
	"net/http"
	"os"
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
	log.Println("The Crawler is live!")
	fmt.Println("The Crawler is live!")
	runCrawler()

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
	fmt.Fprint(res, "We are in the homepage")
}

func runCrawler() {
	fmt.Println("Inside Crawler")
}
