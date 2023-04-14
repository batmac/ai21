package ai21

import "net/http"

type CompleteRequest struct {
	Prompt           string   `json:"prompt"`
	NumResults       int      `json:"numResults,omitempty"`
	MaxTokens        int      `json:"maxTokens,omitempty"`
	MinTokens        int      `json:"minTokens,omitempty"`
	Temperature      float64  `json:"temperature,omitempty"`
	TopP             float64  `json:"topP,omitempty"`
	StopSequences    []string `json:"stopSequences,omitempty"`
	TopKReturn       int      `json:"topKReturn,omitempty"`
	FrequencyPenalty Penalty  `json:"frequencyPenalty,omitempty"`
	PresencePenalty  Penalty  `json:"presencePenalty,omitempty"`
	CountPenalty     Penalty  `json:"countPenalty,omitempty"`
}

type Penalty struct {
	Scale               int  `json:"scale"`
	ApplyToNumbers      bool `json:"applyToNumbers"`
	ApplyToPunctuations bool `json:"applyToPunctuations"`
	ApplyToStopwords    bool `json:"applyToStopwords"`
	ApplyToWhitespaces  bool `json:"applyToWhitespaces"`
	ApplyToEmojis       bool `json:"applyToEmojis"`
}

type CompleteResponse struct {
	ID          string       `json:"id"`
	Prompt      Prompt       `json:"prompt"`
	Completions []Completion `json:"completions"`
}

type Prompt struct {
	Text   string   `json:"text"`
	Tokens []Tokens `json:"tokens"`
}

type Tokens struct {
	GeneratedToken GeneratedToken `json:"generatedToken"`
	TopTokens      interface{}    `json:"topTokens"`
	TextRange      TextRange      `json:"textRange"`
}

type GeneratedToken struct {
	Token      string  `json:"token"`
	Logprob    float64 `json:"logprob"`
	RawLogprob float64 `json:"raw_logprob"`
}

type TextRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Completion struct {
	Data         Data         `json:"data"`
	FinishReason FinishReason `json:"finishReason"`
}

type Data struct {
	Text   string   `json:"text"`
	Tokens []Tokens `json:"tokens"`
}

type FinishReason struct {
	Reason string `json:"reason"`
	Length int    `json:"length"`
}

func (c *Client) Complete(m Model, r CompleteRequest) (CompleteResponse, error) {
	path := "j2-" + string(m) + "/complete"
	return req[CompleteRequest, CompleteResponse](c, http.MethodPost, path, r)
}

func (c *Client) EasyComplete(m Model, prompt string, maxTokens int) (string, error) {
	r := CompleteRequest{
		Prompt:        prompt,
		NumResults:    1,
		MaxTokens:     maxTokens,
		StopSequences: []string{"##"},
	}

	resp, err := c.Complete(m, r)
	if err != nil {
		return "", err
	}
	return resp.Completions[0].Data.Text, nil
}
