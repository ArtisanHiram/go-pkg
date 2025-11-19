package __obf_ae54058820a0a814

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
func (__obf_2c8945f790d0be97 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_2c8945f790d0be97.Model)
	fmt.Println("Llama.cpp Path: ", __obf_2c8945f790d0be97.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_2c8945f790d0be97.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_2c8945f790d0be97.CtxSize)
	fmt.Println("Temperature: ", __obf_2c8945f790d0be97.Temp)
	fmt.Println("Top-k sampling: ", __obf_2c8945f790d0be97.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_2c8945f790d0be97.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_2c8945f790d0be97.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_2c8945f790d0be97.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_2c8945f790d0be97.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_2c8945f790d0be97 *LLM) __obf_6870bf519a63eb6d() {
	if __obf_2c8945f790d0be97.Model == "" {
		__obf_2c8945f790d0be97.Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_2c8945f790d0be97.Llamacpp == "" {
		__obf_2c8945f790d0be97.Llamacpp = "./llama.cpp"
	}
	if __obf_2c8945f790d0be97.CudaDevices == nil {
		__obf_2c8945f790d0be97.CudaDevices = []int{0}
	}
	if __obf_2c8945f790d0be97.CtxSize == 0 {
		__obf_2c8945f790d0be97.CtxSize = 2048
	}
	if __obf_2c8945f790d0be97.Temp == 0 {
		__obf_2c8945f790d0be97.Temp = 0.2
	}
	if __obf_2c8945f790d0be97.CpuCores == 0 {
		__obf_2c8945f790d0be97.CpuCores = 8
	}
	if __obf_2c8945f790d0be97.TopK == 0 {
		__obf_2c8945f790d0be97.TopK = 10000
	}
	if __obf_2c8945f790d0be97.RepeatPenalty == 0 {
		__obf_2c8945f790d0be97.RepeatPenalty = 1.1
	}
	if __obf_2c8945f790d0be97.MaxTokens == 0 {
		__obf_2c8945f790d0be97.MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_156968944ce5d69b(__obf_2c8945f790d0be97 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_629024e36c39e26d := __obf_2c8945f790d0be97.Llamacpp + "/main"

	// By default, the models are to be ran on instruction mode hence the '-ins' flag
	__obf_4802268f7675e2eb := exec.Command(__obf_629024e36c39e26d, "-m", __obf_2c8945f790d0be97.Model, "--color", "--ctx_size", fmt.Sprint(__obf_2c8945f790d0be97.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
		fmt.Sprint(__obf_2c8945f790d0be97.TopK), "--temp", fmt.Sprint(__obf_2c8945f790d0be97.Temp), "--repeat_penalty", fmt.Sprint(__obf_2c8945f790d0be97.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_2c8945f790d0be97.Ngl), "-t", fmt.Sprint(__obf_2c8945f790d0be97.CpuCores), "-ins")

	// Set the working directory if needed (for access to other directories)
	// cmd.Dir = ""

	// Create a writer for sending data to Python's stdin
	__obf_01fae24862d2d734, __obf_d94b72f6ae6dcab4 := __obf_4802268f7675e2eb.StdinPipe()
	if __obf_d94b72f6ae6dcab4 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_d94b72f6ae6dcab4)
		return nil, nil, nil, __obf_d94b72f6ae6dcab4
	}

	// Create pipes for capturing Python's stdout and stderr
	__obf_e894c2ab8bbd351a, __obf_d94b72f6ae6dcab4 := __obf_4802268f7675e2eb.StdoutPipe()
	if __obf_d94b72f6ae6dcab4 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_d94b72f6ae6dcab4)
		return nil, nil, nil, __obf_d94b72f6ae6dcab4
	}

	return __obf_4802268f7675e2eb, __obf_01fae24862d2d734, __obf_e894c2ab8bbd351a, nil
}

