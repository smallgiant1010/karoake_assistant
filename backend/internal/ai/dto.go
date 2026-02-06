package ai

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

