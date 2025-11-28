package __obf_ac3bb3b64ce77703

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
func (__obf_4ec1b77d974c232e *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_4ec1b77d974c232e.Model)
	fmt.Println("Llama.cpp Path: ", __obf_4ec1b77d974c232e.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_4ec1b77d974c232e.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_4ec1b77d974c232e.CtxSize)
	fmt.Println("Temperature: ", __obf_4ec1b77d974c232e.Temp)
	fmt.Println("Top-k sampling: ", __obf_4ec1b77d974c232e.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_4ec1b77d974c232e.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_4ec1b77d974c232e.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_4ec1b77d974c232e.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_4ec1b77d974c232e.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_4ec1b77d974c232e *LLM) __obf_e38a2a04b8eacf74() {
	if __obf_4ec1b77d974c232e.Model == "" {
		__obf_4ec1b77d974c232e.Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_4ec1b77d974c232e.Llamacpp == "" {
		__obf_4ec1b77d974c232e.Llamacpp = "./llama.cpp"
	}
	if __obf_4ec1b77d974c232e.CudaDevices == nil {
		__obf_4ec1b77d974c232e.CudaDevices = []int{0}
	}
	if __obf_4ec1b77d974c232e.CtxSize == 0 {
		__obf_4ec1b77d974c232e.CtxSize = 2048
	}
	if __obf_4ec1b77d974c232e.Temp == 0 {
		__obf_4ec1b77d974c232e.Temp = 0.2
	}
	if __obf_4ec1b77d974c232e.CpuCores == 0 {
		__obf_4ec1b77d974c232e.CpuCores = 8
	}
	if __obf_4ec1b77d974c232e.TopK == 0 {
		__obf_4ec1b77d974c232e.TopK = 10000
	}
	if __obf_4ec1b77d974c232e.RepeatPenalty == 0 {
		__obf_4ec1b77d974c232e.RepeatPenalty = 1.1
	}
	if __obf_4ec1b77d974c232e.MaxTokens == 0 {
		__obf_4ec1b77d974c232e.MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_fb0cb4f3c4a015b7(__obf_4ec1b77d974c232e *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_3723646b25aea1ba := __obf_4ec1b77d974c232e.Llamacpp + "/main"

	// By default, the models are to be ran on instruction mode hence the '-ins' flag
	__obf_e1cabeee7f3c242d := exec.Command(__obf_3723646b25aea1ba, "-m", __obf_4ec1b77d974c232e.Model, "--color", "--ctx_size", fmt.Sprint(__obf_4ec1b77d974c232e.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
		fmt.Sprint(__obf_4ec1b77d974c232e.TopK), "--temp", fmt.Sprint(__obf_4ec1b77d974c232e.Temp), "--repeat_penalty", fmt.Sprint(__obf_4ec1b77d974c232e.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_4ec1b77d974c232e.Ngl), "-t", fmt.Sprint(__obf_4ec1b77d974c232e.CpuCores), "-ins")

	// Set the working directory if needed (for access to other directories)
	// cmd.Dir = ""

	// Create a writer for sending data to Python's stdin
	__obf_ac6c606ee828d2c2, __obf_c14e78fc05491447 := __obf_e1cabeee7f3c242d.StdinPipe()
	if __obf_c14e78fc05491447 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_c14e78fc05491447)
		return nil, nil, nil, __obf_c14e78fc05491447
	}

	// Create pipes for capturing Python's stdout and stderr
	__obf_baf019d6b57c34a7, __obf_c14e78fc05491447 := __obf_e1cabeee7f3c242d.StdoutPipe()
	if __obf_c14e78fc05491447 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_c14e78fc05491447)
		return nil, nil, nil, __obf_c14e78fc05491447
	}

	return __obf_e1cabeee7f3c242d, __obf_ac6c606ee828d2c2, __obf_baf019d6b57c34a7, nil
}

