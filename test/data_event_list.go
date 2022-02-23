package test

import "ical2json/models"

func NoEndEvents() []models.Event {
	return []models.Event{
		{
			DateStart:   "20220220",
			Uid:         "1",
			Description: "Tehe Pero",
			Summary:     "Jnjdunk (not available)",
		},
		{
			DateStart:   "90220220",
			DateEnd:     "90220219",
			Uid:         "3",
			Description: "Japanese",
			Summary:     "日本語😎 (not available)",
		},
	}
}

func RamdomEvents() []models.Event {
	return []models.Event{
		{
			DateStart:   "20220220",
			Description: "Tehe Pero",
			Summary:     "Jnjdunk (not available)",
		},
		{
			DateStart:   "90220220",
			DateEnd:     "90220219",
			Uid:         "3",
			Description: "Japanese",
			Summary:     "日本語😎 (not available)",
		},
		{
			DateStart: "90220220",
			DateEnd:   "90220219",
			Uid:       "3",
			Summary:   "日本語😎 (not available)",
		},
		{
			DateStart:   "90220220",
			DateEnd:     "90220219",
			Uid:         "3",
			Description: "Japanese",
		},
	}
}
