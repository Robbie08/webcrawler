package main

import (
    "net/http"
    "os"
    "fmt"
    )

func main(){
    http.HandleFunc("/shutdown", shutDownServer)
    http.HandleFunc("/", startServer)
    http.ListenAndServe(":8080", nil)
}

// This function is in charge of gracefully shutting down the server
func shutDownServer(res http.ResponseWriter, req *http.Request){
    os.Exit(0)
}

// This function is in charge of handling requests to the homepage
func startServer(res http.ResponseWriter, req *http.Request){
    if req.URL.Path != "/" {
        http.NotFound(res, req)
        return
    }
    fmt.Fprint(res, "We are in the homepage")
}
