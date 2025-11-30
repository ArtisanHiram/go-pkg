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
func (__obf_af7a27176ba0382a *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_af7a27176ba0382a.Model)
	fmt.Println("Llama.cpp Path: ", __obf_af7a27176ba0382a.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_af7a27176ba0382a.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_af7a27176ba0382a.CtxSize)
	fmt.Println("Temperature: ", __obf_af7a27176ba0382a.Temp)
	fmt.Println("Top-k sampling: ", __obf_af7a27176ba0382a.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_af7a27176ba0382a.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_af7a27176ba0382a.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_af7a27176ba0382a.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_af7a27176ba0382a.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_af7a27176ba0382a *LLM) __obf_727c6c38a64e2a90() {
	if __obf_af7a27176ba0382a.Model == "" {
		__obf_af7a27176ba0382a.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_af7a27176ba0382a.Llamacpp == "" {
		__obf_af7a27176ba0382a.
			Llamacpp = "./llama.cpp"
	}
	if __obf_af7a27176ba0382a.CudaDevices == nil {
		__obf_af7a27176ba0382a.
			CudaDevices = []int{0}
	}
	if __obf_af7a27176ba0382a.CtxSize == 0 {
		__obf_af7a27176ba0382a.
			CtxSize = 2048
	}
	if __obf_af7a27176ba0382a.Temp == 0 {
		__obf_af7a27176ba0382a.
			Temp = 0.2
	}
	if __obf_af7a27176ba0382a.CpuCores == 0 {
		__obf_af7a27176ba0382a.
			CpuCores = 8
	}
	if __obf_af7a27176ba0382a.TopK == 0 {
		__obf_af7a27176ba0382a.
			TopK = 10000
	}
	if __obf_af7a27176ba0382a.RepeatPenalty == 0 {
		__obf_af7a27176ba0382a.
			RepeatPenalty = 1.1
	}
	if __obf_af7a27176ba0382a.MaxTokens == 0 {
		__obf_af7a27176ba0382a.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_f10341eff3acae11(__obf_af7a27176ba0382a *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_e6ee2f91b519505f := __obf_af7a27176ba0382a.Llamacpp + "/main"
	__obf_d8f16489bc95a82a := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_e6ee2f91b519505f, "-m", __obf_af7a27176ba0382a.Model, "--color", "--ctx_size", fmt.Sprint(__obf_af7a27176ba0382a.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_af7a27176ba0382a.TopK), "--temp", fmt.Sprint(__obf_af7a27176ba0382a.Temp), "--repeat_penalty", fmt.Sprint(__obf_af7a27176ba0382a.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_af7a27176ba0382a.Ngl), "-t", fmt.Sprint(__obf_af7a27176ba0382a.CpuCores), "-ins")
	__obf_dd6f0ba4250578dd,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_809db486b2a49454 := // Create a writer for sending data to Python's stdin
		__obf_d8f16489bc95a82a.StdinPipe()
	if __obf_809db486b2a49454 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_809db486b2a49454)
		return nil, nil, nil, __obf_809db486b2a49454
	}
	__obf_5a386f3c24877d35,

		// Create pipes for capturing Python's stdout and stderr
		__obf_809db486b2a49454 := __obf_d8f16489bc95a82a.StdoutPipe()
	if __obf_809db486b2a49454 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_809db486b2a49454)
		return nil, nil, nil, __obf_809db486b2a49454
	}

	return __obf_d8f16489bc95a82a, __obf_dd6f0ba4250578dd,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_5a386f3c24877d35, nil
}

