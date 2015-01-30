package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

const (
  MOBILE_DETECT_JSON = "https://raw.githubusercontent.com/serbanghita/Mobile-Detect/master/Mobile_Detect.json"
)

type Detect struct {
  UaMatch struct {
    Phones struct {
      Iphone string `json:"iPhone"`
    } `json:"phones"`
  } `json:"uaMatch"`
}

func getContent(url string) ([]byte, error) {
  // Build request.
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  // Send request via a client.
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  // Defer closing of body.
  defer resp.Body.Close()

  // Read contents into byte array
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  // Return bytes
  return body, nil
}

func main() {
  content, err := getContent(MOBILE_DETECT_JSON)
  if err != nil {
    log.Fatal(err)
  } else {
    var d Detect
    json.Unmarshal(content, &d)
    // Write to mobile_detect.vcl
    f, err := os.Create("mobile_detect.vcl")
    if err != nil {
      log.Fatal(err)
    }
    defer f.Close()

    mbFile, err := f.WriteString(d.UaMatch.Phones.Iphone)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("Wrote mobile_detect.vcl file %d bytes", mbFile)
  }
}
