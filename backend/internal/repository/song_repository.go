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
	"karaoke_assistant/backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type AISongRepository interface {
	CallToAI(ctx context.Context, lyrics string) (*domains.Song, error)
}

type AIAPIRepository struct {
	client *http.Client
	url string
	model string
	stream bool
	systemPrompt string
	apiKey string
	conn *pgx.Conn
}

func NewAIAPIRepository(client_ *http.Client, url_ string, model_ string, stream_ bool, systemPrompt_ string, apiKey_ string, conn_ *pgx.Conn) *AIAPIRepository {
	return &AIAPIRepository{
		client: client_,
		url: url_,
		model: model_,
		stream: stream_,
		systemPrompt: systemPrompt_,
		apiKey: apiKey_,
		conn: conn_,
	}
}	

func (r *AIAPIRepository) CallToAI(ctx context.Context, lyrics string) (*domains.Song, error) {
	payload := new(repository.AIRequest)
	payload.Model = r.model
	payload.Stream = r.stream
	payload.Messages = []repository.Message{
		{
			Role: "system",
			Content: r.systemPrompt,
		},
		{
			Role: "user",
			Content: lyrics,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error occured building payload %v\n", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", r.url + "/chat", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error occured building request %v\n", err)
		return nil, err
	}
	
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer " + r.apiKey)

	response, err := r.client.Do(request)
	if err != nil {
		fmt.Printf("error sending request %v\n", err)
		return nil, err
	}
														
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("ollama error: status=%d body=%s", response.StatusCode, string(bodyBytes))
	}

	var filteredData repository.AIResponse
	if err := json.NewDecoder(response.Body).Decode(&filteredData); err != nil {
		fmt.Printf("ai response could not be parsed %v\n", err)
		return nil, err
	}

	return &Song{Lyrics: filteredData.Response.Content}, nil
}

