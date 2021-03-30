package jk

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_HudsonGet(t *testing.T) {
	c := ClientFromTestENV()
	h, err := c.HudsonGet(context.Background())
	require.NoError(t, err)
	t.Logf("%+v", h)
	v, err := c.ViewGet(context.Background(), os.Getenv("JK_TEST_VIEW"))
	require.NoError(t, err)
	t.Logf("%+v", v)
	j, err := c.JobGetConfig(context.Background(), os.Getenv("JK_TEST_JOB"))
	require.NoError(t, err)
	t.Logf("%s", j)
}