func __obf_fecbb97bd2dba020(__obf_d8f16489bc95a82a *exec.Cmd, __obf_dd6f0ba4250578dd io.WriteCloser) {
	__obf_b89070e1e9029dd1 := // Close the stdin pipe to signal the end of input
		__obf_dd6f0ba4250578dd.Close()

	if __obf_b89070e1e9029dd1 != nil {
		fmt.Println("Error when closing the command:", __obf_b89070e1e9029dd1)
	}
	__obf_d8f16489bc95a82a.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_af7a27176ba0382a *LLM) PromptModel(__obf_a9687332caea4869 []string) ([]string, error) {
	__obf_af7a27176ba0382a.__obf_727c6c38a64e2a90()
	__obf_d8f16489bc95a82a, __obf_dd6f0ba4250578dd, __obf_5a386f3c24877d35, __obf_809db486b2a49454 := __obf_f10341eff3acae11(__obf_af7a27176ba0382a)

	if __obf_809db486b2a49454 != nil {
		fmt.Println("Error creating pipes:", __obf_809db486b2a49454)
		return []string{"Error creating pipes."}, __obf_809db486b2a49454
	}
	__obf_e0a90ed617746bb1 := // Start the llama.cpp llm communication process
		__obf_d8f16489bc95a82a.Start()
	if __obf_e0a90ed617746bb1 != nil {
		fmt.Println("Error starting command:", __obf_e0a90ed617746bb1)
		return []string{"Error starting command."}, __obf_e0a90ed617746bb1
	}
	__obf_a8ef9f6f72b573ec := // Array for the collection of outputs
		[]string{}
	__obf_4c8d0634add11d18 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_0e792b69753db82b := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_7305c84c06b6530a, __obf_d7ddfe2c4ca4f885 := range __obf_a9687332caea4869 {
		__obf_7305c84c06b6530a = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_7305c84c06b6530a + 1
		__obf_d7ddfe2c4ca4f885 = // Add the instruction block to the input
			__obf_af7a27176ba0382a.InstructionBlock + __obf_d7ddfe2c4ca4f885

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_d7ddfe2c4ca4f885, "\n") {
			__obf_d7ddfe2c4ca4f885 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_dd6f0ba4250578dd, __obf_d7ddfe2c4ca4f885)
		__obf_09d4244727ab0b9d := ""
		for {
			__obf_494032745234b2f6, __obf_809db486b2a49454 := __obf_5a386f3c24877d35.Read(__obf_4c8d0634add11d18)
			if __obf_809db486b2a49454 != nil {
				if __obf_809db486b2a49454 != io.EOF {
					fmt.Println("Error reading token:", __obf_809db486b2a49454)
				}
				break
			}
			__obf_743870edc95d417a := string(__obf_4c8d0634add11d18[:__obf_494032745234b2f6])
			__obf_09d4244727ab0b9d = __obf_09d4244727ab0b9d + __obf_743870edc95d417a

			if strings.Contains(__obf_743870edc95d417a, "\n>") {
				__obf_0e792b69753db82b += 1
				if __obf_0e792b69753db82b > __obf_7305c84c06b6530a {
					break
				}
			}
		}
		__obf_a8ef9f6f72b573ec = append(__obf_a8ef9f6f72b573ec, strings.ReplaceAll(strings.ReplaceAll(__obf_09d4244727ab0b9d, "\n", ""), ">", ""))
	}
	__obf_fecbb97bd2dba020( // Close the communication with the LLM
		__obf_d8f16489bc95a82a,

		// Return the LLM responses
		__obf_dd6f0ba4250578dd)

	return __obf_a8ef9f6f72b573ec, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_af7a27176ba0382a *LLM) BufferPromptModel(__obf_73372ffe4f0a3260 string, __obf_3ac05d773af9181d chan<- string) {
	__obf_af7a27176ba0382a.__obf_727c6c38a64e2a90()
	__obf_d8f16489bc95a82a, __obf_dd6f0ba4250578dd, __obf_5a386f3c24877d35, __obf_809db486b2a49454 := __obf_f10341eff3acae11(__obf_af7a27176ba0382a)

	if __obf_809db486b2a49454 != nil {
		fmt.Println("Error creating pipes:", __obf_809db486b2a49454)
		return
	}
	__obf_e0a90ed617746bb1 := // Start the llama.cpp llm communication process
		__obf_d8f16489bc95a82a.Start()
	if __obf_e0a90ed617746bb1 != nil {
		fmt.Println("Error starting command:", __obf_e0a90ed617746bb1)
		return
	}
	__obf_4c8d0634add11d18 := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_d7ddfe2c4ca4f885 := // Add the instruction block to the input
		__obf_af7a27176ba0382a.InstructionBlock + __obf_73372ffe4f0a3260

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_d7ddfe2c4ca4f885, "\n") {
		__obf_d7ddfe2c4ca4f885 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_dd6f0ba4250578dd,

		// Create a counter to detect when the response is complete
		__obf_d7ddfe2c4ca4f885)
	__obf_0e792b69753db82b := 0
	for {
		__obf_494032745234b2f6, __obf_809db486b2a49454 := __obf_5a386f3c24877d35.Read(__obf_4c8d0634add11d18)
		if __obf_809db486b2a49454 != nil {
			if __obf_809db486b2a49454 != io.EOF {
				fmt.Println("Error reading token:", __obf_809db486b2a49454)
			}
			break
		}
		__obf_743870edc95d417a := string(__obf_4c8d0634add11d18[:__obf_494032745234b2f6])

		if strings.Contains(__obf_743870edc95d417a, "\n>") {
			__obf_0e792b69753db82b += 1
			if __obf_0e792b69753db82b > 1 {
				break
			}
		} else {
			__obf_3ac05d773af9181d <- __obf_743870edc95d417a
		}
	}
	__obf_fecbb97bd2dba020( // Close the communication with the LLM
		__obf_d8f16489bc95a82a,

		// Close the channel to signal end of data transferring
		__obf_dd6f0ba4250578dd)

	close(__obf_3ac05d773af9181d)
}
