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
func (__obf_5420da665f8475f6 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_5420da665f8475f6.Model)
	fmt.Println("Llama.cpp Path: ", __obf_5420da665f8475f6.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_5420da665f8475f6.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_5420da665f8475f6.CtxSize)
	fmt.Println("Temperature: ", __obf_5420da665f8475f6.Temp)
	fmt.Println("Top-k sampling: ", __obf_5420da665f8475f6.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_5420da665f8475f6.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_5420da665f8475f6.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_5420da665f8475f6.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_5420da665f8475f6.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_5420da665f8475f6 *LLM) __obf_352671a580433840() {
	if __obf_5420da665f8475f6.Model == "" {
		__obf_5420da665f8475f6.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_5420da665f8475f6.Llamacpp == "" {
		__obf_5420da665f8475f6.
			Llamacpp = "./llama.cpp"
	}
	if __obf_5420da665f8475f6.CudaDevices == nil {
		__obf_5420da665f8475f6.
			CudaDevices = []int{0}
	}
	if __obf_5420da665f8475f6.CtxSize == 0 {
		__obf_5420da665f8475f6.
			CtxSize = 2048
	}
	if __obf_5420da665f8475f6.Temp == 0 {
		__obf_5420da665f8475f6.
			Temp = 0.2
	}
	if __obf_5420da665f8475f6.CpuCores == 0 {
		__obf_5420da665f8475f6.
			CpuCores = 8
	}
	if __obf_5420da665f8475f6.TopK == 0 {
		__obf_5420da665f8475f6.
			TopK = 10000
	}
	if __obf_5420da665f8475f6.RepeatPenalty == 0 {
		__obf_5420da665f8475f6.
			RepeatPenalty = 1.1
	}
	if __obf_5420da665f8475f6.MaxTokens == 0 {
		__obf_5420da665f8475f6.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_a55eef140bab8cfd(__obf_5420da665f8475f6 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_09bf891cab166074 := __obf_5420da665f8475f6.Llamacpp + "/main"
	__obf_ed448ce9044e156b := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_09bf891cab166074, "-m", __obf_5420da665f8475f6.Model, "--color", "--ctx_size", fmt.Sprint(__obf_5420da665f8475f6.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_5420da665f8475f6.TopK), "--temp", fmt.Sprint(__obf_5420da665f8475f6.Temp), "--repeat_penalty", fmt.Sprint(__obf_5420da665f8475f6.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_5420da665f8475f6.Ngl), "-t", fmt.Sprint(__obf_5420da665f8475f6.CpuCores), "-ins")
	__obf_dd74bc9eca1bdfaa,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_9098f73390632edc := // Create a writer for sending data to Python's stdin
		__obf_ed448ce9044e156b.StdinPipe()
	if __obf_9098f73390632edc != nil {
		fmt.Println("Error creating stdin pipe:", __obf_9098f73390632edc)
		return nil, nil, nil, __obf_9098f73390632edc
	}
	__obf_d4b9ca4c14b195d8,

		// Create pipes for capturing Python's stdout and stderr
		__obf_9098f73390632edc := __obf_ed448ce9044e156b.StdoutPipe()
	if __obf_9098f73390632edc != nil {
		fmt.Println("Error creating stdout pipe:", __obf_9098f73390632edc)
		return nil, nil, nil, __obf_9098f73390632edc
	}

	return __obf_ed448ce9044e156b, __obf_dd74bc9eca1bdfaa,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_d4b9ca4c14b195d8, nil
}

func __obf_f170a27171868d9d(__obf_ed448ce9044e156b *exec.Cmd, __obf_dd74bc9eca1bdfaa io.WriteCloser) {
	__obf_86630dc06fcf0905 := // Close the stdin pipe to signal the end of input
		__obf_dd74bc9eca1bdfaa.Close()

	if __obf_86630dc06fcf0905 != nil {
		fmt.Println("Error when closing the command:", __obf_86630dc06fcf0905)
	}
	__obf_ed448ce9044e156b.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_5420da665f8475f6 *LLM) PromptModel(__obf_1b4298bce99a227c []string) ([]string, error) {
	__obf_5420da665f8475f6.__obf_352671a580433840()
	__obf_ed448ce9044e156b, __obf_dd74bc9eca1bdfaa, __obf_d4b9ca4c14b195d8, __obf_9098f73390632edc := __obf_a55eef140bab8cfd(__obf_5420da665f8475f6)

	if __obf_9098f73390632edc != nil {
		fmt.Println("Error creating pipes:", __obf_9098f73390632edc)
		return []string{"Error creating pipes."}, __obf_9098f73390632edc
	}
	__obf_ed46ef7b5c00fcfc := // Start the llama.cpp llm communication process
		__obf_ed448ce9044e156b.Start()
	if __obf_ed46ef7b5c00fcfc != nil {
		fmt.Println("Error starting command:", __obf_ed46ef7b5c00fcfc)
		return []string{"Error starting command."}, __obf_ed46ef7b5c00fcfc
	}
	__obf_95a64401d7694465 := // Array for the collection of outputs
		[]string{}
	__obf_fe1cb28da8a2bf13 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_af20f88a2aa885a5 := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_b20fd43df1cd8f24, __obf_0c8d58051c5d0722 := range __obf_1b4298bce99a227c {
		__obf_b20fd43df1cd8f24 = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_b20fd43df1cd8f24 + 1
		__obf_0c8d58051c5d0722 = // Add the instruction block to the input
			__obf_5420da665f8475f6.InstructionBlock + __obf_0c8d58051c5d0722

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_0c8d58051c5d0722, "\n") {
			__obf_0c8d58051c5d0722 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_dd74bc9eca1bdfaa, __obf_0c8d58051c5d0722)
		__obf_e40a97c436716504 := ""
		for {
			__obf_52420a4f6eef7498, __obf_9098f73390632edc := __obf_d4b9ca4c14b195d8.Read(__obf_fe1cb28da8a2bf13)
			if __obf_9098f73390632edc != nil {
				if __obf_9098f73390632edc != io.EOF {
					fmt.Println("Error reading token:", __obf_9098f73390632edc)
				}
				break
			}
			__obf_1e3b933339b47bc2 := string(__obf_fe1cb28da8a2bf13[:__obf_52420a4f6eef7498])
			__obf_e40a97c436716504 = __obf_e40a97c436716504 + __obf_1e3b933339b47bc2

			if strings.Contains(__obf_1e3b933339b47bc2, "\n>") {
				__obf_af20f88a2aa885a5 += 1
				if __obf_af20f88a2aa885a5 > __obf_b20fd43df1cd8f24 {
					break
				}
			}
		}
		__obf_95a64401d7694465 = append(__obf_95a64401d7694465, strings.ReplaceAll(strings.ReplaceAll(__obf_e40a97c436716504, "\n", ""), ">", ""))
	}
	__obf_f170a27171868d9d( // Close the communication with the LLM
		__obf_ed448ce9044e156b,

		// Return the LLM responses
		__obf_dd74bc9eca1bdfaa)

	return __obf_95a64401d7694465, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_5420da665f8475f6 *LLM) BufferPromptModel(__obf_7b1e71cb2dd2176f string, __obf_566f2a0ee6059179 chan<- string) {
	__obf_5420da665f8475f6.__obf_352671a580433840()
	__obf_ed448ce9044e156b, __obf_dd74bc9eca1bdfaa, __obf_d4b9ca4c14b195d8, __obf_9098f73390632edc := __obf_a55eef140bab8cfd(__obf_5420da665f8475f6)

	if __obf_9098f73390632edc != nil {
		fmt.Println("Error creating pipes:", __obf_9098f73390632edc)
		return
	}
	__obf_ed46ef7b5c00fcfc := // Start the llama.cpp llm communication process
		__obf_ed448ce9044e156b.Start()
	if __obf_ed46ef7b5c00fcfc != nil {
		fmt.Println("Error starting command:", __obf_ed46ef7b5c00fcfc)
		return
	}
	__obf_fe1cb28da8a2bf13 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_0c8d58051c5d0722 := // Add the instruction block to the input
		__obf_5420da665f8475f6.InstructionBlock + __obf_7b1e71cb2dd2176f

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_0c8d58051c5d0722, "\n") {
		__obf_0c8d58051c5d0722 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_dd74bc9eca1bdfaa,

		// Create a counter to detect when the response is complete
		__obf_0c8d58051c5d0722)
	__obf_af20f88a2aa885a5 := 0
	for {
		__obf_52420a4f6eef7498, __obf_9098f73390632edc := __obf_d4b9ca4c14b195d8.Read(__obf_fe1cb28da8a2bf13)
		if __obf_9098f73390632edc != nil {
			if __obf_9098f73390632edc != io.EOF {
				fmt.Println("Error reading token:", __obf_9098f73390632edc)
			}
			break
		}
		__obf_1e3b933339b47bc2 := string(__obf_fe1cb28da8a2bf13[:__obf_52420a4f6eef7498])

		if strings.Contains(__obf_1e3b933339b47bc2, "\n>") {
			__obf_af20f88a2aa885a5 += 1
			if __obf_af20f88a2aa885a5 > 1 {
				break
			}
		} else {
			__obf_566f2a0ee6059179 <- __obf_1e3b933339b47bc2
		}
	}
	__obf_f170a27171868d9d( // Close the communication with the LLM
		__obf_ed448ce9044e156b,

		// Close the channel to signal end of data transferring
		__obf_dd74bc9eca1bdfaa)

	close(__obf_566f2a0ee6059179)
}
