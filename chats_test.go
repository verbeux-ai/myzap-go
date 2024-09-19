package myzap_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTag(t *testing.T) {
	ctx := context.Background()
	err := client.AddChatTag(ctx, "1", os.Getenv("NUMBER"))
	require.NoError(t, err)
}
