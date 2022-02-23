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

package main

import (
  "fmt"
  "ical2json/ics"
  "io/ioutil"
)

func main() {
  bytes, _ := ioutil.ReadFile("test/data/alarm.ics")
  text := string(bytes)
  fmt.Println(ics.ComposeIcs(ics.ParseIcs(text).Events))
}
