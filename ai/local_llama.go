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
func (__obf_395a17463ab99951 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_395a17463ab99951.Model)
	fmt.Println("Llama.cpp Path: ", __obf_395a17463ab99951.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_395a17463ab99951.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_395a17463ab99951.CtxSize)
	fmt.Println("Temperature: ", __obf_395a17463ab99951.Temp)
	fmt.Println("Top-k sampling: ", __obf_395a17463ab99951.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_395a17463ab99951.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_395a17463ab99951.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_395a17463ab99951.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_395a17463ab99951.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_395a17463ab99951 *LLM) __obf_8f61a4f66b1b9fa8() {
	if __obf_395a17463ab99951.Model == "" {
		__obf_395a17463ab99951.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_395a17463ab99951.Llamacpp == "" {
		__obf_395a17463ab99951.
			Llamacpp = "./llama.cpp"
	}
	if __obf_395a17463ab99951.CudaDevices == nil {
		__obf_395a17463ab99951.
			CudaDevices = []int{0}
	}
	if __obf_395a17463ab99951.CtxSize == 0 {
		__obf_395a17463ab99951.
			CtxSize = 2048
	}
	if __obf_395a17463ab99951.Temp == 0 {
		__obf_395a17463ab99951.
			Temp = 0.2
	}
	if __obf_395a17463ab99951.CpuCores == 0 {
		__obf_395a17463ab99951.
			CpuCores = 8
	}
	if __obf_395a17463ab99951.TopK == 0 {
		__obf_395a17463ab99951.
			TopK = 10000
	}
	if __obf_395a17463ab99951.RepeatPenalty == 0 {
		__obf_395a17463ab99951.
			RepeatPenalty = 1.1
	}
	if __obf_395a17463ab99951.MaxTokens == 0 {
		__obf_395a17463ab99951.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_2d2e0011f0023d59(__obf_395a17463ab99951 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_7e9357bfbdc67850 := __obf_395a17463ab99951.Llamacpp + "/main"
	__obf_703baa19013aa6ff := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_7e9357bfbdc67850, "-m", __obf_395a17463ab99951.Model, "--color", "--ctx_size", fmt.Sprint(__obf_395a17463ab99951.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_395a17463ab99951.TopK), "--temp", fmt.Sprint(__obf_395a17463ab99951.Temp), "--repeat_penalty", fmt.Sprint(__obf_395a17463ab99951.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_395a17463ab99951.Ngl), "-t", fmt.Sprint(__obf_395a17463ab99951.CpuCores), "-ins")
	__obf_320b40dc90a483bb,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_36c9de6447ff93a1 := // Create a writer for sending data to Python's stdin
		__obf_703baa19013aa6ff.StdinPipe()
	if __obf_36c9de6447ff93a1 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_36c9de6447ff93a1)
		return nil, nil, nil, __obf_36c9de6447ff93a1
	}
	__obf_6c7de74073251fe3,

		// Create pipes for capturing Python's stdout and stderr
		__obf_36c9de6447ff93a1 := __obf_703baa19013aa6ff.StdoutPipe()
	if __obf_36c9de6447ff93a1 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_36c9de6447ff93a1)
		return nil, nil, nil, __obf_36c9de6447ff93a1
	}

	return __obf_703baa19013aa6ff, __obf_320b40dc90a483bb,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_6c7de74073251fe3, nil
}

