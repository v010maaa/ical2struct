package prop

type featKey map[string]string

func CalendarFeatKey() featKey {
	return featKey{
		"VERSION":  "Version",
		"CALSCALE": "Calscale",
	}
}

func EventFeatKey() featKey {
	return featKey{
		"DTEND":       "DateEnd",
		"DTSTART":     "DateStart",
		"DESCRIPTION": "Description",
		"SUMMARY":     "Summary",
		"UID":         "Uid",
	}
}

func EventScopeFeatKey() featKey {
	return featKey{
		"BEGIN": "Start",
		"END":   "End",
	}
}