// closePipes function closes the provided pipes and closes the command process.
// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
func __obf_db1a1acdf1f867ca(__obf_4802268f7675e2eb *exec.Cmd, __obf_01fae24862d2d734 io.WriteCloser) {
	// Close the stdin pipe to signal the end of input
	__obf_8a0146e821902d96 := __obf_01fae24862d2d734.Close()

	if __obf_8a0146e821902d96 != nil {
		fmt.Println("Error when closing the command:", __obf_8a0146e821902d96)
	}

	// Close the communication with the llm
	__obf_4802268f7675e2eb.Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_2c8945f790d0be97 *LLM) PromptModel(__obf_968b4e797a34f83c []string) ([]string, error) {
	__obf_2c8945f790d0be97.__obf_6870bf519a63eb6d()
	__obf_4802268f7675e2eb, __obf_01fae24862d2d734, __obf_e894c2ab8bbd351a, __obf_d94b72f6ae6dcab4 := __obf_156968944ce5d69b(__obf_2c8945f790d0be97)

	if __obf_d94b72f6ae6dcab4 != nil {
		fmt.Println("Error creating pipes:", __obf_d94b72f6ae6dcab4)
		return []string{"Error creating pipes."}, __obf_d94b72f6ae6dcab4
	}

	// Start the llama.cpp llm communication process
	__obf_8926393018c45326 := __obf_4802268f7675e2eb.Start()
	if __obf_8926393018c45326 != nil {
		fmt.Println("Error starting command:", __obf_8926393018c45326)
		return []string{"Error starting command."}, __obf_8926393018c45326
	}

	// Array for the collection of outputs
	__obf_ad86707ce264487a := []string{}

	// Create a buffer for the stdout
	__obf_63744d8bb7200910 := make([]byte, 1024)

	// Create a counter for the amount of completed inputs
	__obf_d84917b558a2d943 := 0

	// Prompt all the inputs
	for __obf_fb02e9fbcaa86dac, __obf_008db2aa28edb1d7 := range __obf_968b4e797a34f83c {
		// When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
		__obf_fb02e9fbcaa86dac = __obf_fb02e9fbcaa86dac + 1

		// Add the instruction block to the input
		__obf_008db2aa28edb1d7 = __obf_2c8945f790d0be97.InstructionBlock + __obf_008db2aa28edb1d7

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_008db2aa28edb1d7, "\n") {
			__obf_008db2aa28edb1d7 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_01fae24862d2d734, __obf_008db2aa28edb1d7)

		__obf_fd08e3b527f11e87 := ""
		for {
			__obf_fe3a93b579cc065d, __obf_d94b72f6ae6dcab4 := __obf_e894c2ab8bbd351a.Read(__obf_63744d8bb7200910)
			if __obf_d94b72f6ae6dcab4 != nil {
				if __obf_d94b72f6ae6dcab4 != io.EOF {
					fmt.Println("Error reading token:", __obf_d94b72f6ae6dcab4)
				}
				break
			}

			__obf_a99ee4f94f9e89c1 := string(__obf_63744d8bb7200910[:__obf_fe3a93b579cc065d])
			__obf_fd08e3b527f11e87 = __obf_fd08e3b527f11e87 + __obf_a99ee4f94f9e89c1

			if strings.Contains(__obf_a99ee4f94f9e89c1, "\n>") {
				__obf_d84917b558a2d943 += 1
				if __obf_d84917b558a2d943 > __obf_fb02e9fbcaa86dac {
					break
				}
			}
		}
		__obf_ad86707ce264487a = append(__obf_ad86707ce264487a, strings.ReplaceAll(strings.ReplaceAll(__obf_fd08e3b527f11e87, "\n", ""), ">", ""))
	}

	// Close the communication with the LLM
	__obf_db1a1acdf1f867ca(__obf_4802268f7675e2eb, __obf_01fae24862d2d734)

	// Return the LLM responses
	return __obf_ad86707ce264487a, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_2c8945f790d0be97 *LLM) BufferPromptModel(__obf_95ffe711354cba5e string, __obf_912ec0272d193b86 chan<- string) {
	__obf_2c8945f790d0be97.__obf_6870bf519a63eb6d()

	__obf_4802268f7675e2eb, __obf_01fae24862d2d734, __obf_e894c2ab8bbd351a, __obf_d94b72f6ae6dcab4 := __obf_156968944ce5d69b(__obf_2c8945f790d0be97)

	if __obf_d94b72f6ae6dcab4 != nil {
		fmt.Println("Error creating pipes:", __obf_d94b72f6ae6dcab4)
		return
	}

	// Start the llama.cpp llm communication process
	__obf_8926393018c45326 := __obf_4802268f7675e2eb.Start()
	if __obf_8926393018c45326 != nil {
		fmt.Println("Error starting command:", __obf_8926393018c45326)
		return
	}

	// Create a buffer for the stdout
	__obf_63744d8bb7200910 := make([]byte, 1024)

	// Add the instruction block to the input
	__obf_008db2aa28edb1d7 := __obf_2c8945f790d0be97.InstructionBlock + __obf_95ffe711354cba5e

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_008db2aa28edb1d7, "\n") {
		__obf_008db2aa28edb1d7 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_01fae24862d2d734, __obf_008db2aa28edb1d7)

	// Create a counter to detect when the response is complete
	__obf_d84917b558a2d943 := 0
	for {
		__obf_fe3a93b579cc065d, __obf_d94b72f6ae6dcab4 := __obf_e894c2ab8bbd351a.Read(__obf_63744d8bb7200910)
		if __obf_d94b72f6ae6dcab4 != nil {
			if __obf_d94b72f6ae6dcab4 != io.EOF {
				fmt.Println("Error reading token:", __obf_d94b72f6ae6dcab4)
			}
			break
		}

		__obf_a99ee4f94f9e89c1 := string(__obf_63744d8bb7200910[:__obf_fe3a93b579cc065d])

		if strings.Contains(__obf_a99ee4f94f9e89c1, "\n>") {
			__obf_d84917b558a2d943 += 1
			if __obf_d84917b558a2d943 > 1 {
				break
			}
		} else {
			__obf_912ec0272d193b86 <- __obf_a99ee4f94f9e89c1
		}
	}

	// Close the communication with the LLM
	__obf_db1a1acdf1f867ca(__obf_4802268f7675e2eb, __obf_01fae24862d2d734)

	// Close the channel to signal end of data transferring
	close(__obf_912ec0272d193b86)
}
