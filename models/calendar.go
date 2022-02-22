package models

import "fmt"

type Calendar struct {
  Version  string
  Calscale string
  Event    Event
}

func (cal *Calendar) SetCalendar(fv FieldValue) {
  switch field := fv.Field; field {
  case "Version":
    cal.Version = field
  case "Calscale":
    cal.Version = field
  default:
    fmt.Println("Ical2Json not support fieald")
  }
}

func NewCalender() Calendar {
  return Calendar{Version: "GREGORIAN", Calscale: "2.0"}
}
