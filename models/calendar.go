package models

import "fmt"

type Calendar struct {
  Version  string
  Calscale string
  Events   []Event
}

func (cal *Calendar) SetCalendar(fv FieldValue) {
  switch field := fv.Field; field {
  case "Version":
    cal.Version = fv.Value
  case "Calscale":
    cal.Calscale = fv.Value
  default:
    fmt.Println("ical2struct not support fieald")
  }
}

func (cal *Calendar) SetEventList(el []Event) {
  cal.Events = el
}

func NewCalender() Calendar {
  return Calendar{Version: "GREGORIAN", Calscale: "2.0"}
}
