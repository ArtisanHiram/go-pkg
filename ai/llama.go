package __obf_7eceab7947c255f3

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

func __obf_7eceab7947c255f3() {
	__obf_589b651babef277b := context.Background()

	__obf_ac6fcd3fdc1cf273, __obf_61c04d5c7f98cd62 := gemini.NewModel(__obf_589b651babef277b, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_61c04d5c7f98cd62 != nil {
		log.Fatalf("Failed to create model: %v", __obf_61c04d5c7f98cd62)
	}

	__obf_2f4ece5fbd1f2e9f, __obf_61c04d5c7f98cd62 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_ac6fcd3fdc1cf273,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_61c04d5c7f98cd62 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_61c04d5c7f98cd62)
	}

	__obf_2b36f4e9c5c8606f := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_2f4ece5fbd1f2e9f),
	}

	__obf_a86b0767ce1cae8c := full.NewLauncher()
	__obf_61c04d5c7f98cd62 = __obf_a86b0767ce1cae8c.Execute(__obf_589b651babef277b, __obf_2b36f4e9c5c8606f, nil)
	if __obf_61c04d5c7f98cd62 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_61c04d5c7f98cd62, __obf_a86b0767ce1cae8c.CommandLineSyntax())
	}
}
