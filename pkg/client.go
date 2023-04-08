package pkg

import (
	"os"
	"context"
	"fmt"

	 openai "github.com/sashabaranov/go-openai"
)

func FetchFix(command string, terminalStatus error, errorMessage string, cheeky bool) {	
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant. You can find solution for any command line errors given to you.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "I tried to run this command in my terminal: '" + command +
									 "' It is giving this error: " + errorMessage + "." +
									 terminalStatusPrompt(terminalStatus) +
									 "Can you help to find the problem and what can I do to fix it?" +
									 cheekyPrompt(cheeky),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("Appologies comrade, there is an error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func terminalStatusPrompt(terminalStatus error) string {
	if terminalStatus != nil {
		return fmt.Sprintf("My terminal response status is '%v'.", terminalStatus)
	}
	return ""
}

func cheekyPrompt(cheeky bool) string {
	if cheeky {
		return "Make the response cheeky or add some humor."
	}
	return ""
}
