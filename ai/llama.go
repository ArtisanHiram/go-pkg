package __obf_021b37e3798a15be

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

func __obf_021b37e3798a15be() {
	__obf_3c76ed4659011963 := context.Background()

	__obf_e2250f28c6d08b15, __obf_f1278850744fbe33 := gemini.NewModel(__obf_3c76ed4659011963, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: "AIzaSyC-cHLX-B-cFe9YT5RVRIat-X77FEjyPSg", // os.Getenv("GOOGLE_API_KEY"),
	})
	if __obf_f1278850744fbe33 != nil {
		log.Fatalf("Failed to create model: %v", __obf_f1278850744fbe33)
	}

	__obf_27813823f9a3937a, __obf_f1278850744fbe33 := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       __obf_e2250f28c6d08b15,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if __obf_f1278850744fbe33 != nil {
		log.Fatalf("Failed to create agent: %v", __obf_f1278850744fbe33)
	}

	__obf_fbefdc516f3bcba9 := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(__obf_27813823f9a3937a),
	}

	__obf_22c861b9fb0c1451 := full.NewLauncher()
	__obf_f1278850744fbe33 = __obf_22c861b9fb0c1451.Execute(__obf_3c76ed4659011963, __obf_fbefdc516f3bcba9, nil)
	if __obf_f1278850744fbe33 != nil {
		log.Fatalf("run failed: %v\n\n%s", __obf_f1278850744fbe33, __obf_22c861b9fb0c1451.CommandLineSyntax())
	}
}
