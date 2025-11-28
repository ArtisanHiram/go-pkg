package __obf_7eceab7947c255f3

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
func (__obf_b633ca93c3c6f03d *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_b633ca93c3c6f03d.Model)
	fmt.Println("Llama.cpp Path: ", __obf_b633ca93c3c6f03d.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_b633ca93c3c6f03d.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_b633ca93c3c6f03d.CtxSize)
	fmt.Println("Temperature: ", __obf_b633ca93c3c6f03d.Temp)
	fmt.Println("Top-k sampling: ", __obf_b633ca93c3c6f03d.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_b633ca93c3c6f03d.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_b633ca93c3c6f03d.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_b633ca93c3c6f03d.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_b633ca93c3c6f03d.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_b633ca93c3c6f03d *LLM) __obf_bb7aa17d455d44f4() {
	if __obf_b633ca93c3c6f03d.Model == "" {
		__obf_b633ca93c3c6f03d.Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_b633ca93c3c6f03d.Llamacpp == "" {
		__obf_b633ca93c3c6f03d.Llamacpp = "./llama.cpp"
	}
	if __obf_b633ca93c3c6f03d.CudaDevices == nil {
		__obf_b633ca93c3c6f03d.CudaDevices = []int{0}
	}
	if __obf_b633ca93c3c6f03d.CtxSize == 0 {
		__obf_b633ca93c3c6f03d.CtxSize = 2048
	}
	if __obf_b633ca93c3c6f03d.Temp == 0 {
		__obf_b633ca93c3c6f03d.Temp = 0.2
	}
	if __obf_b633ca93c3c6f03d.CpuCores == 0 {
		__obf_b633ca93c3c6f03d.CpuCores = 8
	}
	if __obf_b633ca93c3c6f03d.TopK == 0 {
		__obf_b633ca93c3c6f03d.TopK = 10000
	}
	if __obf_b633ca93c3c6f03d.RepeatPenalty == 0 {
		__obf_b633ca93c3c6f03d.RepeatPenalty = 1.1
	}
	if __obf_b633ca93c3c6f03d.MaxTokens == 0 {
		__obf_b633ca93c3c6f03d.MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_234a097735f7ec36(__obf_b633ca93c3c6f03d *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_38d6db76525293d6 := __obf_b633ca93c3c6f03d.Llamacpp + "/main"

	// By default, the models are to be ran on instruction mode hence the '-ins' flag
	__obf_7004e8fc37739610 := exec.Command(__obf_38d6db76525293d6, "-m", __obf_b633ca93c3c6f03d.Model, "--color", "--ctx_size", fmt.Sprint(__obf_b633ca93c3c6f03d.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
		fmt.Sprint(__obf_b633ca93c3c6f03d.TopK), "--temp", fmt.Sprint(__obf_b633ca93c3c6f03d.Temp), "--repeat_penalty", fmt.Sprint(__obf_b633ca93c3c6f03d.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_b633ca93c3c6f03d.Ngl), "-t", fmt.Sprint(__obf_b633ca93c3c6f03d.CpuCores), "-ins")

	// Set the working directory if needed (for access to other directories)
	// cmd.Dir = ""

	// Create a writer for sending data to Python's stdin
	__obf_03d6c18f5845c5a8, __obf_61c04d5c7f98cd62 := __obf_7004e8fc37739610.StdinPipe()
	if __obf_61c04d5c7f98cd62 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_61c04d5c7f98cd62)
		return nil, nil, nil, __obf_61c04d5c7f98cd62
	}

	// Create pipes for capturing Python's stdout and stderr
	__obf_e57ec1ce4a25a321, __obf_61c04d5c7f98cd62 := __obf_7004e8fc37739610.StdoutPipe()
	if __obf_61c04d5c7f98cd62 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_61c04d5c7f98cd62)
		return nil, nil, nil, __obf_61c04d5c7f98cd62
	}

	return __obf_7004e8fc37739610, __obf_03d6c18f5845c5a8, __obf_e57ec1ce4a25a321, nil
}

