package main

import (
  "launchpad.net/goamz/aws"
  "launchpad.net/goamz/s3"
  "mime"
  "io"
  "bytes"
)

func WriteDockerfile(contents io.Reader, name string) (string, error) {
  mime.AddExtensionType(".tar", "application/x-tar")
  var bytes bytes.Buffer
  auth, err := aws.EnvAuth()

  if err != nil {
    panic(err)
  }

  client := s3.New(auth, aws.USEast)
  bucket := client.Bucket("images.static.triforce.io")
  filename := name + ".tar"
  contentType := mime.TypeByExtension(".tar")
  bytes.ReadFrom(contents)
  putErr := bucket.Put(filename, bytes.Bytes(), contentType, s3.PublicRead)

  if putErr != nil {
    return "", putErr
  }

  return bucket.URL(filename), nil
}
