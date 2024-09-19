package myzap_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTags(t *testing.T) {
	ctx := context.Background()
	result, err := client.GetTags(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestAddNewTag(t *testing.T) {
	ctx := context.Background()
	err := client.CreateTag(ctx, "teste")
	require.NoError(t, err)
}
