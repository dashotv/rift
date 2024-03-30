package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_PageIndex(t *testing.T) {
	c := New("http://localhost:59006/")
	c.SetDebug(true)
	resp, err := c.Page.Index(context.Background(), &PageIndexRequest{Page: 1, Limit: 1})
	assert.NoError(t, err)
	// assert.Equal(t, int64(30), resp.Total)
	fmt.Printf("resp: %+v\n", resp)
}
