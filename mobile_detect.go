package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

const (
  MOBILE_DETECT_JSON = "https://raw.githubusercontent.com/serbanghita/Mobile-Detect/master/Mobile_Detect.json"
)

func main() {
  // Grab json file from github.
  resp, err := http.Get(MOBILE_DETECT_JSON)
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  } else {
    defer resp.Body.Close()
    // Grab contents of json file.
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
    }
    fmt.Printf("%s\n", string(contents))
  }
}


