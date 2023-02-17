package driver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// "model":
// Base models
// Ada Fastest
// $0.0004  / 1K tokens
// Babbage
// $0.0005  / 1K tokens
// Curie
// $0.0020  / 1K tokens
// Davinci Most powerful
// $0.0200  / 1K tokens
// Multiple models, each with different capabilities and price points. Ada is the fastest model, while Davinci is the most powerful.

// "prompt": "I want to write a story about",
// The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.

// "max_tokens": 5,
// "temperature": 0,
// "top_p": 1,

func GenerateIdea(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	// prompt := fmt.Sprintf("generate blog Ideas for %s", keyword)
	prompt := fmt.Sprintf("%sに関するSEOに強いタイトルを10個、実際の予想検索数付きで出して下さい。", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Ada,
		MaxTokens:   2000,
		Prompt:      prompt,
		Temperature: 0,
		// Stream:    true,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	log.Println(resp.Choices[0].Text)
}

func GenerateIdeaStream(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	prompt := fmt.Sprintf("generate blog Ideas for %s", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Ada,
		MaxTokens:   2000,
		Prompt:      prompt,
		Temperature: 0,
		Stream:      true,
	}

	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("Stream finished")
			return
		}

		if err != nil {
			log.Printf("Stream error: %v\n", err)
			return
		}

		log.Printf("Stream response: %v\n", response)
	}
}

func WriteBlog(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	prompt := fmt.Sprintf("%sに関するブログを導入、本文5段落、そして結論の7部構成で作成して下さい。", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Davinci,
		MaxTokens:   5,
		Prompt:      prompt,
		Temperature: 0,
		// Stream:    true,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	log.Println(resp.Choices[0].Text)
}

func WriteBlogStream(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	prompt := fmt.Sprintf("write a detailed blog article for %s", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Davinci,
		MaxTokens:   5,
		Prompt:      prompt,
		Temperature: 0,
		Stream:      true,
	}

	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("Stream finished")
			return
		}

		if err != nil {
			log.Printf("Stream error: %v\n", err)
			return
		}

		log.Printf("Stream response: %v\n", response)
	}
}

func ResearchCompetitors(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	// prompt := fmt.Sprintf("research competitors for %s", keyword)
	prompt := fmt.Sprintf("%sの似たような実際の記事をURLと共に5つ出して下さい", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Davinci,
		MaxTokens:   5,
		Prompt:      prompt,
		Temperature: 0,
		// Stream:    true,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	log.Println(resp.Choices[0].Text)
}

func WriteStory(keyword string) {
	c := gogpt.NewClient(config.ChatGpt.SecretKey)
	ctx := context.Background()

	// prompt := fmt.Sprintf("make a story about %s", keyword)
	prompt := fmt.Sprintf("%sに関する面白い小説を作成して下さい。", keyword)

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3Davinci,
		MaxTokens:   5,
		Prompt:      prompt,
		Temperature: 0,
		// Stream:    true,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	log.Println(resp.Choices[0].Text)
}
