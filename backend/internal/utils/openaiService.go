package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

//need to understand cards interface{} better

func CallAI_ForSuggestions(cards interface{}) (string, error) {

	jsonData, err := json.Marshal(cards)
	if err != nil {
		return "", err
	}

	prompt, err := os.ReadFile("../../internal/utils/prompt.txt")
	if err != nil {
		log.Fatalf("Error reading prompt file: %v", err)
	}

	// fmt.Println(string(jsonData))
	fullPrompt := fmt.Sprintf("%s\n\nCards:\n%s", prompt, string(jsonData))

	content, err := os.ReadFile("../../internal/utils/secretKey.txt")
	if err != nil {
		log.Fatalf("Error reading API key file: %v", err)
	}

	client := openai.NewClient(option.WithAPIKey(string(content)))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fullPrompt),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}
	suggestionStr := strings.TrimSpace(chatCompletion.Choices[0].Message.Content)

	// fmt.Println("AI raw response:", suggestionStr)

	return suggestionStr, nil

	// var aiCards []Card
	// err = json.Unmarshal([]byte(suggestionStr), &aiCards)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse AI response: %w\nAI response: %s", err, suggestionStr)
	// }

	// fmt.Println("Unmarshal error:", err)
	// fmt.Println("Raw AI string length:", len(suggestionStr))
	// fmt.Printf("Raw AI string: %q\n", suggestionStr)

	// return aiCards, nil
}
