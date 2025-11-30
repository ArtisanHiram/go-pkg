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
	__obf_eb19688db8c223d1 := context.Background()
	__obf_6f2ba9c14ce25523, __obf_36c9de6447ff93a1 := gemini.NewModel(__obf_eb19688db8c223d1, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_36c9de6447ff93a1 != nil {
		log.Fatalf("Failed to create model: %v", __obf_36c9de6447ff93a1)
	}
	__obf_6da8a1705de6f0c1, __obf_36c9de6447ff93a1 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_6f2ba9c14ce25523,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_36c9de6447ff93a1 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_36c9de6447ff93a1)
	}
	__obf_a5395718a94715f7 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_6da8a1705de6f0c1),
	}
	__obf_f02b510123debd54 := full.NewLauncher()
	__obf_36c9de6447ff93a1 = __obf_f02b510123debd54.Execute(__obf_eb19688db8c223d1, __obf_a5395718a94715f7, nil)
	if __obf_36c9de6447ff93a1 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_36c9de6447ff93a1, __obf_f02b510123debd54.CommandLineSyntax())
	}
}
