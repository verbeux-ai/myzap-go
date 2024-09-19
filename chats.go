package myzap

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
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
