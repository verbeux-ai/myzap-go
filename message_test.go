package myzap_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/myzap-go"
)

func TestSendTextMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTextMessage(ctx, &myzap.TextMessageRequest{
		Number:     os.Getenv("PHONE"),
		Text:       "Teste",
		MarkIsRead: true,
		TimeTyping: 5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
