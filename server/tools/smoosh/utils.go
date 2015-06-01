package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const tvdbBannerURL = "http://thetvdb.com/banners/"

func isset(data []byte) bool {
	return len(data) > 0 && string(data) != "N/A"
}

var REGEXP_DIGITS = regexp.MustCompile(`\d+`)

func atoi(data []byte) int {
	match := REGEXP_DIGITS.Find(data)
	num, err := strconv.Atoi(string(match))
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func unpipe(data []byte) string {
	s := string(data)
	s = strings.Trim(s, "|")
	return strings.Join(strings.Split(s, "|"), ", ")
}

func uncomma(data []byte) int {
	s := string(data)
	num, err := strconv.Atoi(strings.Join(strings.Split(s, ","), ""))
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseDateTime(data []byte) time.Time {
	time, _ := time.Parse("2006-01-02 15:04:05", string(data))
	return time
}

func parseDate(data []byte) time.Time {
	time, _ := time.Parse("2006-01-02", string(data))
	return time
}
