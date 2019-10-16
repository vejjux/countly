package lib

import "log"

var logDebug = false

func Log (args ...interface{}) {
  if logDebug {
    log.Println(args)
  }
}

func LogSetDebug (on bool) {
  logDebug = on
}
