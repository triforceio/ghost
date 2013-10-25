package main

import (
  "bytes"
  "fmt"
  "text/template"
  "io/ioutil"
  "os"
)

type FileGenerator interface {
  ReadPrelude() string
  ReadEpilogue(packages []string) string
  ReadBase() string
  ReadExtras(packageName string) string
  Write(template string) (ret int, err error)
}

type Dockerfile struct {
  bytes []byte
}

func ParsedSnippet(path string, data interface{}) string {
  text := new(bytes.Buffer)
  t := template.New("prelude")
  contents,err := ioutil.ReadFile(path)
  if err != nil {
    fmt.Println(err)
  }
  t = template.Must(t.Parse(string(contents)))
  t.Execute(text, data)
  return text.String()
}

func FileExists(path string) bool {
  if _, err := os.Stat(path); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}

func (d Dockerfile) ReadPrelude() string {
  return ParsedSnippet("templates/prelude.tmpl", nil)
}

func (d Dockerfile) ReadEpilogue(packages []string) string {
  return ParsedSnippet("templates/epilogue.tmpl", packages)
}

func (d Dockerfile) ReadExtras(packageName string) string {
  path := "templates/extras/" + packageName + ".tmpl"

  if FileExists(path) {
    return ParsedSnippet(path, nil)
  }
  return ""
}

func (d Dockerfile) ReadBase() string {
  return ParsedSnippet("templates/base.tmpl", nil)
}

func (d *Dockerfile) Write(template string) (ret int,err error) {
  d.bytes = []byte(template)
  return len(d.bytes), nil
}

func MakeDockerfile(packages []string, g FileGenerator) {
  var final bytes.Buffer
  final.WriteString(fmt.Sprintln(g.ReadPrelude()))
  final.WriteString(fmt.Sprintln(g.ReadBase()))
  for _, p := range packages {
    line := fmt.Sprintln(g.ReadExtras(p))
    final.WriteString(line)
  }
  final.WriteString(fmt.Sprintln(g.ReadEpilogue(packages)))
  g.Write(final.String())
}

func (d Dockerfile) Contents() string {
  return string(d.bytes[:len(d.bytes)])
}
