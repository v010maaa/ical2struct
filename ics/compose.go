package ics

import (
	"bytes"
	"ical2json/models"
	"log"
	"text/template"
)

const templateFile = "template/ics_simple.tpl"

func ComposeIcs(event []models.Event) string {
	t, err := template.ParseFiles(templateFile)

	if err != nil {
		log.Fatal(err)
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, event); err != nil {
		log.Fatal(err)
	}

	return tpl.String()
}
