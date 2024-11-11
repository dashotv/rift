package scraper

import (
	"os"
	"strings"
)

var version string

func init() {
	b, err := os.ReadFile("../../VERSION")
	if err != nil {
		panic(err)
	}
	version = strings.TrimSuffix(string(b), "\n")
}
