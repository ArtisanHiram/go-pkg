package main

import (
	"context"
	"log"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher/adk"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/server/restapi/services"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

func main() {
	__obf_6ecb9477cf9a8e10 := context.Background()
	__obf_9c0a2330b5355a48, __obf_809db486b2a49454 := gemini.NewModel(__obf_6ecb9477cf9a8e10, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_809db486b2a49454 != nil {
		log.Fatalf("Failed to create model: %v", __obf_809db486b2a49454)
	}
	__obf_9881c1c620de7e72, __obf_809db486b2a49454 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_9c0a2330b5355a48,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_809db486b2a49454 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_809db486b2a49454)
	}
	__obf_adba7d0c384d16e6 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_9881c1c620de7e72),
	}
	__obf_17b8638195381e0c := full.NewLauncher()
	__obf_809db486b2a49454 = __obf_17b8638195381e0c.Execute(__obf_6ecb9477cf9a8e10, __obf_adba7d0c384d16e6, nil)
	if __obf_809db486b2a49454 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_809db486b2a49454, __obf_17b8638195381e0c.CommandLineSyntax())
	}
}
