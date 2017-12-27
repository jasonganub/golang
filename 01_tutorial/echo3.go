// An insane oneliner command-line printer
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
