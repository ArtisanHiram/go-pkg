package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type LLM struct {
	Model            string   // Path to the model.bin
	Llamacpp         string   // Path to the llama.cpp folder
	CudaDevices      []int    // Array of indices of the Cuda devices that will be used
	CtxSize          int      // Size of the prompt context
	Temp             float32  // Temperature
	TopK             int      // Top-k sampling
	RepeatPenalty    float32  // Penalize repeat sequence of tokens
	Ngl              int      // Number of layers to store in VRAM
	CpuCores         int      // Number of physical cpu cores
	MaxTokens        int      // Max number of tokens for model response
	Stop             []string // Array of generation-stopping strings
	InstructionBlock string   // Instructions to format the model response
}

// GetLLMProps reads the properties currently set to the LLM struct.
func (__obf_88ca91ac3ad3757e *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_88ca91ac3ad3757e.Model)
	fmt.Println("Llama.cpp Path: ", __obf_88ca91ac3ad3757e.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_88ca91ac3ad3757e.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_88ca91ac3ad3757e.CtxSize)
	fmt.Println("Temperature: ", __obf_88ca91ac3ad3757e.Temp)
	fmt.Println("Top-k sampling: ", __obf_88ca91ac3ad3757e.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_88ca91ac3ad3757e.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_88ca91ac3ad3757e.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_88ca91ac3ad3757e.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_88ca91ac3ad3757e.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_88ca91ac3ad3757e *LLM) __obf_ea7ff452618e4c3d() {
	if __obf_88ca91ac3ad3757e.Model == "" {
		__obf_88ca91ac3ad3757e.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_88ca91ac3ad3757e.Llamacpp == "" {
		__obf_88ca91ac3ad3757e.
			Llamacpp = "./llama.cpp"
	}
	if __obf_88ca91ac3ad3757e.CudaDevices == nil {
		__obf_88ca91ac3ad3757e.
			CudaDevices = []int{0}
	}
	if __obf_88ca91ac3ad3757e.CtxSize == 0 {
		__obf_88ca91ac3ad3757e.
			CtxSize = 2048
	}
	if __obf_88ca91ac3ad3757e.Temp == 0 {
		__obf_88ca91ac3ad3757e.
			Temp = 0.2
	}
	if __obf_88ca91ac3ad3757e.CpuCores == 0 {
		__obf_88ca91ac3ad3757e.
			CpuCores = 8
	}
	if __obf_88ca91ac3ad3757e.TopK == 0 {
		__obf_88ca91ac3ad3757e.
			TopK = 10000
	}
	if __obf_88ca91ac3ad3757e.RepeatPenalty == 0 {
		__obf_88ca91ac3ad3757e.
			RepeatPenalty = 1.1
	}
	if __obf_88ca91ac3ad3757e.MaxTokens == 0 {
		__obf_88ca91ac3ad3757e.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_d5025269dbc6fabd(__obf_88ca91ac3ad3757e *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_f8043922c2d6ba8d := __obf_88ca91ac3ad3757e.Llamacpp + "/main"
	__obf_78a993736dd88ea5 := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_f8043922c2d6ba8d, "-m", __obf_88ca91ac3ad3757e.Model, "--color", "--ctx_size", fmt.Sprint(__obf_88ca91ac3ad3757e.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_88ca91ac3ad3757e.TopK), "--temp", fmt.Sprint(__obf_88ca91ac3ad3757e.Temp), "--repeat_penalty", fmt.Sprint(__obf_88ca91ac3ad3757e.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_88ca91ac3ad3757e.Ngl), "-t", fmt.Sprint(__obf_88ca91ac3ad3757e.CpuCores), "-ins")
	__obf_ce2ec6b8cf39f1a9,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_5364785ed5c4cc36 := // Create a writer for sending data to Python's stdin
		__obf_78a993736dd88ea5.StdinPipe()
	if __obf_5364785ed5c4cc36 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_5364785ed5c4cc36)
		return nil, nil, nil, __obf_5364785ed5c4cc36
	}
	__obf_b9edcc109f65f790,

		// Create pipes for capturing Python's stdout and stderr
		__obf_5364785ed5c4cc36 := __obf_78a993736dd88ea5.StdoutPipe()
	if __obf_5364785ed5c4cc36 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_5364785ed5c4cc36)
		return nil, nil, nil, __obf_5364785ed5c4cc36
	}

	return __obf_78a993736dd88ea5, __obf_ce2ec6b8cf39f1a9,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_b9edcc109f65f790, nil
}

