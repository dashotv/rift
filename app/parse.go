package app

import (
	"regexp"
	"strconv"
)

var seasonRegex = regexp.MustCompile(`(?i)(?:s|season)(?:\s*)(\d{1,2})`)
var episodeRegex = regexp.MustCompile(`(?i)[e第](?:p)*(?:isode)*\s*(\d+)`)

func ParseFulltitle(title string) (string, int, int, error) {
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

	return title, season, episode, nil
}
