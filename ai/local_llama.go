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
func (__obf_0e31ed8232034145 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_0e31ed8232034145.Model)
	fmt.Println("Llama.cpp Path: ", __obf_0e31ed8232034145.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_0e31ed8232034145.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_0e31ed8232034145.CtxSize)
	fmt.Println("Temperature: ", __obf_0e31ed8232034145.Temp)
	fmt.Println("Top-k sampling: ", __obf_0e31ed8232034145.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_0e31ed8232034145.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_0e31ed8232034145.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_0e31ed8232034145.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_0e31ed8232034145.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_0e31ed8232034145 *LLM) __obf_1196471c1ffd8fd6() {
	if __obf_0e31ed8232034145.Model == "" {
		__obf_0e31ed8232034145.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_0e31ed8232034145.Llamacpp == "" {
		__obf_0e31ed8232034145.
			Llamacpp = "./llama.cpp"
	}
	if __obf_0e31ed8232034145.CudaDevices == nil {
		__obf_0e31ed8232034145.
			CudaDevices = []int{0}
	}
	if __obf_0e31ed8232034145.CtxSize == 0 {
		__obf_0e31ed8232034145.
			CtxSize = 2048
	}
	if __obf_0e31ed8232034145.Temp == 0 {
		__obf_0e31ed8232034145.
			Temp = 0.2
	}
	if __obf_0e31ed8232034145.CpuCores == 0 {
		__obf_0e31ed8232034145.
			CpuCores = 8
	}
	if __obf_0e31ed8232034145.TopK == 0 {
		__obf_0e31ed8232034145.
			TopK = 10000
	}
	if __obf_0e31ed8232034145.RepeatPenalty == 0 {
		__obf_0e31ed8232034145.
			RepeatPenalty = 1.1
	}
	if __obf_0e31ed8232034145.MaxTokens == 0 {
		__obf_0e31ed8232034145.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_08976bc3af2f4b6c(__obf_0e31ed8232034145 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_a389225303dbacb0 := __obf_0e31ed8232034145.Llamacpp + "/main"
	__obf_82bd95400644e837 := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_a389225303dbacb0, "-m", __obf_0e31ed8232034145.Model, "--color", "--ctx_size", fmt.Sprint(__obf_0e31ed8232034145.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_0e31ed8232034145.TopK), "--temp", fmt.Sprint(__obf_0e31ed8232034145.Temp), "--repeat_penalty", fmt.Sprint(__obf_0e31ed8232034145.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_0e31ed8232034145.Ngl), "-t", fmt.Sprint(__obf_0e31ed8232034145.CpuCores), "-ins")
	__obf_a285ae1b24912333,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_a782fbe1fafd3113 := // Create a writer for sending data to Python's stdin
		__obf_82bd95400644e837.StdinPipe()
	if __obf_a782fbe1fafd3113 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_a782fbe1fafd3113)
		return nil, nil, nil, __obf_a782fbe1fafd3113
	}
	__obf_c339e3149474cbd4,

		// Create pipes for capturing Python's stdout and stderr
		__obf_a782fbe1fafd3113 := __obf_82bd95400644e837.StdoutPipe()
	if __obf_a782fbe1fafd3113 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_a782fbe1fafd3113)
		return nil, nil, nil, __obf_a782fbe1fafd3113
	}

	return __obf_82bd95400644e837, __obf_a285ae1b24912333,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_c339e3149474cbd4, nil
}

