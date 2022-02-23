package ics

import (
    "bufio"
    "bytes"

    "github.com/v010maaa/ical2struct/models"
    "github.com/v010maaa/ical2struct/utils"
)

func ParseIcs(text string) models.Calendar {
    buf := bytes.NewBufferString(text)
    scanner := bufio.NewScanner(buf)

    calendar := models.NewCalender()
    event_list := []models.Event{}
    temp_event := models.Event{}

    event_flag := false
    for scanner.Scan() {
        line := scanner.Text()
        if event_flag {
            if fv := utils.HasEvent(line); fv.Field != "none" {
                temp_event.SetEvent(fv)
                continue
            }
        }

        if fv := utils.HasCalendar(line); fv.Field != "none" {
            calendar.SetCalendar(fv)
            continue
        }

        if fv := utils.HasEventScope(line); fv.Value == "VEVENT" {
            if fv.Field == "End" {
                event_list = append(event_list, temp_event)
                event_flag = false
                temp_event = models.Event{}
            } else if fv.Field == "Start" {
                event_flag = true
            }
        }
    }
    calendar.SetEventList(event_list)
    return calendar
}
