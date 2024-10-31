package myzap_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddTag(t *testing.T) {
	ctx := context.Background()
	err := client.AddChatTag(ctx, "1", os.Getenv("PHONE"))
	require.NoError(t, err)
}

func TestStartTyping(t *testing.T) {
	ctx := context.Background()
	err := client.StartTyping(ctx, os.Getenv("PHONE"), time.Second*3)
	require.NoError(t, err)
}

func TestStopTyping(t *testing.T) {
	ctx := context.Background()
	err := client.StopTyping(ctx, os.Getenv("PHONE"))
	require.NoError(t, err)
}
