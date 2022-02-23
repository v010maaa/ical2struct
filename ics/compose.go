package ics

import (
	"bytes"
	"ical2json/models"
	"log"
	"text/template"
)

const templateFile = "template/ics_simple.text.go"

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
