package scraper

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestDonghuaStream_Read(t *testing.T) {
	url := "https://donghuastream.org/"
	client := New("donghuastream", zap.NewExample().Sugar())
	list := client.Read(url)
	for _, l := range list {
		fmt.Printf("naruldonghua: %s\n", l)
	}
}

func TestDonghuaStream_Parse(t *testing.T) {
	url := "https://donghuastream.org/"
	client := New("donghuastream", zap.NewExample().Sugar())
	list := client.Parse(url)
	for _, l := range list {
		fmt.Printf("donghuastream: %+v\n", l)
	}
}
