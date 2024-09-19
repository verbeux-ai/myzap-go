package myzap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type statusResponseInternal struct {
	Data   StatusResponse `json:"data"`
	Result int            `json:"result"`
	Status string         `json:"status"`
}

type StatusResponse struct {
	Attempts       int         `json:"attempts"`
	AttemptsStart  int         `json:"attempts_start"`
	Battery        interface{} `json:"battery"`
	BlockStoreAdds string      `json:"blockStoreAdds"`
	ClientToken    interface{} `json:"clientToken"`
	Connected      interface{} `json:"connected"`
	CreatedAt      time.Time   `json:"created_at"`
	Id             int         `json:"id"`
	Is24H          interface{} `json:"is24h"`
	IsResponse     interface{} `json:"isResponse"`
	LastConnect    time.Time   `json:"last_connect"`
	LastDisconnect time.Time   `json:"last_disconnect"`
	LastStart      time.Time   `json:"last_start"`
	Lc             interface{} `json:"lc"`
	Lg             interface{} `json:"lg"`
	Locales        interface{} `json:"locales"`
	Number         string      `json:"number"`
	Platform       string      `json:"platform"`
	Plugged        interface{} `json:"plugged"`
	ProtoVersion   interface{} `json:"protoVersion"`
	Pushname       string      `json:"pushname"`
	QrCode         string      `json:"qrCode"`
	Ref            string      `json:"ref"`
	RefTTL         string      `json:"refTTL"`
	ServerToken    interface{} `json:"serverToken"`
	Session        string      `json:"session"`
	Sessionkey     string      `json:"sessionkey"`
	SmbTos         string      `json:"smbTos"`
	State          string      `json:"state"`
	Status         string      `json:"status"`
	Tos            interface{} `json:"tos"`
	UpdatedAt      time.Time   `json:"updated_at"`
	UrlCode        string      `json:"urlCode"`
	UserId         int         `json:"user_id"`
	WaJsVersion    string      `json:"wa_js_version"`
	WaVersion      string      `json:"wa_version"`
	WhConnect      string      `json:"wh_connect"`
	WhMessage      string      `json:"wh_message"`
	WhQrcode       string      `json:"wh_qrcode"`
	WhStatus       string      `json:"wh_status"`
}

type QrCodeResponse struct {
	ImageBase64 string `json:"img"`
}

func (s *Client) Status(ctx context.Context) (*StatusResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodPost, getConnectionStatusEndpoint)
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
		return nil, fmt.Errorf("failed to get status with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn statusResponseInternal
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}

func (s *Client) QrCodeImage(ctx context.Context) (*QrCodeResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, getQrCodeEndpoint+"?session="+s.sessionKey+"&sessionkey="+s.sessionKey)
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

	var toReturn QrCodeResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
