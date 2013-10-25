package main

import (
  "flag"
  "fmt"
  "strings"
)

func main() {
  flag.Parse()
  args := flag.Args()
  packages := strings.Split(args[0], ",")
  formatted := FmtName(args[0])
  dockerfile := new(Dockerfile)
  fmt.Println("SHA: %s", formatted)
  MakeDockerfile(packages, dockerfile)
  fmt.Println("Contents: ", dockerfile.Contents())
}
