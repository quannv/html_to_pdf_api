package main

import (
    "encoding/json"
//     "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "bytes"
// "strings"
)

func main() {
    API_KEY := "1234abcdf"

//     content, err := ioutil.ReadFile("test.html")
//     text := string(content)
//     fmt.Println(text)
//     if err != nil {
//         log.Fatal(err)
//     }

//     message := map[string]interface{}{
//             "url":  "https://en.wikipedia.org/wiki/PDF",
//     }
//
//     bytesRepresentation, err := json.Marshal(message)
//     if err != nil {
//         log.Fatalln(err)
//     }

    data := url.Values{}
    data.Set("html", `Lazy Test`)


    client := http.Client{}
    request, err := http.NewRequest("POST", "http://localhost:9000/render", bytes.NewBufferString(data.Encode()))
    if err != nil {
        log.Fatalln(err)
    }
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")


    q := request.URL.Query()
    q.Add("apikey", API_KEY)
    request.URL.RawQuery = q.Encode()

//     request.ParseForm()
//     value := request.FormValue(text)

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
        ioutil.WriteFile("wikipedia.pdf", body, 0644)
    } else {
        // An error occurred
        var result map[string]interface{}

        json.NewDecoder(resp.Body).Decode(&result)

        log.Println(result)
        log.Println(result["data"])
    }
}
