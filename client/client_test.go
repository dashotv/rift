package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_PageIndex(t *testing.T) {
	c := New("http://localhost:9006/api")
	resp, err := c.PageService.Index(context.Background(), &Request{})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), resp.Total)
}
