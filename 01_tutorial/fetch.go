// Prints the content found at a URL
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
  for _, url := range os.Args[1:] {
    // fetch the url data and print the error if it exists
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }

    // parse only the body and check for errors once again
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", b)
  }
}