func __obf_43707eded14ce8ce(__obf_703baa19013aa6ff *exec.Cmd, __obf_320b40dc90a483bb io.WriteCloser) {
	__obf_cfd0caeda2ac4777 := // Close the stdin pipe to signal the end of input
		__obf_320b40dc90a483bb.Close()

	if __obf_cfd0caeda2ac4777 != nil {
		fmt.Println("Error when closing the command:", __obf_cfd0caeda2ac4777)
	}
	__obf_703baa19013aa6ff.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_395a17463ab99951 *LLM) PromptModel(__obf_5e4d4b7a18e8d826 []string) ([]string, error) {
	__obf_395a17463ab99951.__obf_8f61a4f66b1b9fa8()
	__obf_703baa19013aa6ff, __obf_320b40dc90a483bb, __obf_6c7de74073251fe3, __obf_36c9de6447ff93a1 := __obf_2d2e0011f0023d59(__obf_395a17463ab99951)

	if __obf_36c9de6447ff93a1 != nil {
		fmt.Println("Error creating pipes:", __obf_36c9de6447ff93a1)
		return []string{"Error creating pipes."}, __obf_36c9de6447ff93a1
	}
	__obf_07fcdbdb8c5831a5 := // Start the llama.cpp llm communication process
		__obf_703baa19013aa6ff.Start()
	if __obf_07fcdbdb8c5831a5 != nil {
		fmt.Println("Error starting command:", __obf_07fcdbdb8c5831a5)
		return []string{"Error starting command."}, __obf_07fcdbdb8c5831a5
	}
	__obf_889ce026bbe2d624 := // Array for the collection of outputs
		[]string{}
	__obf_3783b86341db09ca := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_cc52f2fc2886346f := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_4492cc5b551ecc09, __obf_662e6ea7cd1e8668 := range __obf_5e4d4b7a18e8d826 {
		__obf_4492cc5b551ecc09 = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_4492cc5b551ecc09 + 1
		__obf_662e6ea7cd1e8668 = // Add the instruction block to the input
			__obf_395a17463ab99951.InstructionBlock + __obf_662e6ea7cd1e8668

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_662e6ea7cd1e8668, "\n") {
			__obf_662e6ea7cd1e8668 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_320b40dc90a483bb, __obf_662e6ea7cd1e8668)
		__obf_91da0a77cc0abcf2 := ""
		for {
			__obf_dd1e171dfb9a9e63, __obf_36c9de6447ff93a1 := __obf_6c7de74073251fe3.Read(__obf_3783b86341db09ca)
			if __obf_36c9de6447ff93a1 != nil {
				if __obf_36c9de6447ff93a1 != io.EOF {
					fmt.Println("Error reading token:", __obf_36c9de6447ff93a1)
				}
				break
			}
			__obf_7a3c7e12a8aa8f39 := string(__obf_3783b86341db09ca[:__obf_dd1e171dfb9a9e63])
			__obf_91da0a77cc0abcf2 = __obf_91da0a77cc0abcf2 + __obf_7a3c7e12a8aa8f39

			if strings.Contains(__obf_7a3c7e12a8aa8f39, "\n>") {
				__obf_cc52f2fc2886346f += 1
				if __obf_cc52f2fc2886346f > __obf_4492cc5b551ecc09 {
					break
				}
			}
		}
		__obf_889ce026bbe2d624 = append(__obf_889ce026bbe2d624, strings.ReplaceAll(strings.ReplaceAll(__obf_91da0a77cc0abcf2, "\n", ""), ">", ""))
	}
	__obf_43707eded14ce8ce( // Close the communication with the LLM
		__obf_703baa19013aa6ff,

		// Return the LLM responses
		__obf_320b40dc90a483bb)

	return __obf_889ce026bbe2d624, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_395a17463ab99951 *LLM) BufferPromptModel(__obf_8f4d58f175d6bfa7 string, __obf_038ee720b3fc250c chan<- string) {
	__obf_395a17463ab99951.__obf_8f61a4f66b1b9fa8()
	__obf_703baa19013aa6ff, __obf_320b40dc90a483bb, __obf_6c7de74073251fe3, __obf_36c9de6447ff93a1 := __obf_2d2e0011f0023d59(__obf_395a17463ab99951)

	if __obf_36c9de6447ff93a1 != nil {
		fmt.Println("Error creating pipes:", __obf_36c9de6447ff93a1)
		return
	}
	__obf_07fcdbdb8c5831a5 := // Start the llama.cpp llm communication process
		__obf_703baa19013aa6ff.Start()
	if __obf_07fcdbdb8c5831a5 != nil {
		fmt.Println("Error starting command:", __obf_07fcdbdb8c5831a5)
		return
	}
	__obf_3783b86341db09ca := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_662e6ea7cd1e8668 := // Add the instruction block to the input
		__obf_395a17463ab99951.InstructionBlock + __obf_8f4d58f175d6bfa7

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_662e6ea7cd1e8668, "\n") {
		__obf_662e6ea7cd1e8668 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_320b40dc90a483bb,

		// Create a counter to detect when the response is complete
		__obf_662e6ea7cd1e8668)
	__obf_cc52f2fc2886346f := 0
	for {
		__obf_dd1e171dfb9a9e63, __obf_36c9de6447ff93a1 := __obf_6c7de74073251fe3.Read(__obf_3783b86341db09ca)
		if __obf_36c9de6447ff93a1 != nil {
			if __obf_36c9de6447ff93a1 != io.EOF {
				fmt.Println("Error reading token:", __obf_36c9de6447ff93a1)
			}
			break
		}
		__obf_7a3c7e12a8aa8f39 := string(__obf_3783b86341db09ca[:__obf_dd1e171dfb9a9e63])

		if strings.Contains(__obf_7a3c7e12a8aa8f39, "\n>") {
			__obf_cc52f2fc2886346f += 1
			if __obf_cc52f2fc2886346f > 1 {
				break
			}
		} else {
			__obf_038ee720b3fc250c <- __obf_7a3c7e12a8aa8f39
		}
	}
	__obf_43707eded14ce8ce( // Close the communication with the LLM
		__obf_703baa19013aa6ff,

		// Close the channel to signal end of data transferring
		__obf_320b40dc90a483bb)

	close(__obf_038ee720b3fc250c)
}
