package main

import (
  "fmt"
  "io/ioutil"

  "github.com/v010maaa/ical2struct/ics"
)

func main() {
  bytes, _ := ioutil.ReadFile("test/data/alarm.ics")
  text := string(bytes)
  fmt.Println(ics.ComposeIcs(ics.ParseIcs(text).Events))
}