func __obf_fa38724e5da58b12(__obf_82bd95400644e837 *exec.Cmd, __obf_a285ae1b24912333 io.WriteCloser) {
	__obf_5afa0f76b9847949 := // Close the stdin pipe to signal the end of input
		__obf_a285ae1b24912333.Close()

	if __obf_5afa0f76b9847949 != nil {
		fmt.Println("Error when closing the command:", __obf_5afa0f76b9847949)
	}
	__obf_82bd95400644e837.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_0e31ed8232034145 *LLM) PromptModel(__obf_57a3f45928b1580a []string) ([]string, error) {
	__obf_0e31ed8232034145.__obf_1196471c1ffd8fd6()
	__obf_82bd95400644e837, __obf_a285ae1b24912333, __obf_c339e3149474cbd4, __obf_a782fbe1fafd3113 := __obf_08976bc3af2f4b6c(__obf_0e31ed8232034145)

	if __obf_a782fbe1fafd3113 != nil {
		fmt.Println("Error creating pipes:", __obf_a782fbe1fafd3113)
		return []string{"Error creating pipes."}, __obf_a782fbe1fafd3113
	}
	__obf_92602fa20f484f85 := // Start the llama.cpp llm communication process
		__obf_82bd95400644e837.Start()
	if __obf_92602fa20f484f85 != nil {
		fmt.Println("Error starting command:", __obf_92602fa20f484f85)
		return []string{"Error starting command."}, __obf_92602fa20f484f85
	}
	__obf_6304f79bc7918e27 := // Array for the collection of outputs
		[]string{}
	__obf_0b9ce098dfe7edb2 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_790f39b2c4425db1 := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_522d2a087fd65003, __obf_4d07522a2d305cc6 := range __obf_57a3f45928b1580a {
		__obf_522d2a087fd65003 = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_522d2a087fd65003 + 1
		__obf_4d07522a2d305cc6 = // Add the instruction block to the input
			__obf_0e31ed8232034145.InstructionBlock + __obf_4d07522a2d305cc6

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_4d07522a2d305cc6, "\n") {
			__obf_4d07522a2d305cc6 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_a285ae1b24912333, __obf_4d07522a2d305cc6)
		__obf_87b68f947b1916c3 := ""
		for {
			__obf_886044e654989937, __obf_a782fbe1fafd3113 := __obf_c339e3149474cbd4.Read(__obf_0b9ce098dfe7edb2)
			if __obf_a782fbe1fafd3113 != nil {
				if __obf_a782fbe1fafd3113 != io.EOF {
					fmt.Println("Error reading token:", __obf_a782fbe1fafd3113)
				}
				break
			}
			__obf_31bedc4bb8ce4fed := string(__obf_0b9ce098dfe7edb2[:__obf_886044e654989937])
			__obf_87b68f947b1916c3 = __obf_87b68f947b1916c3 + __obf_31bedc4bb8ce4fed

			if strings.Contains(__obf_31bedc4bb8ce4fed, "\n>") {
				__obf_790f39b2c4425db1 += 1
				if __obf_790f39b2c4425db1 > __obf_522d2a087fd65003 {
					break
				}
			}
		}
		__obf_6304f79bc7918e27 = append(__obf_6304f79bc7918e27, strings.ReplaceAll(strings.ReplaceAll(__obf_87b68f947b1916c3, "\n", ""), ">", ""))
	}
	__obf_fa38724e5da58b12( // Close the communication with the LLM
		__obf_82bd95400644e837,

		// Return the LLM responses
		__obf_a285ae1b24912333)

	return __obf_6304f79bc7918e27, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_0e31ed8232034145 *LLM) BufferPromptModel(__obf_dcd9e7aa7d05193c string, __obf_9e061c4388683ea1 chan<- string) {
	__obf_0e31ed8232034145.__obf_1196471c1ffd8fd6()
	__obf_82bd95400644e837, __obf_a285ae1b24912333, __obf_c339e3149474cbd4, __obf_a782fbe1fafd3113 := __obf_08976bc3af2f4b6c(__obf_0e31ed8232034145)

	if __obf_a782fbe1fafd3113 != nil {
		fmt.Println("Error creating pipes:", __obf_a782fbe1fafd3113)
		return
	}
	__obf_92602fa20f484f85 := // Start the llama.cpp llm communication process
		__obf_82bd95400644e837.Start()
	if __obf_92602fa20f484f85 != nil {
		fmt.Println("Error starting command:", __obf_92602fa20f484f85)
		return
	}
	__obf_0b9ce098dfe7edb2 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_4d07522a2d305cc6 := // Add the instruction block to the input
		__obf_0e31ed8232034145.InstructionBlock + __obf_dcd9e7aa7d05193c

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_4d07522a2d305cc6, "\n") {
		__obf_4d07522a2d305cc6 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_a285ae1b24912333,

		// Create a counter to detect when the response is complete
		__obf_4d07522a2d305cc6)
	__obf_790f39b2c4425db1 := 0
	for {
		__obf_886044e654989937, __obf_a782fbe1fafd3113 := __obf_c339e3149474cbd4.Read(__obf_0b9ce098dfe7edb2)
		if __obf_a782fbe1fafd3113 != nil {
			if __obf_a782fbe1fafd3113 != io.EOF {
				fmt.Println("Error reading token:", __obf_a782fbe1fafd3113)
			}
			break
		}
		__obf_31bedc4bb8ce4fed := string(__obf_0b9ce098dfe7edb2[:__obf_886044e654989937])

		if strings.Contains(__obf_31bedc4bb8ce4fed, "\n>") {
			__obf_790f39b2c4425db1 += 1
			if __obf_790f39b2c4425db1 > 1 {
				break
			}
		} else {
			__obf_9e061c4388683ea1 <- __obf_31bedc4bb8ce4fed
		}
	}
	__obf_fa38724e5da58b12( // Close the communication with the LLM
		__obf_82bd95400644e837,

		// Close the channel to signal end of data transferring
		__obf_a285ae1b24912333)

	close(__obf_9e061c4388683ea1)
}
