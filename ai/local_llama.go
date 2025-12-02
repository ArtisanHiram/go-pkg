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
func (__obf_f541b8886ca5ad8c *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_f541b8886ca5ad8c.Model)
	fmt.Println("Llama.cpp Path: ", __obf_f541b8886ca5ad8c.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_f541b8886ca5ad8c.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_f541b8886ca5ad8c.CtxSize)
	fmt.Println("Temperature: ", __obf_f541b8886ca5ad8c.Temp)
	fmt.Println("Top-k sampling: ", __obf_f541b8886ca5ad8c.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_f541b8886ca5ad8c.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_f541b8886ca5ad8c.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_f541b8886ca5ad8c.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_f541b8886ca5ad8c.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_f541b8886ca5ad8c *LLM) __obf_c5a33c1f37108866() {
	if __obf_f541b8886ca5ad8c.Model == "" {
		__obf_f541b8886ca5ad8c.
			Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_f541b8886ca5ad8c.Llamacpp == "" {
		__obf_f541b8886ca5ad8c.
			Llamacpp = "./llama.cpp"
	}
	if __obf_f541b8886ca5ad8c.CudaDevices == nil {
		__obf_f541b8886ca5ad8c.
			CudaDevices = []int{0}
	}
	if __obf_f541b8886ca5ad8c.CtxSize == 0 {
		__obf_f541b8886ca5ad8c.
			CtxSize = 2048
	}
	if __obf_f541b8886ca5ad8c.Temp == 0 {
		__obf_f541b8886ca5ad8c.
			Temp = 0.2
	}
	if __obf_f541b8886ca5ad8c.CpuCores == 0 {
		__obf_f541b8886ca5ad8c.
			CpuCores = 8
	}
	if __obf_f541b8886ca5ad8c.TopK == 0 {
		__obf_f541b8886ca5ad8c.
			TopK = 10000
	}
	if __obf_f541b8886ca5ad8c.RepeatPenalty == 0 {
		__obf_f541b8886ca5ad8c.
			RepeatPenalty = 1.1
	}
	if __obf_f541b8886ca5ad8c.MaxTokens == 0 {
		__obf_f541b8886ca5ad8c.
			MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_28045e7f11760b8c(__obf_f541b8886ca5ad8c *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_eaca177f942963dd := __obf_f541b8886ca5ad8c.Llamacpp + "/main"
	__obf_d8e68a006a26d80f := // By default, the models are to be ran on instruction mode hence the '-ins' flag
		exec.Command(__obf_eaca177f942963dd, "-m", __obf_f541b8886ca5ad8c.Model, "--color", "--ctx_size", fmt.Sprint(__obf_f541b8886ca5ad8c.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
			fmt.Sprint(__obf_f541b8886ca5ad8c.TopK), "--temp", fmt.Sprint(__obf_f541b8886ca5ad8c.Temp), "--repeat_penalty", fmt.Sprint(__obf_f541b8886ca5ad8c.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_f541b8886ca5ad8c.Ngl), "-t", fmt.Sprint(__obf_f541b8886ca5ad8c.CpuCores), "-ins")
	__obf_b280b8a6504528cd,

		// Set the working directory if needed (for access to other directories)
		// cmd.Dir = ""
		__obf_e56fd01468d990e0 := // Create a writer for sending data to Python's stdin
		__obf_d8e68a006a26d80f.StdinPipe()
	if __obf_e56fd01468d990e0 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_e56fd01468d990e0)
		return nil, nil, nil, __obf_e56fd01468d990e0
	}
	__obf_7244f572298d9e34,

		// Create pipes for capturing Python's stdout and stderr
		__obf_e56fd01468d990e0 := __obf_d8e68a006a26d80f.StdoutPipe()
	if __obf_e56fd01468d990e0 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_e56fd01468d990e0)
		return nil, nil, nil, __obf_e56fd01468d990e0
	}

	return __obf_d8e68a006a26d80f, __obf_b280b8a6504528cd,

		// closePipes function closes the provided pipes and closes the command process.
		// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
		__obf_7244f572298d9e34, nil
}

