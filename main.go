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

package main // mainパッケージであることを宣言

import (
  "ical2json/ics"
)

func main() {

  text := `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTEND;VALUE=DATE:20220220
DTSTART;VALUE=DATE:20220219
UID:1
DESCRIPTION:Checkinn
SUMMARY:Checkinn (not available)
END:VEVENT
BEGIN:VEVENT
DTEND;VALUE=DATE:90220220
DTSTART;VALUE=DATE:90220219
UID:3
DESCRIPTION:Checkinn
SUMMARY:Checkinn (not available)
END:VEVENT
END:VCALENDAR`
  // 最初に実行されるmain()関数を定義
  ics.ParseIcs(text)

  //   event_list := []models.Event{
  //     {
  //         DateStart:"20220220",
  //         DateEnd: "20220219",
  //         Uid:"1",
  //         Description:"Checkinn",
  //         Summary:"Checkinn (not available)",
  //     },
  //     {
  //         DateStart:"90220220",
  //         DateEnd: "90220219",
  //         Uid:"3",
  //         Description:"Checkinn",
  //         Summary:"Checkinn (not available)",
  //     },
  // }
  //   fmt.Println(ics.JsonToIcals(event_list))
}
