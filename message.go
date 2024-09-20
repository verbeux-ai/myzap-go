package myzap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/verbeux-ai/myzap-go/listener"
)

type TextMessageRequest struct {
	Number     string `json:"number"`
	Text       string `json:"text"`
	TimeTyping int    `json:"timeTyping"`
	MarkIsRead bool   `json:"markIsRead"`
}

type textMessageRequest struct {
	Session string `json:"session"`
	*TextMessageRequest
}

type textMessageResponseInternal struct {
	Data   listener.WebhookMessageData `json:"data"`
	Result int                         `json:"result"`
}

type TextMessageResponse struct {
	*listener.WebhookMessageData
}

func (s *Client) SendTextMessage(ctx context.Context, d *TextMessageRequest) (*TextMessageResponse, error) {
	if d == nil {
		return nil, fmt.Errorf("missing request object")
	}

	req := textMessageRequest{
		Session:            s.sessionKey,
		TextMessageRequest: d,
	}

	req.TextMessageRequest.TimeTyping = req.TextMessageRequest.TimeTyping * 1000

	resp, err := s.request(ctx, req, http.MethodPost, sendTextEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to send text message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn textMessageResponseInternal
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &TextMessageResponse{&toReturn.Data}, nil
}
