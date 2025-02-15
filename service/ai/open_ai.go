package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/env"
	"io"
	"net/http"
)

type openAI struct{}

func (o openAI) Search(ctx context.Context, message string) (map[string]any, bool) {
	const openAPIURL = "https://api.openai.com/v1/chat/completions"
	jsonBody, err := json.Marshal(defaultOpenAIRequestWithMessages(message))
	if err != nil {
		//TODO add logs
		return nil, false
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, openAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		//TODO add logs
		return nil, false
	}
	o.setOpenAIChatRequestHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//TODO add logs
		return nil, false
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//TODO add logs
		return nil, false
	}
	bodyJSON := make(map[string]any)
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		//TODO add logs

	}
	if resp.StatusCode != 200 {
		return nil, false
	}

	choicesResponse, ok := bodyJSON["choices"]
	if !ok {
		return nil, false
	}
	choicesResponseArray, ok := choicesResponse.([]any)
	if !ok {
		return nil, false
	}

	if len(choicesResponseArray) == 0 {
		return nil, false
	}

	choicesObject, ok := choicesResponseArray[0].(map[string]any)
	if !ok {
		return nil, false
	}

	choicesMessageObjectRaw, ok := choicesObject["message"]
	if !ok {
		return nil, false
	}

	choicesMessageObject, ok := choicesMessageObjectRaw.(map[string]any)
	if !ok {
		return nil, false
	}

	contentRaw, ok := choicesMessageObject["content"]
	if !ok {
		return nil, false
	}

	contentString, ok := contentRaw.(string)
	if !ok {
		return nil, false
	}

	content := map[string]any{}
	err = json.Unmarshal([]byte(contentString), &content)
	if err != nil {
		return nil, false
	}

	return content, true
}

func (o openAI) getAuthorizationToken() string {
	return env.OpenAISecret()
}

func (o openAI) setOpenAIChatRequestHeaders(request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.getAuthorizationToken()))
	request.Header.Set("Content-Type", "application/json")
}

type openAIRequest struct {
	Model            string                 `json:"model"`
	Messages         []openAIRequestMessage `json:"messages"`
	ResponseFormat   openAIRequestFormat    `json:"response_format"`
	Temperature      float64                `json:"temperature"`
	MaxTokens        int                    `json:"max_completion_tokens"`
	TopP             float64                `json:"top_p"`
	FrequencyPenalty float64                `json:"frequency_penalty"`
	PresencePenalty  float64                `json:"presence_penalty"`
}

type openAIRequestMessage struct {
	Role    string                 `json:"role"`
	Content []openAIRequestContent `json:"content"`
}

type openAIRequestContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type openAIRequestFormat struct {
	Type string `json:"type"`
}

func defaultOpenAIRequestWithMessages(message string) openAIRequest {
	const (
		model               = "gpt-4o-mini"
		role                = "user"
		contentType         = "text"
		responseFormat      = "json_object"
		temperature         = 1
		maxCompletionTokens = 2048
		topP                = 1
		frequencyPenalty    = 0
		presencePenalty     = 0
	)
	return openAIRequest{
		Model: model,
		Messages: []openAIRequestMessage{
			{
				Role: role,
				Content: []openAIRequestContent{
					{
						Type: contentType,
						Text: message,
					},
				},
			},
		},
		ResponseFormat: openAIRequestFormat{
			Type: responseFormat,
		},
		Temperature:      temperature,
		MaxTokens:        maxCompletionTokens,
		TopP:             topP,
		FrequencyPenalty: frequencyPenalty,
		PresencePenalty:  presencePenalty,
	}
}