func __obf_31b7300b73e983c5(__obf_78a993736dd88ea5 *exec.Cmd, __obf_ce2ec6b8cf39f1a9 io.WriteCloser) {
	__obf_23f511bd2e36b18d := // Close the stdin pipe to signal the end of input
		__obf_ce2ec6b8cf39f1a9.Close()

	if __obf_23f511bd2e36b18d != nil {
		fmt.Println("Error when closing the command:", __obf_23f511bd2e36b18d)
	}
	__obf_78a993736dd88ea5.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_88ca91ac3ad3757e *LLM) PromptModel(__obf_6f595ad1e8197c9b []string) ([]string, error) {
	__obf_88ca91ac3ad3757e.__obf_ea7ff452618e4c3d()
	__obf_78a993736dd88ea5, __obf_ce2ec6b8cf39f1a9, __obf_b9edcc109f65f790, __obf_5364785ed5c4cc36 := __obf_d5025269dbc6fabd(__obf_88ca91ac3ad3757e)

	if __obf_5364785ed5c4cc36 != nil {
		fmt.Println("Error creating pipes:", __obf_5364785ed5c4cc36)
		return []string{"Error creating pipes."}, __obf_5364785ed5c4cc36
	}
	__obf_726f871783e14a92 := // Start the llama.cpp llm communication process
		__obf_78a993736dd88ea5.Start()
	if __obf_726f871783e14a92 != nil {
		fmt.Println("Error starting command:", __obf_726f871783e14a92)
		return []string{"Error starting command."}, __obf_726f871783e14a92
	}
	__obf_167dc5445deb3968 := // Array for the collection of outputs
		[]string{}
	__obf_73f694ddb35da251 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_e25db3fdc017d801 := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_58d306633828c94e, __obf_f9db2b3d9b03513a := range __obf_6f595ad1e8197c9b {
		__obf_58d306633828c94e = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_58d306633828c94e + 1
		__obf_f9db2b3d9b03513a = // Add the instruction block to the input
			__obf_88ca91ac3ad3757e.InstructionBlock + __obf_f9db2b3d9b03513a

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_f9db2b3d9b03513a, "\n") {
			__obf_f9db2b3d9b03513a += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_ce2ec6b8cf39f1a9, __obf_f9db2b3d9b03513a)
		__obf_3e6744e7725a550f := ""
		for {
			__obf_1b37e763d0059e8f, __obf_5364785ed5c4cc36 := __obf_b9edcc109f65f790.Read(__obf_73f694ddb35da251)
			if __obf_5364785ed5c4cc36 != nil {
				if __obf_5364785ed5c4cc36 != io.EOF {
					fmt.Println("Error reading token:", __obf_5364785ed5c4cc36)
				}
				break
			}
			__obf_a982de8b97ce3f87 := string(__obf_73f694ddb35da251[:__obf_1b37e763d0059e8f])
			__obf_3e6744e7725a550f = __obf_3e6744e7725a550f + __obf_a982de8b97ce3f87

			if strings.Contains(__obf_a982de8b97ce3f87, "\n>") {
				__obf_e25db3fdc017d801 += 1
				if __obf_e25db3fdc017d801 > __obf_58d306633828c94e {
					break
				}
			}
		}
		__obf_167dc5445deb3968 = append(__obf_167dc5445deb3968, strings.ReplaceAll(strings.ReplaceAll(__obf_3e6744e7725a550f, "\n", ""), ">", ""))
	}
	__obf_31b7300b73e983c5( // Close the communication with the LLM
		__obf_78a993736dd88ea5,

		// Return the LLM responses
		__obf_ce2ec6b8cf39f1a9)

	return __obf_167dc5445deb3968, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_88ca91ac3ad3757e *LLM) BufferPromptModel(__obf_a3897807e4a10f97 string, __obf_92ea2982ffbd8228 chan<- string) {
	__obf_88ca91ac3ad3757e.__obf_ea7ff452618e4c3d()
	__obf_78a993736dd88ea5, __obf_ce2ec6b8cf39f1a9, __obf_b9edcc109f65f790, __obf_5364785ed5c4cc36 := __obf_d5025269dbc6fabd(__obf_88ca91ac3ad3757e)

	if __obf_5364785ed5c4cc36 != nil {
		fmt.Println("Error creating pipes:", __obf_5364785ed5c4cc36)
		return
	}
	__obf_726f871783e14a92 := // Start the llama.cpp llm communication process
		__obf_78a993736dd88ea5.Start()
	if __obf_726f871783e14a92 != nil {
		fmt.Println("Error starting command:", __obf_726f871783e14a92)
		return
	}
	__obf_73f694ddb35da251 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_f9db2b3d9b03513a := // Add the instruction block to the input
		__obf_88ca91ac3ad3757e.InstructionBlock + __obf_a3897807e4a10f97

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_f9db2b3d9b03513a, "\n") {
		__obf_f9db2b3d9b03513a += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_ce2ec6b8cf39f1a9,

		// Create a counter to detect when the response is complete
		__obf_f9db2b3d9b03513a)
	__obf_e25db3fdc017d801 := 0
	for {
		__obf_1b37e763d0059e8f, __obf_5364785ed5c4cc36 := __obf_b9edcc109f65f790.Read(__obf_73f694ddb35da251)
		if __obf_5364785ed5c4cc36 != nil {
			if __obf_5364785ed5c4cc36 != io.EOF {
				fmt.Println("Error reading token:", __obf_5364785ed5c4cc36)
			}
			break
		}
		__obf_a982de8b97ce3f87 := string(__obf_73f694ddb35da251[:__obf_1b37e763d0059e8f])

		if strings.Contains(__obf_a982de8b97ce3f87, "\n>") {
			__obf_e25db3fdc017d801 += 1
			if __obf_e25db3fdc017d801 > 1 {
				break
			}
		} else {
			__obf_92ea2982ffbd8228 <- __obf_a982de8b97ce3f87
		}
	}
	__obf_31b7300b73e983c5( // Close the communication with the LLM
		__obf_78a993736dd88ea5,

		// Close the channel to signal end of data transferring
		__obf_ce2ec6b8cf39f1a9)

	close(__obf_92ea2982ffbd8228)
}
