package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetVideos(t *testing.T) {
	l := zap.NewNop().Sugar()
	m := NewMyAnime(l)
	urls := m.Read("https://jhdanime.live/?s=ancient+lords")
	assert.NotEmpty(t, urls, "expected results")
	for _, u := range urls {
		fmt.Printf("url: %s\n", u)
	}
}
