package myzap_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/myzap-go"
)

func TestStartInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.Start(ctx, &myzap.StartRequest{
		Session:   os.Getenv("SESSION_KEY"),
		WhStatus:  "https://webhook.site/34c00a56-a2f6-41b1-8a6a-c44b14677c54",
		WhMessage: "https://webhook.site/34c00a56-a2f6-41b1-8a6a-c44b14677c54",
		WhQrcode:  "https://webhook.site/34c00a56-a2f6-41b1-8a6a-c44b14677c54",
		WhConnect: "https://webhook.site/34c00a56-a2f6-41b1-8a6a-c44b14677c54",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestDeleteInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.Delete(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
