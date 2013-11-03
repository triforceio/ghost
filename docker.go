package main

import (
  "fmt"
  "net"
  "net/http"
  "net/http/httputil"
  "io"
  "io/ioutil"
)

type Client struct {
  addr string
  APIVersion string
}

func (c Client) BuildImage(contents io.Reader) error {
  req, err := http.NewRequest("POST", fmt.Sprintf("%s", "/build"), contents)

  if err != nil {
    fmt.Printf("Error creating request: %s", err)
    return nil
  }

  req.Host = c.addr
  req.Header.Set("Content-Type", "application/tar")

  conn,err := net.Dial("tcp", c.addr)

  if err != nil {
    fmt.Printf("Error dialing Docker API: %s", err)
    return nil
  }

  clientconn := httputil.NewClientConn(conn, nil)
  resp, err := clientconn.Do(req)
  defer clientconn.Close()

  if err != nil {
    fmt.Printf("Got error connecting to Docker: %s", err)
    return nil
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  if resp.StatusCode < 200 || resp.StatusCode >= 400 {
    if len(body) == 0 {
      return fmt.Errorf("Error: %s", http.StatusText(resp.StatusCode))
    }
    return fmt.Errorf("Error: %s", body)
  }

  fmt.Println(string(body))

  return nil
}

/* func main() { */
/*   client := Client{"172.17.42.1:4243", "1.6"} */
/*   client.Request("/images") */
/* } */
