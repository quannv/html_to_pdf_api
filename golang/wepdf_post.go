package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "bytes"
)

func main() {
    API_KEY := "YOUR_API_KEY"

    content, err := ioutil.ReadFile("YOUR_FILE.html")
    text := string(content)

    var jsonStr = []byte(text)

    client := http.Client{}
    request, err := http.NewRequest("POST", "https://api.wepdf.io/render", bytes.NewBuffer(jsonStr))
    if err != nil {
        log.Fatalln(err)
    }
    request.Header.Set("Content-Type", "text/html")

    // add query parameters
    query := request.URL.Query()
    query.Add("apikey", API_KEY)
    query.Add("landscape", "true")
    request.URL.RawQuery = query.Encode()

    response, err := client.Do(request)
    if err != nil {
        log.Fatalln(err)
    }

    if response.StatusCode >= 200 && response.StatusCode < 300 {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalln(err)
        }
        // write the response to file
        ioutil.WriteFile("wepdf.pdf", body, 0644)
    } else {
        // An error occurred
        var result map[string]interface{}
        json.NewDecoder(response.Body).Decode(&result)

        log.Println(result)
        log.Println(result["data"])
    }
}
