package main

import (
  "launchpad.net/goamz/aws"
  "launchpad.net/goamz/s3"
)

func WriteDockerfile(contents []byte, name string) (string, error) {
  auth, err := aws.EnvAuth()

  if err != nil {
    panic(err)
  }

  client := s3.New(auth, aws.USEast)
  bucket := client.Bucket("images.static.triforce.io")
  filename := "Dockerfile-" + name
  putErr := bucket.Put(filename, contents, "text/plain", s3.PublicRead)

  if putErr != nil {
    return "", putErr
  }

  return bucket.URL(filename), nil
}
