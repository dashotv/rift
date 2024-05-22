package scraper

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

type jhdanimeTest struct {
	name string
	url  string
}

func TestJhdAnime_Read(t *testing.T) {
	tests := []jhdanimeTest{
		{name: "hidden sect leader", url: "https://jhdanime.live/category/one-of-the-hidden-sect-leaders-is-shocking-s1/"},
	}
	client := NewJhdAnime(zap.NewExample().Sugar())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := client.Read(tt.url)
			for _, l := range list {
				fmt.Printf("jhdanime: %s\n", l)
			}
		})
	}
}
