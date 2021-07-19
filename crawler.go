package crawler

import (
    "fmt"
    "net.http"
    "net/url"
)



// Grab package using Get request. 
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
