package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAnimeXin_Read(t *testing.T) {
	url := "https://animexin.dev/"
	client := New("animexin", zap.NewExample().Sugar())
	list := client.Read(url)
	for _, l := range list {
		fmt.Printf("animexin: %s\n", l)
	}
}

func TestAnimeXin_Parse(t *testing.T) {
	url := "https://animexin.dev/"
	client := New("animexin", zap.NewExample().Sugar())
	list := client.Parse(url)
	require.NotEmpty(t, list, "expected results")
	for _, l := range list {
		fmt.Printf("animexin: %+v\n", l)
	}
}

func TestAnimeXin_ParseIndividual(t *testing.T) {
	url := "https://animexin.dev/100-000-years-of-refining-qi/"
	client := New("animexin", zap.NewExample().Sugar())
	list := client.Parse(url)
	require.NotEmpty(t, list, "expected results")
	for _, l := range list {
		fmt.Printf("animexin: %+v\n", l)
	}
}
