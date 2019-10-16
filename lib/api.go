package lib

import (
  "net/url"
  "runtime"
  "strconv"
  "time"
)

type ApiMap struct {
  base url.Values
}

func CreateApiMap (appKey string, sdkVersion string) (apiMap ApiMap, err error) {
  apiMap = ApiMap{
    base: url.Values{},
  }

  deviceId, err := GetDeviceId()
  if err != nil { return }

  apiMap.base.Add("app_key", appKey)
  apiMap.base.Add("device_id", deviceId)
  apiMap.base.Add("sdk_name", "golang-native-" + runtime.GOOS)
  apiMap.base.Add("sdk_version", sdkVersion)

  return
}

func (apiMap ApiMap) MakeValues () url.Values {
  values := url.Values{}

  for key, val := range apiMap.base {
    values[key] = val
  }

  _, tzOffset := time.Now().Zone()
  tzOffset = tzOffset / 60

  values.Add("dow", strconv.Itoa(int(time.Now().Weekday())))
  values.Add("tz", strconv.Itoa(tzOffset))
  values.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
  values.Add("hour", strconv.Itoa(time.Now().Hour()))

  return values
}