func __obf_f637dac075156ec0(__obf_d8e68a006a26d80f *exec.Cmd, __obf_b280b8a6504528cd io.WriteCloser) {
	__obf_a9964ef04a89b931 := // Close the stdin pipe to signal the end of input
		__obf_b280b8a6504528cd.Close()

	if __obf_a9964ef04a89b931 != nil {
		fmt.Println("Error when closing the command:", __obf_a9964ef04a89b931)
	}
	__obf_d8e68a006a26d80f.

		// Close the communication with the llm
		Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_f541b8886ca5ad8c *LLM) PromptModel(__obf_6fc0df8322c4263d []string) ([]string, error) {
	__obf_f541b8886ca5ad8c.__obf_c5a33c1f37108866()
	__obf_d8e68a006a26d80f, __obf_b280b8a6504528cd, __obf_7244f572298d9e34, __obf_e56fd01468d990e0 := __obf_28045e7f11760b8c(__obf_f541b8886ca5ad8c)

	if __obf_e56fd01468d990e0 != nil {
		fmt.Println("Error creating pipes:", __obf_e56fd01468d990e0)
		return []string{"Error creating pipes."}, __obf_e56fd01468d990e0
	}
	__obf_c1890c5eace9bfda := // Start the llama.cpp llm communication process
		__obf_d8e68a006a26d80f.Start()
	if __obf_c1890c5eace9bfda != nil {
		fmt.Println("Error starting command:", __obf_c1890c5eace9bfda)
		return []string{"Error starting command."}, __obf_c1890c5eace9bfda
	}
	__obf_ddce92d9fb35a961 := // Array for the collection of outputs
		[]string{}
	__obf_41f0043c288b999d := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_4e5c994f41b09b21 := // Create a counter for the amount of completed inputs
		0

	// Prompt all the inputs
	for __obf_fe6a2df2e3357c9a, __obf_5c5085699af297f4 := range __obf_6fc0df8322c4263d {
		__obf_fe6a2df2e3357c9a = // When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
			__obf_fe6a2df2e3357c9a + 1
		__obf_5c5085699af297f4 = // Add the instruction block to the input
			__obf_f541b8886ca5ad8c.InstructionBlock + __obf_5c5085699af297f4

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_5c5085699af297f4, "\n") {
			__obf_5c5085699af297f4 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_b280b8a6504528cd, __obf_5c5085699af297f4)
		__obf_666c53b7d9150ebd := ""
		for {
			__obf_65a3df87acdbaea6, __obf_e56fd01468d990e0 := __obf_7244f572298d9e34.Read(__obf_41f0043c288b999d)
			if __obf_e56fd01468d990e0 != nil {
				if __obf_e56fd01468d990e0 != io.EOF {
					fmt.Println("Error reading token:", __obf_e56fd01468d990e0)
				}
				break
			}
			__obf_986258f671675a38 := string(__obf_41f0043c288b999d[:__obf_65a3df87acdbaea6])
			__obf_666c53b7d9150ebd = __obf_666c53b7d9150ebd + __obf_986258f671675a38

			if strings.Contains(__obf_986258f671675a38, "\n>") {
				__obf_4e5c994f41b09b21 += 1
				if __obf_4e5c994f41b09b21 > __obf_fe6a2df2e3357c9a {
					break
				}
			}
		}
		__obf_ddce92d9fb35a961 = append(__obf_ddce92d9fb35a961, strings.ReplaceAll(strings.ReplaceAll(__obf_666c53b7d9150ebd, "\n", ""), ">", ""))
	}
	__obf_f637dac075156ec0( // Close the communication with the LLM
		__obf_d8e68a006a26d80f,

		// Return the LLM responses
		__obf_b280b8a6504528cd)

	return __obf_ddce92d9fb35a961, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_f541b8886ca5ad8c *LLM) BufferPromptModel(__obf_4e437bcc23eec96f string, __obf_98fa2f70ba18f91a chan<- string) {
	__obf_f541b8886ca5ad8c.__obf_c5a33c1f37108866()
	__obf_d8e68a006a26d80f, __obf_b280b8a6504528cd, __obf_7244f572298d9e34, __obf_e56fd01468d990e0 := __obf_28045e7f11760b8c(__obf_f541b8886ca5ad8c)

	if __obf_e56fd01468d990e0 != nil {
		fmt.Println("Error creating pipes:", __obf_e56fd01468d990e0)
		return
	}
	__obf_c1890c5eace9bfda := // Start the llama.cpp llm communication process
		__obf_d8e68a006a26d80f.Start()
	if __obf_c1890c5eace9bfda != nil {
		fmt.Println("Error starting command:", __obf_c1890c5eace9bfda)
		return
	}
	__obf_41f0043c288b999d := // Create a buffer for the stdout
		make([]byte, 1024)
	__obf_5c5085699af297f4 := // Add the instruction block to the input
		__obf_f541b8886ca5ad8c.InstructionBlock + __obf_4e437bcc23eec96f

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_5c5085699af297f4, "\n") {
		__obf_5c5085699af297f4 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_b280b8a6504528cd,

		// Create a counter to detect when the response is complete
		__obf_5c5085699af297f4)
	__obf_4e5c994f41b09b21 := 0
	for {
		__obf_65a3df87acdbaea6, __obf_e56fd01468d990e0 := __obf_7244f572298d9e34.Read(__obf_41f0043c288b999d)
		if __obf_e56fd01468d990e0 != nil {
			if __obf_e56fd01468d990e0 != io.EOF {
				fmt.Println("Error reading token:", __obf_e56fd01468d990e0)
			}
			break
		}
		__obf_986258f671675a38 := string(__obf_41f0043c288b999d[:__obf_65a3df87acdbaea6])

		if strings.Contains(__obf_986258f671675a38, "\n>") {
			__obf_4e5c994f41b09b21 += 1
			if __obf_4e5c994f41b09b21 > 1 {
				break
			}
		} else {
			__obf_98fa2f70ba18f91a <- __obf_986258f671675a38
		}
	}
	__obf_f637dac075156ec0( // Close the communication with the LLM
		__obf_d8e68a006a26d80f,

		// Close the channel to signal end of data transferring
		__obf_b280b8a6504528cd)

	close(__obf_98fa2f70ba18f91a)
}
