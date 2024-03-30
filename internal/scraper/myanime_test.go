package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideos(t *testing.T) {
	m := NewMyAnime()
	urls := m.Read("https://jhdanime.live/?s=ancient+lords")
	assert.NotEmpty(t, urls, "expected results")
	for _, u := range urls {
		fmt.Printf("url: %s\n", u)
	}
}
