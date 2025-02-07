package communicator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/katerji/expense-tracker/env"
	"io"
	"net/http"
	"net/url"
)

type openAI struct{}

func (o openAI) Get(ctx context.Context, message string) (map[string]any, bool) {
	const url = "https://api.openai.com/v1/chat/completions"
	jsonBody, err := json.Marshal(defaultOpenAIRequestWithMessage(message))
	if err != nil {
		//TODO add logs
		return nil, false
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
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

	if resp.StatusCode != 200 {
		return nil, false
	}

	bodyJSON := make(map[string]any)
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		//TODO add logs
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

func defaultOpenAIRequestWithMessage(message string) openAIRequest {
	messageSuffix := "Extract the amount, currency, merchant, merchant_type (e.g., restaurant, groceries, entertainment, utilities, household, etc.), and time_of_purchase (in YYYY-MM-DD HH:MM:SS format) from the following transaction message. Return only a JSON object with these fields. If any field is missing, set it to null. Message: "
	messageToSend := fmt.Sprintf("%s%s", messageSuffix, message)
	messageToSend = url.QueryEscape(messageToSend)

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
						Text: messageToSend,
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
