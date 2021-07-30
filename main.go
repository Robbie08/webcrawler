package main

import (
	"encoding/json"
	"fmt"
	"github.com/Robbie08/webcrawler/pkg/crawler"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{}) // making our logs into JSON format
	log.SetOutput(ioutil.Discard)          // COMMENT THIS LINE OUT IF YOU WANT TO VIEW THE LOGS also "io/ioutil"
	crawler.Dummy()
	http.HandleFunc("/shutdown", shutDownServer)
	http.HandleFunc("/run", runWebcrawler)
	http.HandleFunc("/", startServer)
	http.ListenAndServe(":8080", nil)
}

func runWebcrawler(res http.ResponseWriter, req *http.Request) {
	printBanner()
	handleIP(res, req)
	runCrawler() // start the webcrawler
}

// This function is in charge of gracefully shutting down the server
func shutDownServer(res http.ResponseWriter, req *http.Request) {
	log.Info("Server shutting down gracefully...")
	fmt.Println("Server shutting down gracefully...")
	os.Exit(0)
}

// This function is in charge of handling requests to the homepage
func startServer(res http.ResponseWriter, req *http.Request) {
	log.Info("Someone hit the homepage")
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		log.Info("Page not found")
		return
	}
	fmt.Fprint(res, "You are home")
}

// This function will scrape the URLs from the website passed in as command line arguments
func runCrawler() {
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

func handleIP(res http.ResponseWriter, r *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	// get and verify if we got an ip
	ip := grabIP(r)

	if ip == "" {
		log.Warn("IP address not found")
		fmt.Println("No IP found ...")
		return
	}

	resp, _ := json.Marshal(map[string]string{
		"ip": ip,
	})
	fmt.Println("\nYour IP :) -> ", ip)
	res.Write(resp)
}

// Function that grabs the requests IP Address by reading from the Header
func grabIP(r *http.Request) string {
	// get the ip from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")

	if ip != "" {
		log.WithFields(
			log.Fields{
				"IP":   ip,
				"Type": "X-REAL-IP",
			},
		).Info("IP Address fetched")
		return ip
	}

	// if we could not get the X-REAL-IP then get the "forwarded-for" ip
	ip = r.Header.Get("X-FORWARDED-FOR")

	if ip != "" {
		log.WithFields(
			log.Fields{
				"IP":   ip,
				"Type": "X-FORWARDED-FOR-IP",
			},
		).Info("IP Address fetched")
		return ip
	}

	// if we couldn't fetch any ip by now, then we can just assume local host [for now]
	ip = r.RemoteAddr
	log.WithFields(
		log.Fields{
			"IP":   ip,
			"Type": "REMOTE-IP",
		},
	).Info("IP Address fetched")
	return ip
}
