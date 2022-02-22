package utils

import (
	"ical2json/models"
	"ical2json/prop"
	"strings"
)

// FeatKey map feature sentence to field name we defined
type FeatKey map[string]string

//models.SetInfo
func HasEvent(line string) models.FieldValue {
	result := models.NewFieldValue()
	pre_text := strings.Split(line, ":")[0]

	if key, ok := prop.EventFeatKey()[pre_text]; ok {
		result.Field = key
		result.Value = line[len(pre_text)+1:]
	}
	return result
}

func HasCalendar(line string) models.FieldValue {
	result := models.NewFieldValue()
	pre_text := strings.Split(line, ":")[0]

	if key, ok := prop.CalendarFeatKey()[pre_text]; ok {
		result.Field = key
		result.Value = line[len(pre_text)+1:]
	}
	return result
}

func HasEventScope(line string) models.FieldValue {
	result := models.NewFieldValue()
	pre_text := strings.Split(line, ":")[0]

	if key, ok := prop.EventScopeFeatKey()[pre_text]; ok {
		result.Field = key
		result.Value = line[len(pre_text)+1:]
	}
	return result
}