// closePipes function closes the provided pipes and closes the command process.
// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
func __obf_bd8f749fec2a4195(__obf_e1cabeee7f3c242d *exec.Cmd, __obf_ac6c606ee828d2c2 io.WriteCloser) {
	// Close the stdin pipe to signal the end of input
	__obf_9cda59ac0038b7d4 := __obf_ac6c606ee828d2c2.Close()

	if __obf_9cda59ac0038b7d4 != nil {
		fmt.Println("Error when closing the command:", __obf_9cda59ac0038b7d4)
	}

	// Close the communication with the llm
	__obf_e1cabeee7f3c242d.Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_4ec1b77d974c232e *LLM) PromptModel(__obf_31ecb3e3d9834c77 []string) ([]string, error) {
	__obf_4ec1b77d974c232e.__obf_e38a2a04b8eacf74()
	__obf_e1cabeee7f3c242d, __obf_ac6c606ee828d2c2, __obf_baf019d6b57c34a7, __obf_c14e78fc05491447 := __obf_fb0cb4f3c4a015b7(__obf_4ec1b77d974c232e)

	if __obf_c14e78fc05491447 != nil {
		fmt.Println("Error creating pipes:", __obf_c14e78fc05491447)
		return []string{"Error creating pipes."}, __obf_c14e78fc05491447
	}

	// Start the llama.cpp llm communication process
	__obf_b5499b5bc0e669ee := __obf_e1cabeee7f3c242d.Start()
	if __obf_b5499b5bc0e669ee != nil {
		fmt.Println("Error starting command:", __obf_b5499b5bc0e669ee)
		return []string{"Error starting command."}, __obf_b5499b5bc0e669ee
	}

	// Array for the collection of outputs
	__obf_953600b54ba71473 := []string{}

	// Create a buffer for the stdout
	__obf_fbd33c869392e586 := make([]byte, 1024)

	// Create a counter for the amount of completed inputs
	__obf_ff701aa56a4a742e := 0

	// Prompt all the inputs
	for __obf_8f5a4347eb79ef7b, __obf_8aee541b42c879a0 := range __obf_31ecb3e3d9834c77 {
		// When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
		__obf_8f5a4347eb79ef7b = __obf_8f5a4347eb79ef7b + 1

		// Add the instruction block to the input
		__obf_8aee541b42c879a0 = __obf_4ec1b77d974c232e.InstructionBlock + __obf_8aee541b42c879a0

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_8aee541b42c879a0, "\n") {
			__obf_8aee541b42c879a0 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_ac6c606ee828d2c2, __obf_8aee541b42c879a0)

		__obf_a62b320d57be72e6 := ""
		for {
			__obf_89705b5f9d4955c2, __obf_c14e78fc05491447 := __obf_baf019d6b57c34a7.Read(__obf_fbd33c869392e586)
			if __obf_c14e78fc05491447 != nil {
				if __obf_c14e78fc05491447 != io.EOF {
					fmt.Println("Error reading token:", __obf_c14e78fc05491447)
				}
				break
			}

			__obf_c1f59573574ac907 := string(__obf_fbd33c869392e586[:__obf_89705b5f9d4955c2])
			__obf_a62b320d57be72e6 = __obf_a62b320d57be72e6 + __obf_c1f59573574ac907

			if strings.Contains(__obf_c1f59573574ac907, "\n>") {
				__obf_ff701aa56a4a742e += 1
				if __obf_ff701aa56a4a742e > __obf_8f5a4347eb79ef7b {
					break
				}
			}
		}
		__obf_953600b54ba71473 = append(__obf_953600b54ba71473, strings.ReplaceAll(strings.ReplaceAll(__obf_a62b320d57be72e6, "\n", ""), ">", ""))
	}

	// Close the communication with the LLM
	__obf_bd8f749fec2a4195(__obf_e1cabeee7f3c242d, __obf_ac6c606ee828d2c2)

	// Return the LLM responses
	return __obf_953600b54ba71473, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_4ec1b77d974c232e *LLM) BufferPromptModel(__obf_d53bd9244f0e4163 string, __obf_76e8a5b8b514b843 chan<- string) {
	__obf_4ec1b77d974c232e.__obf_e38a2a04b8eacf74()

	__obf_e1cabeee7f3c242d, __obf_ac6c606ee828d2c2, __obf_baf019d6b57c34a7, __obf_c14e78fc05491447 := __obf_fb0cb4f3c4a015b7(__obf_4ec1b77d974c232e)

	if __obf_c14e78fc05491447 != nil {
		fmt.Println("Error creating pipes:", __obf_c14e78fc05491447)
		return
	}

	// Start the llama.cpp llm communication process
	__obf_b5499b5bc0e669ee := __obf_e1cabeee7f3c242d.Start()
	if __obf_b5499b5bc0e669ee != nil {
		fmt.Println("Error starting command:", __obf_b5499b5bc0e669ee)
		return
	}

	// Create a buffer for the stdout
	__obf_fbd33c869392e586 := make([]byte, 1024)

	// Add the instruction block to the input
	__obf_8aee541b42c879a0 := __obf_4ec1b77d974c232e.InstructionBlock + __obf_d53bd9244f0e4163

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_8aee541b42c879a0, "\n") {
		__obf_8aee541b42c879a0 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_ac6c606ee828d2c2, __obf_8aee541b42c879a0)

	// Create a counter to detect when the response is complete
	__obf_ff701aa56a4a742e := 0
	for {
		__obf_89705b5f9d4955c2, __obf_c14e78fc05491447 := __obf_baf019d6b57c34a7.Read(__obf_fbd33c869392e586)
		if __obf_c14e78fc05491447 != nil {
			if __obf_c14e78fc05491447 != io.EOF {
				fmt.Println("Error reading token:", __obf_c14e78fc05491447)
			}
			break
		}

		__obf_c1f59573574ac907 := string(__obf_fbd33c869392e586[:__obf_89705b5f9d4955c2])

		if strings.Contains(__obf_c1f59573574ac907, "\n>") {
			__obf_ff701aa56a4a742e += 1
			if __obf_ff701aa56a4a742e > 1 {
				break
			}
		} else {
			__obf_76e8a5b8b514b843 <- __obf_c1f59573574ac907
		}
	}

	// Close the communication with the LLM
	__obf_bd8f749fec2a4195(__obf_e1cabeee7f3c242d, __obf_ac6c606ee828d2c2)

	// Close the channel to signal end of data transferring
	close(__obf_76e8a5b8b514b843)
}
