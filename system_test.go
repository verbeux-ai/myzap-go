package myzap_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	ctx := context.Background()
	result, err := client.Status(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	l, _ := json.Marshal(result)
	fmt.Print(string(l))
}

func TestQrCode(t *testing.T) {
	ctx := context.Background()
	result, err := client.QrCodeImage(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
