package ai21

import (
	"errors"
	"net/http"
)

type SummarizeRequest struct {
	Source     string     `json:"source"`
	SourceType SourceType `json:"sourceType"`
	Focus      string     `json:"focus,omitempty"`
}

type SummarizeResponse struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
	Detail  string `json:"detail"`
}

var (
	ErrSourceTooLong = errors.New("source too long")
	ErrFocusTooLong  = errors.New("focus too long")
)

func (c *Client) Summarize(r SummarizeRequest) (SummarizeResponse, error) {
	return req[SummarizeRequest, SummarizeResponse](c, http.MethodPost, "summarize", r)
}

func (c *Client) EasySummarize(src string, srcType SourceType, focus string) (string, error) {
	if len(src) > 50_000 {
		return "", ErrSourceTooLong
	}
	if len(focus) > 50 {
		return "", ErrFocusTooLong
	}
	resp, err := c.Summarize(SummarizeRequest{
		Source:     src,
		SourceType: srcType,
		Focus:      focus,
	})
	if err != nil {
		return "", err
	}
	return resp.Summary, nil
}
