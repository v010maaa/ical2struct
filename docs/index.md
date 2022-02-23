# Overview

A minimal iCalendar / ICS parser for Golang.

The functions are the following two:

1. Converting iCalendar / ICS text to a struct
2. Converting a struct to an ics format

Since this is a project that started the other day, we are looking for friends who will contribute to this project.

[Go Doc](https://pkg.go.dev/github.com/v010maaa/ical2struct/ics)

## Import package

```go
go get github.com/v010maaa/ical2struct
```

## Use package

```go
import (
	github.com/v010maaa/ical2struct/ics
    github.com/v010maaa/ical2struct/models
)
```

The response type is defined in models.

```go
text := ics.ComposeIcs(event_list)
calendar := ics.ParseIcs(text)
```

# Supported property

Please request, if we want to use more props.

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

<meta name="google-site-verification" content="PnpQheTxB5s_HSy1bY1T_khavRWjDUbd74QdSi2AZxQ" />
