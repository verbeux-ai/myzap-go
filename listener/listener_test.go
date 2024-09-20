package listener_test

import (
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/myzap-go/listener"
)

func TestListener_OnMessage(t *testing.T) {
	client := listener.NewMessageListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnMessage(func(message *listener.WebhookMessage) error {
		defer wg.Done()
		require.NotEmpty(t, message)
		return nil
	})

	wg.Wait()

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(messageExample1)))
	require.NoError(t, err)

}
