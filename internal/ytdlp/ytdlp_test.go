package ytdlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYtdlpList(t *testing.T) {
	list, err := ProcessURL("https://myanime.live/2024/10/24/shixiong-a-shixiong-big-brother-episode-60-english-sub/")
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
