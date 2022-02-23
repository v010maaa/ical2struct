package main

import (
  "fmt"
  "ical2struct/ics"
  "io/ioutil"
)

func main() {
  bytes, _ := ioutil.ReadFile("test/data/alarm.ics")
  text := string(bytes)
  fmt.Println(ics.ComposeIcs(ics.ParseIcs(text).Events))
}
