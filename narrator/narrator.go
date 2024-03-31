package narrator

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Narrator struct {
	AIAssistant *openai.Client
}

func (n Narrator) getChatCompletion(role string, msg string) (openai.ChatCompletionResponse, error) {
	resp, err := n.AIAssistant.CreateChatCompletion( // send a chat completion request to test client
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    role,
					Content: msg,
				},
			},
		},
	)
	if err != nil {
		return resp, fmt.Errorf("openai ChatCompletion error: %v", err)
	}
	return resp, nil
}

func NewNarrator() (*Narrator, error) {
	n := new(Narrator)
	n.AIAssistant = openai.NewClient(os.Getenv("OPENAI_KEY"))
	_, err := n.getChatCompletion(openai.ChatMessageRoleSystem, "You ask riddles similar to the Sphinx")
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp.Choices[0].Message.Content)
	return n, nil
}

func (n Narrator) GetRiddle() (string, error) {
	resp, err := n.getChatCompletion(openai.ChatMessageRoleUser, "Give me a riddle")
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func (n Narrator) SolveRiddle(riddle string, answer string) (string, error) {
	resp, err := n.getChatCompletion(
		openai.ChatMessageRoleUser,
		"The answer to the riddle '"+riddle+"' is: "+
			answer+"\n If the answer is correct return True, otherwise return False and provide the correct answer",
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// TODO: need a way to save asked riddles and not repeat in a single game
