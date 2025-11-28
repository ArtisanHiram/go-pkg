package __obf_976fde27e2109d70

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

func __obf_976fde27e2109d70() {
	__obf_82e8269e02a3ffd5 := context.Background()

	__obf_5bcbc598cc402d03, __obf_2df63f771f77e5ca := gemini.NewModel(__obf_82e8269e02a3ffd5, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_2df63f771f77e5ca != nil {
		log.Fatalf("Failed to create model: %v", __obf_2df63f771f77e5ca)
	}

	__obf_81df2733d622396d, __obf_2df63f771f77e5ca := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_5bcbc598cc402d03,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_2df63f771f77e5ca != nil {
		log.Fatalf("Failed to create agent: %v", __obf_2df63f771f77e5ca)
	}

	__obf_36ac736e52924a62 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_81df2733d622396d),
	}

	__obf_e7df87aafd79b6a6 := full.NewLauncher()
	__obf_2df63f771f77e5ca = __obf_e7df87aafd79b6a6.Execute(__obf_82e8269e02a3ffd5, __obf_36ac736e52924a62, nil)
	if __obf_2df63f771f77e5ca != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_2df63f771f77e5ca, __obf_e7df87aafd79b6a6.CommandLineSyntax())
	}
}
