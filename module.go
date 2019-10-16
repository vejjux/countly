package countly

import (
  "encoding/json"
  "net/url"
  "strconv"
  "time"

  "countly/lib"
)

type Countly struct {
  url      string
  apiMap   lib.ApiMap
  updated  int64
  dataChan chan url.Values
  endChan  chan struct{}
  doneChan chan struct{}
}

func CreateSession (urlStr string, appKey string, appVersion string, heartBeat int, debug bool) (c Countly, err error) {
  lib.LogSetDebug(debug)

  _, err = url.Parse(urlStr)
  if err != nil { return }

  apiMap, err := lib.CreateApiMap(appKey, version)
  if err != nil { return }

  c = Countly{
    url:      urlStr,
    apiMap:   apiMap,
    updated:  time.Now().Unix(),
    dataChan: make(chan url.Values),
    endChan:  make(chan struct{}),
    doneChan: make(chan struct{}),
  }

  metrics := lib.GetMetrics(appVersion)
  metrictsStr, err := metrics.JSON()
  if err != nil {return}

  lib.Log("Starting Countly session", urlStr, appKey)

  go c.consumeData()

  query := c.apiMap.MakeValues()
  query.Add("begin_session", "1")
  query.Add("metrics", metrictsStr)
  c.dataChan <- query

  if heartBeat > 0 {
    go c.heartBeat(heartBeat)
  }

  return
}

func (c *Countly) Event (key string) (err error){
  lib.Log("Sending Countly event", key)
  query := c.apiMap.MakeValues()

  type Event struct {
    Key string `json:"key"`
    Count int `json:"count"`
  }
  events := []Event{Event{Key: key, Count: 1}}
  bin, err := json.Marshal(events)
  if err != nil { return }
  eventsStr := string(bin)

  query.Add("events", eventsStr)
  c.dataChan <- query
  return
}

func (c *Countly) End () chan struct{} {
  lib.Log("Closing Countly session")
  query := c.apiMap.MakeValues()
  query.Add("end_session", "1")
  query.Add("session_duration", strconv.FormatInt(time.Now().Unix() - c.updated ,10))
  c.dataChan <- query
  close(c.endChan)
  return c.doneChan
}

func (c *Countly) heartBeat (interval int) {
  timer := time.NewTicker(time.Duration(interval) * time.Second)

  for {
    select {
    case <- c.endChan:
      timer.Stop()
      lib.Log("Heart-beat done")
      return

    case <- timer.C:
      lib.Log("Sending Countly heart-beat")
      query := c.apiMap.MakeValues()
      query.Add("session_duration", strconv.FormatInt(time.Now().Unix() - c.updated ,10))
      c.dataChan <- query
    }
  }
}

func (c *Countly) consumeData () {
  for {
    select {
    case <- c.endChan:
      close(c.doneChan)
      lib.Log("All done")
      return

    case d := <- c.dataChan:
      body, err := lib.MakeRequest(c.url, d)

      if err != nil {
        lib.Log("Req", c.url, err.Error())
        continue
      }

      lib.Log("Req", c.url, d.Encode(), body)
    }
  }
}


