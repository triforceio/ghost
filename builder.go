package main

import (
  "bytes"
  "fmt"
  "text/template"
  "io/ioutil"
)

type FileGenerator interface {
  ReadPrelude() string
  ReadEpilogue() string
  ReadBase() string
  ReadExtras(packageName string) string
  Write(template string) (ret int, err error)
}

type Dockerfile struct {
  bytes []byte
}

func ParsedSnippet(path string) string {
  text := new(bytes.Buffer)
  t := template.New("prelude")
  contents,err := ioutil.ReadFile("templates/prelude.tmpl")
  if err != nil {
    fmt.Println(err)
  }
  t = template.Must(t.Parse(string(contents)))
  t.Execute(text, "")
  return text.String()
}

func (d Dockerfile) ReadPrelude() string {
  return ParsedSnippet("templates/prelude.tmpl")
}

func (d Dockerfile) ReadEpilogue() string {
  return ""
}

func (d Dockerfile) ReadExtras(packageName string) string {
  return ""
}

func (d Dockerfile) ReadBase() string {
  return ""
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
  final.WriteString(fmt.Sprintln(g.ReadEpilogue()))
  g.Write(final.String())
}

func (d Dockerfile) Contents() string {
  return string(d.bytes[:len(d.bytes)])
}
