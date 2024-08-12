package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseCase struct {
	title   string
	season  int
	episode int
}

func TestParseFulltitle(t *testing.T) {
	subjects := []parseCase{{"wspws1ep78eng", 1, 78}}
	for _, subject := range subjects {
		t.Run(subject.title, func(t *testing.T) {
			s, e := ParseFulltitle(subject.title)
			assert.Equal(t, subject.season, s)
			assert.Equal(t, subject.episode, e)
		})
	}
}

func TestParseURL(t *testing.T) {
	subjects := []parseCase{
		{"https://myanime.live/2024/05/19/battle-through-the-heavens-season-5-episode-97-english-sub/", 5, 97},
		{"https://myanime.live/2024/05/19/xiaobing-chuanqi-legend-of-soldier-episode-04-english-sub/", 0, 4},
		{"https://donghuastream.org/swallowed-star-season-4-episode-47-132-subtitles/", 4, 132},
	}
	for _, subject := range subjects {
		t.Run(subject.title, func(t *testing.T) {
			s, e := ParseURL(subject.title)
			assert.Equal(t, subject.season, s)
			assert.Equal(t, subject.episode, e)
		})
	}
}
