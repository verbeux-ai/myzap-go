package myzap_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	ctx := context.Background()
	result, err := client.Status(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestQrCode(t *testing.T) {
	ctx := context.Background()
	result, err := client.QrCodeImage(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
