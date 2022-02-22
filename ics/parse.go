/*
ここはパッケージコメントの最初になるから見出しではないよー

Hで始まり単一行かつ句読点なしかつ前が見出しではないのでこれは見出し

段落段落
段落段落
段落段落

次の段落
次の段落

    整形済みテキスt

次のやつはリンクになるはず。
https://golang.org/
*/
package ics

import (
    "bufio"
    "bytes"
    "fmt"
    "ical2json/models"
    "ical2json/utils"
)

func ParseIcs(text string) {
    buf := bytes.NewBufferString(text)
    scanner := bufio.NewScanner(buf)

    calender := models.NewCalender()

    for scanner.Scan() {
        // event := models.Event{}
        // property.SetEventJson(scanner.Text())
        line := scanner.Text()

        if fk := utils.HasCalendar(line); fk.Field != "none" {
            fmt.Println(fk)
            calender.SetCalendar(fk)
            continue
        }

        if fk := utils.HasEventScope(line); fk.Value == "VEVENT" {
            if fk.Field == "start" {
                fmt.Println("VVVVVVVVVV._start_.VVVVVVVVVVVV")
            } else if fk.Field == "end" {
                fmt.Println("VVVVVVVVVV.end.VVVVVVVVVVVV\n")
            }
            continue
        }

        if fk := utils.HasEvent(line); fk.Field != "none" {
            fmt.Println(fk)
        }
    }
}
