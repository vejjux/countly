package lib

import (
  "encoding/json"
  "runtime"

  "github.com/denisbrodbeck/machineid"
)

type Metrics struct {
  AppVersion string `json:"_app_version"`
  OS 		     string `json:"_os"`
  OSVersion  string `json:"_os_version"`
}

func (m *Metrics) JSON () (jsonStr string, err error) {
  bin, err := json.Marshal(m)
  if err != nil { return }
  jsonStr = string(bin)
  return
}

func GetMetrics (appVersion string) Metrics {
  return Metrics{
    AppVersion: appVersion,
    OS:         GetPlatform(),
    OSVersion:  runtime.GOARCH,
  }
}

func GetDeviceId () (id string, err error) {
   return machineid.ProtectedID("CountLy")
}
