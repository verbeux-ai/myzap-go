package myzap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WebhookDeliveryRequest struct {
	Value string `json:"value,omitempty"`
}

type StartRequest struct {
	Session   string `json:"session"`
	WhStatus  string `json:"wh_status"`
	WhMessage string `json:"wh_message"`
	WhQrcode  string `json:"wh_qrcode"`
	WhConnect string `json:"wh_connect"`
}

type StartResponse struct {
	Result  string `json:"result"`
	Session string `json:"session"`
	State   string `json:"state"`
	Status  string `json:"status"`
}

func (s *Client) Start(ctx context.Context, d *StartRequest) (*StartResponse, error) {
	resp, err := s.request(ctx, d, http.MethodPost, startInstanceEndpoint)
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
		return nil, fmt.Errorf("failed to start sessionKey with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn StartResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

type DeleteResponse struct {
	Logout  bool   `json:"logout"`
	Close   bool   `json:"close"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (s *Client) Delete(ctx context.Context) (*DeleteResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodPost, deleteInstanceEndpoint)
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
		return nil, fmt.Errorf("failed to start sessionKey with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn DeleteResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
