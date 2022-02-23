/*
Package ics convert ics to structures.

Basic rules in ics format are processed using conditions without the use of lexers.
Currently, it supports only a limited number of properties, but it is a file structure that can be applied to all properties in the future.
	struct -> icls
	icls -> struct

Performance and validation haven't been considered yet, but it does provide minimal configuration functionality that can be used to link third-party calendar features such as airbnb.
Then, it has already started operation at an early stage of MVP and will improve it based on many requests.

https://golang.org/
*/
package ics

import (
	"bytes"
	"log"
	"text/template"

	"github.com/v010maaa/ical2struct/models"
)

const temp_simple = `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
{{ range $i, $el := . -}}
{{ if and ($el.DateStart) ($el.DateEnd) -}}
BEGIN:VEVENT
DTEND;VALUE={{$el.DateEnd}}
DTSTART;VALUE={{$el.DateStart}}
{{ if ($el.Uid) -}}
UID:{{$el.Uid}}
{{ end -}}
{{ if ($el.Description) -}}
DESCRIPTION:{{$el.Description}}
{{ end -}}
{{ if ($el.Summary) -}}
SUMMARY:{{$el.Summary}}
{{ end -}}
END:VEVENT
{{ end -}}
{{ end -}}
END:VCALENDAR`

func ComposeIcs(el []models.Event) string {
	t, err := template.New("simple_temp").Parse(temp_simple)

	if err != nil {
		log.Fatal(err)
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, el); err != nil {
		log.Fatal(err)
	}

	return tpl.String()
}
