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

type ImageMessageRequest struct {
	Path    string `json:"path"`
	Caption string `json:"caption,omitempty"`
	Number  string `json:"number"`
}

type imageMessageRequest struct {
	Session string `json:"session"`
	*ImageMessageRequest
}

type ImageMessageResponse struct {
	Result    int         `json:"result"`
	Type      string      `json:"type"`
	MessageId interface{} `json:"messageId"`
	Session   string      `json:"session"`
	File      string      `json:"file"`
	Data      interface{} `json:"data"`
}

type imageMessageResponseInternal struct {
	Data   ImageMessageResponse `json:"data"`
	Result int                  `json:"result"`
}

func (s *Client) SendImageMessage(ctx context.Context, d *ImageMessageRequest) (*ImageMessageResponse, error) {
	if d == nil {
		return nil, fmt.Errorf("missing request object")
	}

	req := imageMessageRequest{
		Session:             s.sessionKey,
		ImageMessageRequest: d,
	}

	resp, err := s.request(ctx, req, http.MethodPost, sendImageEndpoint)
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
		return nil, fmt.Errorf("failed to send image message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn imageMessageResponseInternal
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}

type FileMessageRequest struct {
	Path   string `json:"path"`
	Number string `json:"number"`
}

type fileMessageRequest struct {
	Session string `json:"session"`
	*FileMessageRequest
}

type FileMessageResponse struct {
	Ack           int                       `json:"ack"`
	Id            string                    `json:"id"`
	SendMsgResult FileMessageResponseResult `json:"sendMsgResult"`
}

type FileMessageResponseResult struct {
	MessageSendResult string `json:"messageSendResult"`
}

type fileMessageResponseInternal struct {
	Data    FileMessageResponse `json:"data"`
	Result  int                 `json:"result"`
	Type    string              `json:"type"`
	Session string              `json:"session"`
}

func (s *Client) SendFileMessage(ctx context.Context, d *FileMessageRequest) (*FileMessageResponse, error) {
	if d == nil {
		return nil, fmt.Errorf("missing request object")
	}

	req := fileMessageRequest{
		Session:            s.sessionKey,
		FileMessageRequest: d,
	}

	resp, err := s.request(ctx, req, http.MethodPost, sendFileEndpoint)
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
		return nil, fmt.Errorf("failed to send image message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn fileMessageResponseInternal
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}
