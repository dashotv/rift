package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseCase struct {
	title   string
	season  int
	episode int
}

func TestParse(t *testing.T) {
	subjects := []parseCase{{"wspws1ep78eng", 1, 78}}
	for _, subject := range subjects {
		t.Run(subject.title, func(t *testing.T) {
			_, s, e, err := ParseFulltitle(subject.title)
			assert.NoError(t, err)
			assert.Equal(t, subject.season, s)
			assert.Equal(t, subject.episode, e)
		})
	}
}
