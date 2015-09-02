package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "encoding/json"

)

// Extract all http** links from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {

    resp, err := http.Get(url)

    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()
    if err != nil {
        fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
        return
    }

    b := resp.Body
    defer b.Close() // close Body when the function returns

    z := html.NewTokenizer(b)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()
            // fmt.Println(t.Data)
            ch <- t.Data
        }
    }
}

func getTags(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    
    fmt.Println("GET params:", r.URL.Query())
     website := r.URL.Query().Get("website")
    fmt.Println("website requested: ", website)
    if website == "" {
        return
    }
    foundTags := make(map[string]bool)

    url := website
    // Channels
    chTags := make(chan string)
    chFinished := make(chan bool) 

    go crawl(url, chTags, chFinished)
    
    // Subscribe to both channels
    for c := 0; c < 1; {
        select {
        case url := <-chTags:
            foundTags[url] = true //setting its value to true if a tag is found
        case <-chFinished:
            c++
        }
    }

      js, err := json.Marshal(foundTags)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.Header().Set("Content-Type", "application/json")
      w.Write(js)

    close(chTags)
}
func main() {
    http.HandleFunc("/", getTags)
    http.ListenAndServe(":9000", nil)
}