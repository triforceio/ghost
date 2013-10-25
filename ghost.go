package main

import (
  "flag"
  "fmt"
)

func main() {
  flag.Parse()
  args := flag.Args()
  formatted := FmtName(args[0])
  fmt.Println("SHA: %s", formatted)
  dockerfile := new(Dockerfile)
  MakeDockerfile([]string{"blah"}, dockerfile)
  fmt.Println("Contents: ", dockerfile.Contents())
}
