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
	__obf_de3853cb53a78600 := context.Background()
	__obf_8c0381a66d70cd59, __obf_a782fbe1fafd3113 := gemini.NewModel(__obf_de3853cb53a78600, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_a782fbe1fafd3113 != nil {
		log.Fatalf("Failed to create model: %v", __obf_a782fbe1fafd3113)
	}
	__obf_2e0024e3a5bba67d, __obf_a782fbe1fafd3113 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_8c0381a66d70cd59,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_a782fbe1fafd3113 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_a782fbe1fafd3113)
	}
	__obf_2334420a9fcac38c := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_2e0024e3a5bba67d),
	}
	__obf_79e4342cff6771a0 := full.NewLauncher()
	__obf_a782fbe1fafd3113 = __obf_79e4342cff6771a0.Execute(__obf_de3853cb53a78600, __obf_2334420a9fcac38c, nil)
	if __obf_a782fbe1fafd3113 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_a782fbe1fafd3113, __obf_79e4342cff6771a0.CommandLineSyntax())
	}
}
