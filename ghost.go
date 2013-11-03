package main

import (
  "flag"
  "fmt"
  "strings"
  "bytes"
  "archive/tar"
  "io"
)

type Archive io.Reader

func PackageFile(content string) (Archive, error) {
  buf := new(bytes.Buffer)
  tw := tar.NewWriter(buf)

  hdr := &tar.Header{
    Name: "Dockerfile",
    Size: int64(len(content)),
    Typeflag: tar.TypeReg,
  }
  if err := tw.WriteHeader(hdr); err != nil {
    return nil, err
  }
  if _, err := tw.Write([]byte(content)); err != nil {
    return nil, err
  }

  if err := tw.Close(); err != nil {
    return nil, err
  }
  return buf, nil
}

func main() {
  flag.Parse()
  args := flag.Args()
  packages := strings.Split(args[0], ",")
  formatted := FmtName(args[0])
  dockerfile := new(Dockerfile)
  fmt.Println("SHA: %s", formatted)
  MakeDockerfile(packages, dockerfile)
  archive, err := PackageFile(dockerfile.Contents())

  if err != nil {
    fmt.Println("Error compressing file: ", err);
  }

  fmt.Println("writing Dockerfile to S3 bucket...")
  url, err := WriteDockerfile(archive, formatted)

  if url != "" {
    fmt.Println(url)
  } else {
    fmt.Println("Error writing Dockerfile: ", err)
  }

  client := Client{"172.17.42.1:4243", "1.6"}
  client.BuildImage(archive)
}
