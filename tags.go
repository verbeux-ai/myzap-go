package myzap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type tagsRequestInternal struct {
	Session string `json:"session"`
}

type getTagsResponseInternal struct {
	Result int            `json:"result"`
	Token  []TagsResponse `json:"token"`
}

type TagsResponse struct {
	Color      interface{} `json:"color"`
	ColorIndex int         `json:"colorIndex"`
	Count      int         `json:"count"`
	HexColor   string      `json:"hexColor"`
	Id         string      `json:"id"`
	Name       string      `json:"name"`
}

func (s *Client) GetTags(ctx context.Context) ([]TagsResponse, error) {
	req := tagsRequestInternal{
		Session: s.sessionKey,
	}

	resp, err := s.request(ctx, req, http.MethodPost, tagsEndpoint)
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
		return nil, fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn getTagsResponseInternal
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return toReturn.Token, nil
}

type createTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	ID string `json:"id"`
}

func (s *Client) CreateTag(ctx context.Context, name string) error {
	req := &createTagRequest{Name: name}

	var err error
	for range 3 {
		err = s.createTagRetry(ctx, req)
		if err != nil {
			continue
		}

		err = nil
	}

	return err
}

func (s *Client) createTagRetry(ctx context.Context, req *createTagRequest) error {
	resp, err := s.request(ctx, req, http.MethodPost, createTagEndpoint)
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
		return fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	tags, err := s.GetTags(ctx)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		if strings.EqualFold(tag.Name, req.Name) {
			return nil
		}
	}

	return ErrTagNotFound
}
