package models

import "fmt"

type Event struct {
	DateStart   string
	DateEnd     string
	Uid         string
	Description string
	Summary     string
}

func (ev *Event) SetEvent(fv FieldValue) {
	switch field := fv.Field; field {
	case "DateStart":
		ev.DateStart = fv.Value
	case "DateEnd":
		ev.DateEnd = fv.Value
	case "Uid":
		ev.Uid = fv.Value
	case "Description":
		ev.Description = fv.Value
	case "Summary":
		ev.Summary = fv.Value
	default:
		fmt.Println("ical2struct not support field")
	}
}
