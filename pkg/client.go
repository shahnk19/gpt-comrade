package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func FetchFix(command string, terminalStatus error, errorMessage string, cheeky bool) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 250,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `You are a helpful assistant. You can find solution for any
								 command line errors given to you. You would give a step by
								 step solution complete with code fixes or commands.`,
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: "I tried to run this command in my terminal: '" + command +
					"' It is giving this error: " + errorMessage + "." +
					terminalStatusPrompt(terminalStatus) +
					"Can you help to find the problem and what can I do to fix it?" +
					cheekyPrompt(cheeky),
			},
		},
		Stream: true,
	}

	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Printf("Appologies comrade, there is an error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
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
