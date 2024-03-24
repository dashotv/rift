package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYtdlpList(t *testing.T) {
	list, err := ProcessURL("https://myanime.live/2024/03/23/soul-land-2-the-peerless-tang-sect-episode-41-english-sub/")
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotEmpty(t, list)

	for _, e := range list {
		fmt.Printf("url: %s\n", e.WebpageURL)
		fmt.Printf("url: %s\n", e.URL)
		// spew.Dump(e)
	}
}
