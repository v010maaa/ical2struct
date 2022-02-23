package ics_test

import (
	"fmt"
	"ical2struct/ics"
	"ical2struct/models"
)

func ExampleComposeIcs() {
	event_list := []models.Event{
		{
			DateStart:   "20220223",
			DateEnd:     "20220224",
			Uid:         "1",
			Description: "English Class",
			Summary:     "English (not available)",
		},
		{
			DateStart:   "20220228",
			DateEnd:     "20220301",
			Uid:         "2",
			Description: "Japanese",
			Summary:     "Japanese (reserved)",
		},
	}

	text := ics.ComposeIcs(event_list)
	fmt.Println(text)
	// Output: BEGIN:VCALENDAR
	// VERSION:2.0
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// DTEND;VALUE=20220224
	// DTSTART;VALUE=20220223
	// UID:1
	// DESCRIPTION:English Class
	// SUMMARY:English (not available)
	// END:VEVENT
	// BEGIN:VEVENT
	// DTEND;VALUE=20220301
	// DTSTART;VALUE=20220228
	// UID:2
	// DESCRIPTION:Japanese
	// SUMMARY:Japanese (reserved)
	// END:VEVENT
	// END:VCALENDAR
}

func ExampleParseIcs() {
	text := `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTEND;VALUE=20220224
DTSTART;VALUE=20220223
UID:1
DESCRIPTION:English Class
SUMMARY:English (not available)
END:VEVENT
BEGIN:VEVENT
DTEND;VALUE=20220301
DTSTART;VALUE=20220228
UID:2
DESCRIPTION:Japanese
SUMMARY:Japanese (reserved)
END:VEVENT
END:VCALENDAR`

	calendar := ics.ParseIcs(text)
	fmt.Printf(calendar.Calscale)
	//Output: GREGORIAN
}
