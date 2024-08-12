package app

import (
	"regexp"
	"strconv"
)

var seasonRegex = regexp.MustCompile(`(?i)(?:s|season)(?:\s*)(\d{1,2})`)
var episodeRegex = regexp.MustCompile(`(?i)[e第](?:p)*(?:isode)*\s*(\d+)`)
var urlRegex = regexp.MustCompile(`(?i)(?:season[\s-]+(\d+))*[\s-]+episode[\s-]+(\d+)(?:[\s-]+(\d+))*`)

func ParseFulltitle(title string) (int, int) {
	season := 0
	episode := 0

	if results := seasonRegex.FindStringSubmatch(title); len(results) == 2 {
		s, _ := strconv.Atoi(results[1])
		season = s
	}

	if results := episodeRegex.FindStringSubmatch(title); len(results) == 2 {
		e, _ := strconv.Atoi(results[1])
		episode = e
	}

	return season, episode
}

func ParseURL(url string) (int, int) {
	season := 0
	episode := 0

	if results := urlRegex.FindStringSubmatch(url); len(results) >= 3 {
		if results[1] != "" {
			s, _ := strconv.Atoi(results[1])
			season = s
		}
		if results[3] != "" {
			a, _ := strconv.Atoi(results[3])
			episode = a
		} else {
			e, _ := strconv.Atoi(results[2])
			episode = e
		}
	}

	return season, episode
}
