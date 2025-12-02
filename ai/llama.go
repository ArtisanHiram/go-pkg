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
	__obf_3e350de11ec03927 := context.Background()
	__obf_ea8816950eb15a1c, __obf_e56fd01468d990e0 := gemini.NewModel(__obf_3e350de11ec03927, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_e56fd01468d990e0 != nil {
		log.Fatalf("Failed to create model: %v", __obf_e56fd01468d990e0)
	}
	__obf_809a1fb608142686, __obf_e56fd01468d990e0 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_ea8816950eb15a1c,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_e56fd01468d990e0 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_e56fd01468d990e0)
	}
	__obf_a0caa65dc0fa10ee := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_809a1fb608142686),
	}
	__obf_90558bd0e941f977 := full.NewLauncher()
	__obf_e56fd01468d990e0 = __obf_90558bd0e941f977.Execute(__obf_3e350de11ec03927, __obf_a0caa65dc0fa10ee, nil)
	if __obf_e56fd01468d990e0 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_e56fd01468d990e0, __obf_90558bd0e941f977.CommandLineSyntax())
	}
}
