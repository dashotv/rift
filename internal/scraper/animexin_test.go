package scraper

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestAnimeXin_Read(t *testing.T) {
	url := "https://animexin.vip/"
	client := NewAnimeXin(zap.NewExample().Sugar())
	list := client.Read(url)
	for _, l := range list {
		fmt.Printf("animexin: %s\n", l)
	}
}

func TestAnimeXin_Parse(t *testing.T) {
	url := "https://animexin.vip/"
	client := NewAnimeXin(zap.NewExample().Sugar())
	list := client.Parse(url)
	for _, l := range list {
		fmt.Printf("animexin: %+v\n", l)
	}
}