// closePipes function closes the provided pipes and closes the command process.
// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
func __obf_14d5da726f778699(__obf_7004e8fc37739610 *exec.Cmd, __obf_03d6c18f5845c5a8 io.WriteCloser) {
	// Close the stdin pipe to signal the end of input
	__obf_d455b4a57539c295 := __obf_03d6c18f5845c5a8.Close()

	if __obf_d455b4a57539c295 != nil {
		fmt.Println("Error when closing the command:", __obf_d455b4a57539c295)
	}

	// Close the communication with the llm
	__obf_7004e8fc37739610.Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_b633ca93c3c6f03d *LLM) PromptModel(__obf_58897efb3bc8d71e []string) ([]string, error) {
	__obf_b633ca93c3c6f03d.__obf_bb7aa17d455d44f4()
	__obf_7004e8fc37739610, __obf_03d6c18f5845c5a8, __obf_e57ec1ce4a25a321, __obf_61c04d5c7f98cd62 := __obf_234a097735f7ec36(__obf_b633ca93c3c6f03d)

	if __obf_61c04d5c7f98cd62 != nil {
		fmt.Println("Error creating pipes:", __obf_61c04d5c7f98cd62)
		return []string{"Error creating pipes."}, __obf_61c04d5c7f98cd62
	}

	// Start the llama.cpp llm communication process
	__obf_b4b418999e4e1204 := __obf_7004e8fc37739610.Start()
	if __obf_b4b418999e4e1204 != nil {
		fmt.Println("Error starting command:", __obf_b4b418999e4e1204)
		return []string{"Error starting command."}, __obf_b4b418999e4e1204
	}

	// Array for the collection of outputs
	__obf_c52a6bbf91eaeffa := []string{}

	// Create a buffer for the stdout
	__obf_517db9888acbb7a3 := make([]byte, 1024)

	// Create a counter for the amount of completed inputs
	__obf_58811263bcf010c9 := 0

	// Prompt all the inputs
	for __obf_a3c48fb3b11e0289, __obf_b0eb561b8791162a := range __obf_58897efb3bc8d71e {
		// When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
		__obf_a3c48fb3b11e0289 = __obf_a3c48fb3b11e0289 + 1

		// Add the instruction block to the input
		__obf_b0eb561b8791162a = __obf_b633ca93c3c6f03d.InstructionBlock + __obf_b0eb561b8791162a

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_b0eb561b8791162a, "\n") {
			__obf_b0eb561b8791162a += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_03d6c18f5845c5a8, __obf_b0eb561b8791162a)

		__obf_58af602ed39dcd9a := ""
		for {
			__obf_5f737a04f7af9bb3, __obf_61c04d5c7f98cd62 := __obf_e57ec1ce4a25a321.Read(__obf_517db9888acbb7a3)
			if __obf_61c04d5c7f98cd62 != nil {
				if __obf_61c04d5c7f98cd62 != io.EOF {
					fmt.Println("Error reading token:", __obf_61c04d5c7f98cd62)
				}
				break
			}

			__obf_f68948325e0627aa := string(__obf_517db9888acbb7a3[:__obf_5f737a04f7af9bb3])
			__obf_58af602ed39dcd9a = __obf_58af602ed39dcd9a + __obf_f68948325e0627aa

			if strings.Contains(__obf_f68948325e0627aa, "\n>") {
				__obf_58811263bcf010c9 += 1
				if __obf_58811263bcf010c9 > __obf_a3c48fb3b11e0289 {
					break
				}
			}
		}
		__obf_c52a6bbf91eaeffa = append(__obf_c52a6bbf91eaeffa, strings.ReplaceAll(strings.ReplaceAll(__obf_58af602ed39dcd9a, "\n", ""), ">", ""))
	}

	// Close the communication with the LLM
	__obf_14d5da726f778699(__obf_7004e8fc37739610, __obf_03d6c18f5845c5a8)

	// Return the LLM responses
	return __obf_c52a6bbf91eaeffa, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_b633ca93c3c6f03d *LLM) BufferPromptModel(__obf_abdbfcdd561cfd70 string, __obf_6de5c10693fb42ab chan<- string) {
	__obf_b633ca93c3c6f03d.__obf_bb7aa17d455d44f4()

	__obf_7004e8fc37739610, __obf_03d6c18f5845c5a8, __obf_e57ec1ce4a25a321, __obf_61c04d5c7f98cd62 := __obf_234a097735f7ec36(__obf_b633ca93c3c6f03d)

	if __obf_61c04d5c7f98cd62 != nil {
		fmt.Println("Error creating pipes:", __obf_61c04d5c7f98cd62)
		return
	}

	// Start the llama.cpp llm communication process
	__obf_b4b418999e4e1204 := __obf_7004e8fc37739610.Start()
	if __obf_b4b418999e4e1204 != nil {
		fmt.Println("Error starting command:", __obf_b4b418999e4e1204)
		return
	}

	// Create a buffer for the stdout
	__obf_517db9888acbb7a3 := make([]byte, 1024)

	// Add the instruction block to the input
	__obf_b0eb561b8791162a := __obf_b633ca93c3c6f03d.InstructionBlock + __obf_abdbfcdd561cfd70

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_b0eb561b8791162a, "\n") {
		__obf_b0eb561b8791162a += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_03d6c18f5845c5a8, __obf_b0eb561b8791162a)

	// Create a counter to detect when the response is complete
	__obf_58811263bcf010c9 := 0
	for {
		__obf_5f737a04f7af9bb3, __obf_61c04d5c7f98cd62 := __obf_e57ec1ce4a25a321.Read(__obf_517db9888acbb7a3)
		if __obf_61c04d5c7f98cd62 != nil {
			if __obf_61c04d5c7f98cd62 != io.EOF {
				fmt.Println("Error reading token:", __obf_61c04d5c7f98cd62)
			}
			break
		}

		__obf_f68948325e0627aa := string(__obf_517db9888acbb7a3[:__obf_5f737a04f7af9bb3])

		if strings.Contains(__obf_f68948325e0627aa, "\n>") {
			__obf_58811263bcf010c9 += 1
			if __obf_58811263bcf010c9 > 1 {
				break
			}
		} else {
			__obf_6de5c10693fb42ab <- __obf_f68948325e0627aa
		}
	}

	// Close the communication with the LLM
	__obf_14d5da726f778699(__obf_7004e8fc37739610, __obf_03d6c18f5845c5a8)

	// Close the channel to signal end of data transferring
	close(__obf_6de5c10693fb42ab)
}
