package main

import (
  "flag"
  "fmt"
)

func main() {
  flag.Parse()
  args := flag.Args()
  formatted := FmtName(args[0])
  fmt.Printf("%s", formatted)
}
