package ytdlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYtdlpList(t *testing.T) {
	list, err := ProcessURL("https://jhdanime.live/2024/05/17/one-of-the-hidden-sect-leaders-is-shocking-ep-20/")
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotEmpty(t, list)

	for _, e := range list {
		fmt.Printf("web: %s\n", e.WebpageURL)
		fmt.Printf("url: %s\n", e.URL)
		// spew.Dump(e)
	}
}
