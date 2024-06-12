package scraper

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestNarulDonghua_Read(t *testing.T) {
	url := "https://naruldonghua.xyz/"
	client := NewNarulDonghua(zap.NewExample().Sugar())
	list := client.Read(url)
	for _, l := range list {
		fmt.Printf("naruldonghua: %s\n", l)
	}
}

func TestNarulDonghua_Parse(t *testing.T) {
	url := "https://naruldonghua.xyz/"
	client := NewNarulDonghua(zap.NewExample().Sugar())
	list := client.Parse(url)
	for _, l := range list {
		fmt.Printf("naruldonghua: %+v\n", l)
	}
}
