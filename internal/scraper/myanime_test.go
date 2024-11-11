package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetVideos(t *testing.T) {
	l := zap.NewNop().Sugar()
	m := New("myanime", l)
	urls := m.Read("https://myanime.live/")
	assert.NotEmpty(t, urls, "expected results")
	for _, u := range urls {
		fmt.Printf("url: %s\n", u)
	}
}

func TestMyAnime_Parse(t *testing.T) {
	url := "https://myanime.live/"
	client := New("myanime", zap.NewExample().Sugar())
	list := client.Parse(url)
	for _, l := range list {
		fmt.Printf("myanime: %+v\n", l)
	}
}
