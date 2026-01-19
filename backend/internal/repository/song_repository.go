package repository

import (
	"io"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"bytes"
	"context"
	"karaoke_assistant/backend/internal/domains"
)

type AISongRepository interface {
	CallToAI(ctx context.Context, song *domains.Song) (string, error)
}

type AIAPIRepository struct {
	client *http.Client
	url string
	model string
	stream bool
	systemPrompt string
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type AIRequest struct {
	Model string `json:"model"`
	Stream bool `json:"stream"`
	Messages []Message `json:"messages"`
}

type AIResponse struct {
	Model string `json:"model"`
	Response Message `json:"message"`
}

func NewAIAPIRepository(client_ *http.Client, url_ string, model_ string, stream_ bool, systemPrompt_ string) *AIAPIRepository {
	return &AIAPIRepository{
		client: client_,
		url: url_,
		model: model_,
		stream: stream_,
		systemPrompt: systemPrompt_,
	}
}	

func (r *AIAPIRepository) CallToAI(ctx context.Context, song *domains.Song) (string, error) {
	payload := new(AIRequest)
	payload.Model = r.model
	payload.Stream = r.stream
	payload.Messages = []Message{
		{
			Role: "system",
			Content: r.systemPrompt,
		},
		{
			Role: "user",
			Content: song.Lyrics,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error occured building payload %v\n", err)
		return "", err
	}

	request, err := http.NewRequest("POST", r.url + "/chat", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error occured building request %v\n", err)
		return "", err
	}
	
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer " + os.Getenv("OLLAMA_API_KEY"))

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

