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

func TestSendImageMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendImageMessage(ctx, &myzap.ImageMessageRequest{
		Number: os.Getenv("PHONE"),
		Path:   "https://cdn.mos.cms.futurecdn.net/VFLt5vHV7aCoLrLGjP9Qwm-1200-80.jpg",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendFileMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendFileMessage(ctx, &myzap.FileMessageRequest{
		Number: os.Getenv("PHONE"),
		Path:   "https://storage.googleapis.com/verbeux-public-images/cdn/ebook-corretores.pdf",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
