package main

import (
  "log"
  "time"

  "github.com/vejjux/countly"
)

func main() {
  appKey := "AppTeskKeyGetFromCountly"

  cs, err := countly.CreateSession("https://example-countly-url.com/i", appKey, "1.0", 60, true)
  if err != nil {
    log.Println(err.Error())
    return
  }

  time.Sleep(10 * time.Second)
  _ = cs.Event("install")
  time.Sleep(190 * time.Second)
  _ = cs.Event("uninstall")
  time.Sleep(10 * time.Second)

  csEnd := cs.End()
  <- csEnd
}
