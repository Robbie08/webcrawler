package main

import (
    "net/http"
    "os"
    "io/ioutil"
    "log"
    "fmt"
    "github.com/Robbie08/webcrawler/pkg/crawler"
    )

func main(){
    log.Println("Server is running!")
    crawler.PrintFromCrawler("Diggity Daug")
    http.HandleFunc("/shutdown", shutDownServer)
    http.HandleFunc("/run", runWebcrawler)
    http.HandleFunc("/", startServer)
    http.ListenAndServe(":8080", nil)
}


func runWebcrawler(res http.ResponseWriter, req *http.Request){
    targetUrl := "https://ortizrobert.herokuapp.com/"
    log.Println("Target: ", targetUrl)

    resp, err := http.Get(targetUrl)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}

// This function is in charge of gracefully shutting down the server
func shutDownServer(res http.ResponseWriter, req *http.Request){
    log.Println("Server shutting down...")
    os.Exit(0)
}

// This function is in charge of handling requests to the homepage
func startServer(res http.ResponseWriter, req *http.Request){
    log.Println("Someone hit the homepage")
    if req.URL.Path != "/" {
        http.NotFound(res, req)
        return
    }
    fmt.Fprint(res, "We are in the homepage")
}
