package lib

import "runtime"

var platformRef = map[string]string{
  "android": "Android",
  "darwin":  "macOS",
  "freebsd": "Unix",
  "js":      "Web",
  "linux":   "Linux",
  "nacl":    "Nacl",
  "netbsd":  "Unix",
  "openbsd": "Unix",
  "solaris": "Unix",
  "windows": "Windows",
}

func GetPlatform () string {
  if platform, ok := platformRef[runtime.GOOS]; ok {
    return platform
  }

  return "Unknown"
}
