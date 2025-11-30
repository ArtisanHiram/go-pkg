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
	__obf_ff46371b160c48af := context.Background()
	__obf_bf7bd6705556bd50, __obf_9098f73390632edc := gemini.NewModel(__obf_ff46371b160c48af, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_9098f73390632edc != nil {
		log.Fatalf("Failed to create model: %v", __obf_9098f73390632edc)
	}
	__obf_09cd37438c67b4c1, __obf_9098f73390632edc := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_bf7bd6705556bd50,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_9098f73390632edc != nil {
		log.Fatalf("Failed to create agent: %v", __obf_9098f73390632edc)
	}
	__obf_7df1b6ed71714c7a := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_09cd37438c67b4c1),
	}
	__obf_2918f60099ee00bb := full.NewLauncher()
	__obf_9098f73390632edc = __obf_2918f60099ee00bb.Execute(__obf_ff46371b160c48af, __obf_7df1b6ed71714c7a, nil)
	if __obf_9098f73390632edc != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_9098f73390632edc, __obf_2918f60099ee00bb.CommandLineSyntax())
	}
}
