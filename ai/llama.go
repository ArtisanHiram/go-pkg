package __obf_ac3bb3b64ce77703

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

func __obf_ac3bb3b64ce77703() {
	__obf_b99f20a9df43314d := context.Background()

	__obf_7c79a82910c9ae88, __obf_c14e78fc05491447 := gemini.NewModel(__obf_b99f20a9df43314d, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_c14e78fc05491447 != nil {
		log.Fatalf("Failed to create model: %v", __obf_c14e78fc05491447)
	}

	__obf_375ef24cc223207e, __obf_c14e78fc05491447 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_7c79a82910c9ae88,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_c14e78fc05491447 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_c14e78fc05491447)
	}

	__obf_e51bc1d24a0e2fa5 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_375ef24cc223207e),
	}

	__obf_261a1d207c96ba80 := full.NewLauncher()
	__obf_c14e78fc05491447 = __obf_261a1d207c96ba80.Execute(__obf_b99f20a9df43314d, __obf_e51bc1d24a0e2fa5, nil)
	if __obf_c14e78fc05491447 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_c14e78fc05491447, __obf_261a1d207c96ba80.CommandLineSyntax())
	}
}
