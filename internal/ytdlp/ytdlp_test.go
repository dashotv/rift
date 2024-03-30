package ytdlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYtdlpList(t *testing.T) {
	list, err := ProcessURL("https://jhdanime.live/2024/03/24/ancient-lords-episode-7-multi-sub/")
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotEmpty(t, list)

	for _, e := range list {
		fmt.Printf("web: %s\n", e.WebpageURL)
		fmt.Printf("url: %s\n", e.URL)
		// spew.Dump(e)
	}
}
