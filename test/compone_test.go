package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/v010maaa/ical2struct/ics"
)

func TestNomalParse(t *testing.T) {
	bytes, _ := ioutil.ReadFile("data/normal.ics")
	expected := string(bytes)
	cal := ics.ParseIcs(expected)
	actual := ics.ComposeIcs(cal.Events)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}

}

func TestAirBnbParse(t *testing.T) {
	input_bytes, _ := ioutil.ReadFile("data/airbnb.ics")
	expected_bytes, _ := ioutil.ReadFile("data/airbnb_expected.ics")

	input := string(input_bytes)
	expected := string(expected_bytes)

	cal := ics.ParseIcs(input)
	actual := ics.ComposeIcs(cal.Events)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}

func TestAlarmParse(t *testing.T) {
	input_bytes, _ := ioutil.ReadFile("data/alarm.ics")
	expected_bytes, _ := ioutil.ReadFile("data/alarm_expected.ics")

	input := string(input_bytes)
	expected := string(expected_bytes)

	cal := ics.ParseIcs(input)
	actual := ics.ComposeIcs(cal.Events)
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}

func TestNoEndCompose(t *testing.T) {
	event_list := NoEndEvents()
	actual := ics.ComposeIcs(event_list)

	expected_bytes, _ := ioutil.ReadFile("data/no_end_expected.ics")

	expected := string(expected_bytes)
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}

func TestRandomCompose(t *testing.T) {
	event_list := RamdomEvents()
	actual := ics.ComposeIcs(event_list)

	expected_bytes, _ := ioutil.ReadFile("data/ramdom_expected.ics")
	expected := string(expected_bytes)
	fmt.Println(actual)
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
}
