package main

import (
  "testing"
  "fmt"
)

type TestDockerfile struct {
  Dockerfile
  Stream string
}

func (d TestDockerfile) ReadPrelude() string {
  return "testing prelude"
}

func (d TestDockerfile) ReadBase() string {
  return "testing base"
}

func (d TestDockerfile) ReadExtras(packageName string) string {
  if(packageName == "test1") {
    return "install test1"
  } else if(packageName == "test2") {
    return "install test2"
  } else {
    return fmt.Sprint("install", packageName)
  }
}

func (d TestDockerfile) ReadEpilogue() string {
  return "testing epilogue"
}

func (d *TestDockerfile) Write(template string) (ret int, err error) {
  d.bytes = []byte(template)
  // in prod code write straight to S3
  return len(d.bytes), nil
}


func ExampleWrite() {
  b := new(TestDockerfile)
  fmt.Println(b.Contents()) // Ensures b.Contents() is empty
  MakeDockerfile([]string{"test1", "test2"}, b)
  fmt.Println(b.Contents())
  // Output:
  // testing prelude
  // testing base
  // install test1
  // install test2
  // testing epilogue
}

func TestGeneratesEpilogue(t *testing.T) {
  t.Skip("Generator includes epilogue template")
}

func TestGeneratesBase(t *testing.T) {
  t.Skip("Includes base image template")
}

func TestGeneratesPackages(t *testing.T) {
  t.Skip("Generates install line for each specified package")
}

func TestGeneratesConfigs(t *testing.T) {
  t.Skip("Adds configs tarball for base & each specified package")
}

func TestGeneratesExtras(t *testing.T) {
  t.Skip("Includes extras template for each specified package")
}
