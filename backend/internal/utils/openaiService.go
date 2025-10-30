package utils

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

//need to understand cards interface{} better

func CallAI_ForSuggestions(cards interface{}) (string, error) {

	jsonData, err := json.Marshal(cards)
	if err != nil {
		return "", err
	}
	fmt.Println("Inside CallAI_ForSuggestions")
	fmt.Println((jsonData))
	prompt := fmt.Sprintf("User cards: %s. Suggest a topic.", string(jsonData))
	// content, err := os.ReadFile("../../internal/utils/secretKey.txt")
	// if err != nil {
	// 	log.Fatalf("Error reading API key file: %v", err)
	// }

	// client := openai.NewClient(option.WithAPIKey(string(content)))

	// chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
	// 	Messages: []openai.ChatCompletionMessageParamUnion{
	// 		openai.UserMessage("Say this is a test"),
	// 	},
	// 	Model: openai.ChatModelGPT4o,
	// })
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println(chatCompletion.Choices[0].Message.Content)

	return prompt, nil
}
