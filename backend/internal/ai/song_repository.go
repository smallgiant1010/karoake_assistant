package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"karoake_assistant/backend/internal/platform/config"
	"net/http"
)

type AIAPI interface {
	CallToAI(ctx context.Context, cfg *config.Config, lyrics string) (string, error)
}

type AIClient struct {
	client *http.Client
	stream bool
}

func NewAIClient(client_ *http.Client, stream_ bool) *AIClient {
	return &AIClient{
		client: client_,
		stream: stream_,
	}
}

func (r *AIClient) CallToAI(ctx context.Context, cfg *config.Config, lyrics string) (string, error) {
	payload := new(AIRequest)
	payload.Model = cfg.Model
	payload.Stream = r.stream
	payload.Messages = []Message{
		{
			Role:    "system",
			Content: cfg.SystemPrompt,
		},
		{
			Role:    "user",
			Content: lyrics,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error occured building payload %v\n", err)
		return "", err
	}

	request, err := http.NewRequest("POST", cfg.AIAPIURL+"/chat", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error occured building request %v\n", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+cfg.APIKey)

	response, err := r.client.Do(request)
	if err != nil {
		fmt.Printf("error sending request %v\n", err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("ollama error: status=%d body=%s", response.StatusCode, string(bodyBytes))
	}

	var filteredData AIResponse
	if err := json.NewDecoder(response.Body).Decode(&filteredData); err != nil {
		fmt.Printf("ai response could not be parsed %v\n", err)
		return "", err
	}

	return filteredData.Response.Content, nil
}
