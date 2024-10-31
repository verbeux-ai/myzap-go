package myzap

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type setTagRequestInternal struct {
	Number  string                        `json:"number"`
	Actions []setTagRequestActionInternal `json:"actions"`
}

type setTagRequestActionInternal struct {
	LabelID string `json:"labelId"`
	Type    string `json:"type"` // add or remove
}

func (s *Client) AddChatTag(ctx context.Context, tagID string, phone string) error {
	req := setTagRequestInternal{
		Number: phone,
		Actions: []setTagRequestActionInternal{
			{
				LabelID: tagID,
				Type:    "add",
			},
		},
	}

	resp, err := s.request(ctx, req, http.MethodPost, addOrRemoveTagEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyErr := errors.New(string(body))
		return fmt.Errorf("failed to set chat tag with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}

type startTypingRequestInternal struct {
	Number string `json:"number"`
	Time   int    `json:"time"`
}

// StartTyping start typing for a specific number for specific time.
// duration in milliseconds
func (s *Client) StartTyping(ctx context.Context, phone string, duration time.Duration) error {
	if duration < time.Millisecond {
		return errors.New("duration must be greater than 1ms")
	}

	req := startTypingRequestInternal{
		Number: phone,
		Time:   int(duration / time.Millisecond),
	}

	resp, err := s.request(ctx, req, http.MethodPost, startTypingEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyErr := errors.New(string(body))
		return fmt.Errorf("failed to start typing with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}

type stopTypingRequestInternal struct {
	Number string `json:"number"`
}

func (s *Client) StopTyping(ctx context.Context, phone string) error {
	req := stopTypingRequestInternal{
		Number: phone,
	}

	resp, err := s.request(ctx, req, http.MethodPost, stopTypingEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyErr := errors.New(string(body))
		return fmt.Errorf("failed to stop typing with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}
