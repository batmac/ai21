package ai21

import (
	"errors"
	"net/http"
)

// https://docs.ai21.com/reference/contextual-answers-ref

type ContextualAnswersRequest struct {
	Context  string `json:"context"`
	Question string `json:"question"`
}

type ContextualAnswersResponse struct {
	Answer string `json:"answer"`
	ID     string `json:"id"`
	Detail string `json:"detail"`
}

var (
	ErrContextTooLong  = errors.New("context too long")
	ErrQuestionTooLong = errors.New("question too long")
)

func (c *Client) ContextualAnswers(r ContextualAnswersRequest) (ContextualAnswersResponse, error) {
	return req[ContextualAnswersRequest, ContextualAnswersResponse](c, http.MethodPost, "experimental/answer", r)
}

func (c *Client) EasyContextualAnswers(context, question string) (string, error) {
	if len(context) > 10000 {
		return "", ErrContextTooLong
	}
	if len(question) > 160 {
		return "", ErrQuestionTooLong
	}
	resp, err := c.ContextualAnswers(ContextualAnswersRequest{
		Context:  context,
		Question: question,
	})
	if err != nil {
		return "", err
	}
	return resp.Answer, nil
}
