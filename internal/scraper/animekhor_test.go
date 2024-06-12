package scraper

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestAnimeKhor_Read(t *testing.T) {
	url := "https://animekhor.xyz/"
	client := NewAnimeKhor(zap.NewExample().Sugar())
	list := client.Read(url)
	for _, l := range list {
		fmt.Printf("animekhor: %s\n", l)
	}
}

func TestAnimeKhor_Parse(t *testing.T) {
	url := "https://animekhor.xyz/"
	client := NewAnimeKhor(zap.NewExample().Sugar())
	list := client.Parse(url)
	for _, l := range list {
		fmt.Printf("animekhor: %+v\n", l)
	}
}
