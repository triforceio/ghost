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
  fmt.Println("writing Dockerfile to S3 bucket...")
  url, err := WriteDockerfile([]byte(dockerfile.Contents()), formatted)

  if url != "" {
    fmt.Println(url)
  } else {
    fmt.Println("Error writing Dockerfile: ", err)
  }
}
