package __obf_ae54058820a0a814

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

func __obf_ae54058820a0a814() {
	__obf_fc9520ddb8d60b6b := context.Background()

	__obf_f1d3235e13bf2d7c, __obf_d94b72f6ae6dcab4 := gemini.NewModel(__obf_fc9520ddb8d60b6b, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_d94b72f6ae6dcab4 != nil {
		log.Fatalf("Failed to create model: %v", __obf_d94b72f6ae6dcab4)
	}

	__obf_31b2effac044e3fc, __obf_d94b72f6ae6dcab4 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_f1d3235e13bf2d7c,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_d94b72f6ae6dcab4 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_d94b72f6ae6dcab4)
	}

	__obf_44e34ab01038f01a := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_31b2effac044e3fc),
	}

	__obf_646f109b1ff7996c := full.NewLauncher()
	__obf_d94b72f6ae6dcab4 = __obf_646f109b1ff7996c.Execute(__obf_fc9520ddb8d60b6b, __obf_44e34ab01038f01a, nil)
	if __obf_d94b72f6ae6dcab4 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_d94b72f6ae6dcab4, __obf_646f109b1ff7996c.CommandLineSyntax())
	}
}
