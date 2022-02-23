package utils

import (
	"strings"

	"github.com/v010maaa/ical2struct/models"
	"github.com/v010maaa/ical2struct/prop"
)

// FeatKey map feature sentence to field name we defined
type FeatKey map[string]string

func HasCalendar(line string) models.FieldValue {
	fv := models.NewFieldValue()
	pre_text := strings.Split(line, ":")[0]

	if key, ok := prop.CalendarFeatKey()[pre_text]; ok {
		fv.Field = key
		fv.Value = line[len(pre_text)+1:]
	}
	return fv
}

//models.SetInfo
func HasEvent(line string) models.FieldValue {
	fv := models.NewFieldValue()

	if strings.HasPrefix(line, "DTEND") || strings.HasPrefix(line, "DTSTART") {
		pre_text := strings.Split(line, ";")[0]
		if key, ok := prop.EventFeatKey()[pre_text]; ok {
			fv.Field = key
			fv.Value = line[strings.Index(line, "=")+1:]
		}
	} else {
		pre_text := strings.Split(line, ":")[0]
		if key, ok := prop.EventFeatKey()[pre_text]; ok {
			fv.Field = key
			fv.Value = line[len(pre_text)+1:]
		}
	}
	return fv
}

func HasEventScope(line string) models.FieldValue {
	fv := models.NewFieldValue()

	pre_text := strings.Split(line, ":")[0]

	if key, ok := prop.EventScopeFeatKey()[pre_text]; ok {
		fv.Field = key
		fv.Value = line[len(pre_text)+1:]
	}
	return fv
}
