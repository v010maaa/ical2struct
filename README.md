# ical2struct

golang package to parse ical to json

get start

```
sss
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
