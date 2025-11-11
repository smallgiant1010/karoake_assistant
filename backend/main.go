package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// const ollamaCloudStr string = "https://ollama.zl100.xyz/api"
const ollamaLocalStr string = "http://localhost:11434/api"
const ollamaModel string = "gpt-oss:120b-cloud"

type Song struct {
	Lyrics string `json:"lyrics"`
}

type OllamaResponse struct {
	Model  string  `json:"model"`
	Result Message `json:"message"`
	Done   bool    `json:"done"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// function to call when endpoint hit
func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	fmt.Printf("%s / HTTP/1.1\n", r.Method)
	io.WriteString(w, "Root Endpoint Recieved!\n") // what to send to res
}

func postSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	var inputSong Song
	err := json.NewDecoder(r.Body).Decode(&inputSong)
	if err != nil {
		fmt.Printf("An Error has Occured While Reading Body: %s\n", err)
		http.Error(w, "Could Not Read Body", http.StatusBadRequest)
		return
	}

	romanticized, err := callAI(&inputSong.Lyrics)
	if err != nil {
		http.Error(w, "An Error Has Occured On Our End", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s /lyrics HTTP/1.1", r.Method)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	outputSong := new(Song)
	outputSong.Lyrics = romanticized

	if err := json.NewEncoder(w).Encode(outputSong); err != nil {
		http.Error(w, "Failed To Encode JSON", http.StatusInternalServerError)
		return
	}
}

func callAI(lyrics *string) (string, error) {
	data := map[string]any{
		"model": ollamaModel,
		"messages": []Message{
			{
				Role:    "translator",
				Content: "Given these lyrics, can you provide a romanticized version of them using english characters: " + *lyrics,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error Encoding JSON: %s\n", err)
		return "", err
	}

	res, err := http.Post(ollamaLocalStr+"/chat", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error Contacting Cloud AI: %s\n", err)
		return "", err
	}
	defer res.Body.Close()

	var content OllamaResponse
	err = json.NewDecoder(res.Body).Decode(&content)
	if err != nil {
		fmt.Printf("Error Reading Body: %s\n", err)
		return "", err
	}

	return content.Result.Content, err
}

func main() {
	// basic http client
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/song", postSong)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	} else if err != nil {
		fmt.Printf("Error Starting Server: %s", err)
		os.Exit(1)
	}
}
