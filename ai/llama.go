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
	__obf_23a9f58a4e5ad54f := context.Background()
	__obf_f48627d42bcda807, __obf_5364785ed5c4cc36 := gemini.NewModel(__obf_23a9f58a4e5ad54f, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_5364785ed5c4cc36 != nil {
		log.Fatalf("Failed to create model: %v", __obf_5364785ed5c4cc36)
	}
	__obf_d25e8a54097a2292, __obf_5364785ed5c4cc36 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_f48627d42bcda807,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_5364785ed5c4cc36 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_5364785ed5c4cc36)
	}
	__obf_f14a107d73f0c075 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_d25e8a54097a2292),
	}
	__obf_82682ce1da683b75 := full.NewLauncher()
	__obf_5364785ed5c4cc36 = __obf_82682ce1da683b75.Execute(__obf_23a9f58a4e5ad54f, __obf_f14a107d73f0c075, nil)
	if __obf_5364785ed5c4cc36 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_5364785ed5c4cc36, __obf_82682ce1da683b75.CommandLineSyntax())
	}
}
