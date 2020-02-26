package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    API_KEY := "YOUR_API_KEY"

    client := http.Client{}
    request, err := http.NewRequest("GET", "https://api.wepdf.io/render", nil)

    q := request.URL.Query()
    q.Add("apikey", API_KEY)
    q.Add("url", "https://en.wikipedia.org/wiki/PDF")
    request.URL.RawQuery = q.Encode()
    if err != nil {
        log.Fatalln(err)
    }

    resp, err := client.Do(request)
    if err != nil {
        log.Fatalln(err)
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln(err)
        }
        // write the response to file
        ioutil.WriteFile("wepdf.pdf", body, 0644)
    } else {
        // An error occurred
        var result map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&result)
        log.Println(result)
        log.Println(result["data"])
    }
}
