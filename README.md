[![Go Reference](https://pkg.go.dev/badge/github.com/v010maaa/ical2struct.svg)](https://pkg.go.dev/github.com/v010maaa/ical2struct)

# ical2struct

golang package to convert ICS/ICal text to struct

# Install

```
go get github.com/v010maaa/ical2struct
```

# Usage

```go
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
```

```go
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
```

```go
type Calendar struct {
  Version  string
  Calscale string
  Events   []Event
}


type Event struct {
	DateStart   string
	DateEnd     string
	Uid         string
	Description string
	Summary     string
}

```

Please check some example [here](https://pkg.go.dev/github.com/v010maaa/ical2struct/ics)

# Supported property

## Calender

| Prop     | example   |
| -------- | --------- |
| VERSION  | 2.0       |
| CALSCALE | GREGORIAN |

## Event

| Prop        | example   |
| ----------- | --------- |
| DTSTART     | 20220224  |
| DTEND       | 20220223  |
| UID         | 1         |
| SUMMARY     | Text text |
| DESCRIPTION | text text |

# License

[MIT](https://github.com/v010maaa/ical2struct/blob/main/LICENSE)
