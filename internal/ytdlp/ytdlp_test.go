package ytdlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYtdlpList(t *testing.T) {
	list, err := ProcessURL("https://geo.dailymotion.com/player/x9mpt.html?video=k6iaxyaxLPkIRyEEzYE")
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotEmpty(t, list)

	for _, e := range list {
		fmt.Printf("name: %s\n", e.Fulltitle)
		fmt.Printf("web: %s\n", e.WebpageURL)
		fmt.Printf("url: %s\n", e.URL)
		// spew.Dump(e)
	}
}
