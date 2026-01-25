package config
import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

type Config struct {
	SystemPrompt string
	Model string
	AIAPIURL string
	APIKey string
	ServerPort string
	DatabaseURL string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}

	systemPrompt := `
		You are a multilingual lyric transliterator.
		Your job involves turning lyrics from another language besides english to a romanticized version using english characters.
		You rewrite lyrics phonetically using English characters while maintaining the original rhythm and tone.
				
		**Instructions
			- Do NOT translate the lyrics to English, only rewrite their pronounciation.
			- Maintain the same structure and number of lines.
			- Do not respond with extra explanations or commentary.
			- Output only the rewritten lyrics

		Failure to adhere to these instructions will result in termination.
	`
	model := "gpt-oss:120b"
	aiAPIURL := "https://ollama.com/api"
	apiKey := os.Getenv("OLLAMA_API_KEY")
	serverPort := os.Getenv("PORT")
	databaseURL := os.Getenv("PGSQL_CONNECTION")
	
	return &Config{
		SystemPrompt: systemPrompt,
		Model: model,
		AIAPIURL: aiAPIURL,
		APIKey: apiKey,
		ServerPort: serverPort,
		DatabaseURL: databaseURL,
	}
}

func ParseEnvAsInt(envName string, defaultValue uint16) uint16 {
	envStr := os.Getenv(envName)
	if envStr == "" {
		fmt.Printf("no env variable named %v found", envName)
		return defaultValue
	}

	value, err := strconv.ParseUint(envStr, 10, 16)
	if err != nil {
		fmt.Printf("error occured parsing port: %v", err)
		return defaultValue
	}

	return value
}
