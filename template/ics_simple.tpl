BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
{{- range $i, $event := . }}
{{ if and ($event.DateStart) ($event.DateEnd) -}}
BEGIN:VEVENT
DTEND;VALUE=DATE:{{$event.DateStart}}
DTSTART;VALUE=DATE:{{$event.DateEnd}}
{{ if ($event.Uid) -}}
UID:{{$event.Uid}}
{{- end }}
{{ if ($event.Description) -}}
DESCRIPTION:{{$event.Description}}
{{- end }}
{{ if ($event.Summary) -}}
SUMMARY:{{$event.Summary}}
{{- end }}
END:VEVENT
{{- end }}
{{- end }}
END:VCALENDAR