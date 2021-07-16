package main

import (
    "net/http"
    "log"
    "os"
    "io/ioutil"
    "fmt"
    )

func main(){
    http.HandleFunc("/shutdown", shutDownServer)
    http.HandleFunc("/", startServer)
    http.HandleFunc("/activate", startWebCrawler)
    http.ListenAndServe(":8080", nil)
}

// This function is in charge of gracefully shutting down the server
func shutDownServer(res http.ResponseWriter, req *http.Request){
    log.Println("Shutting Down Server")
    os.Exit(0)
}


// This function is in charge of handling requests to the homepage
func startServer(res http.ResponseWriter, req *http.Request){
    if req.URL.Path != "/" {
        http.NotFound(res, req)
        return
    }
    log.Println("Something hit the homepage!")
    fmt.Fprint(res, "We are in the homepage")
}

func startWebCrawler(res http.ResponseWriter, req *http.Request) {
    // This function will allow us to curl an http address
    targetURL := "https://finance.yahoo.com/quote/NVDA?p=NVDA&.tsrc=fin-srch"

    log.Print(targetURL)
    resp, err := http.Get(targetURL)

    // check if the request was successful or not
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close();

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}
