package lib

import (
  "io/ioutil"
  "net/http"
  "net/url"
)

func MakeRequest(u string, d url.Values) (body string, err error) {
  resp, err := http.PostForm(u, d)
  if err != nil { return }

  raw, err := ioutil.ReadAll(resp.Body)
  if err != nil { return }

  body = string(raw)
  return
}
