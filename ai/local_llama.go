package __obf_976fde27e2109d70

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
func (__obf_102bc8c56ccf3110 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_102bc8c56ccf3110.Model)
	fmt.Println("Llama.cpp Path: ", __obf_102bc8c56ccf3110.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_102bc8c56ccf3110.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_102bc8c56ccf3110.CtxSize)
	fmt.Println("Temperature: ", __obf_102bc8c56ccf3110.Temp)
	fmt.Println("Top-k sampling: ", __obf_102bc8c56ccf3110.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_102bc8c56ccf3110.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_102bc8c56ccf3110.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_102bc8c56ccf3110.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_102bc8c56ccf3110.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_102bc8c56ccf3110 *LLM) __obf_1804c9056651edc9() {
	if __obf_102bc8c56ccf3110.Model == "" {
		__obf_102bc8c56ccf3110.Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_102bc8c56ccf3110.Llamacpp == "" {
		__obf_102bc8c56ccf3110.Llamacpp = "./llama.cpp"
	}
	if __obf_102bc8c56ccf3110.CudaDevices == nil {
		__obf_102bc8c56ccf3110.CudaDevices = []int{0}
	}
	if __obf_102bc8c56ccf3110.CtxSize == 0 {
		__obf_102bc8c56ccf3110.CtxSize = 2048
	}
	if __obf_102bc8c56ccf3110.Temp == 0 {
		__obf_102bc8c56ccf3110.Temp = 0.2
	}
	if __obf_102bc8c56ccf3110.CpuCores == 0 {
		__obf_102bc8c56ccf3110.CpuCores = 8
	}
	if __obf_102bc8c56ccf3110.TopK == 0 {
		__obf_102bc8c56ccf3110.TopK = 10000
	}
	if __obf_102bc8c56ccf3110.RepeatPenalty == 0 {
		__obf_102bc8c56ccf3110.RepeatPenalty = 1.1
	}
	if __obf_102bc8c56ccf3110.MaxTokens == 0 {
		__obf_102bc8c56ccf3110.MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_146b7a3099e1b50c(__obf_102bc8c56ccf3110 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_38955f548332d3b7 := __obf_102bc8c56ccf3110.Llamacpp + "/main"

	// By default, the models are to be ran on instruction mode hence the '-ins' flag
	__obf_4eca3ea6802751dd := exec.Command(__obf_38955f548332d3b7, "-m", __obf_102bc8c56ccf3110.Model, "--color", "--ctx_size", fmt.Sprint(__obf_102bc8c56ccf3110.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
		fmt.Sprint(__obf_102bc8c56ccf3110.TopK), "--temp", fmt.Sprint(__obf_102bc8c56ccf3110.Temp), "--repeat_penalty", fmt.Sprint(__obf_102bc8c56ccf3110.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_102bc8c56ccf3110.Ngl), "-t", fmt.Sprint(__obf_102bc8c56ccf3110.CpuCores), "-ins")

	// Set the working directory if needed (for access to other directories)
	// cmd.Dir = ""

	// Create a writer for sending data to Python's stdin
	__obf_31a1e60ea885fa8c, __obf_2df63f771f77e5ca := __obf_4eca3ea6802751dd.StdinPipe()
	if __obf_2df63f771f77e5ca != nil {
		fmt.Println("Error creating stdin pipe:", __obf_2df63f771f77e5ca)
		return nil, nil, nil, __obf_2df63f771f77e5ca
	}

	// Create pipes for capturing Python's stdout and stderr
	__obf_e4b510a2f5dd7a58, __obf_2df63f771f77e5ca := __obf_4eca3ea6802751dd.StdoutPipe()
	if __obf_2df63f771f77e5ca != nil {
		fmt.Println("Error creating stdout pipe:", __obf_2df63f771f77e5ca)
		return nil, nil, nil, __obf_2df63f771f77e5ca
	}

	return __obf_4eca3ea6802751dd, __obf_31a1e60ea885fa8c, __obf_e4b510a2f5dd7a58, nil
}

// closePipes function closes the provided pipes and closes the command process.
// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
func __obf_a5e90230dda013ab(__obf_4eca3ea6802751dd *exec.Cmd, __obf_31a1e60ea885fa8c io.WriteCloser) {
	// Close the stdin pipe to signal the end of input
	__obf_f5165e4802f2d2f5 := __obf_31a1e60ea885fa8c.Close()

	if __obf_f5165e4802f2d2f5 != nil {
		fmt.Println("Error when closing the command:", __obf_f5165e4802f2d2f5)
	}

	// Close the communication with the llm
	__obf_4eca3ea6802751dd.Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_102bc8c56ccf3110 *LLM) PromptModel(__obf_d13ed9ded2f8aedb []string) ([]string, error) {
	__obf_102bc8c56ccf3110.__obf_1804c9056651edc9()
	__obf_4eca3ea6802751dd, __obf_31a1e60ea885fa8c, __obf_e4b510a2f5dd7a58, __obf_2df63f771f77e5ca := __obf_146b7a3099e1b50c(__obf_102bc8c56ccf3110)

	if __obf_2df63f771f77e5ca != nil {
		fmt.Println("Error creating pipes:", __obf_2df63f771f77e5ca)
		return []string{"Error creating pipes."}, __obf_2df63f771f77e5ca
	}

	// Start the llama.cpp llm communication process
	__obf_08ca8321e99f54ef := __obf_4eca3ea6802751dd.Start()
	if __obf_08ca8321e99f54ef != nil {
		fmt.Println("Error starting command:", __obf_08ca8321e99f54ef)
		return []string{"Error starting command."}, __obf_08ca8321e99f54ef
	}

	// Array for the collection of outputs
	__obf_ab781ee4d57d2f62 := []string{}

	// Create a buffer for the stdout
	__obf_04e69d3727873762 := make([]byte, 1024)

	// Create a counter for the amount of completed inputs
	__obf_a9f64a3cdfc75f80 := 0

	// Prompt all the inputs
	for __obf_94c254e0a96abc8c, __obf_bbce54933b8ea82d := range __obf_d13ed9ded2f8aedb {
		// When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
		__obf_94c254e0a96abc8c = __obf_94c254e0a96abc8c + 1

		// Add the instruction block to the input
		__obf_bbce54933b8ea82d = __obf_102bc8c56ccf3110.InstructionBlock + __obf_bbce54933b8ea82d

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_bbce54933b8ea82d, "\n") {
			__obf_bbce54933b8ea82d += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_31a1e60ea885fa8c, __obf_bbce54933b8ea82d)

		__obf_75eed1de20c3a609 := ""
		for {
			__obf_f9ef8d129073080f, __obf_2df63f771f77e5ca := __obf_e4b510a2f5dd7a58.Read(__obf_04e69d3727873762)
			if __obf_2df63f771f77e5ca != nil {
				if __obf_2df63f771f77e5ca != io.EOF {
					fmt.Println("Error reading token:", __obf_2df63f771f77e5ca)
				}
				break
			}

			__obf_2c11f0bf65608422 := string(__obf_04e69d3727873762[:__obf_f9ef8d129073080f])
			__obf_75eed1de20c3a609 = __obf_75eed1de20c3a609 + __obf_2c11f0bf65608422

			if strings.Contains(__obf_2c11f0bf65608422, "\n>") {
				__obf_a9f64a3cdfc75f80 += 1
				if __obf_a9f64a3cdfc75f80 > __obf_94c254e0a96abc8c {
					break
				}
			}
		}
		__obf_ab781ee4d57d2f62 = append(__obf_ab781ee4d57d2f62, strings.ReplaceAll(strings.ReplaceAll(__obf_75eed1de20c3a609, "\n", ""), ">", ""))
	}

	// Close the communication with the LLM
	__obf_a5e90230dda013ab(__obf_4eca3ea6802751dd, __obf_31a1e60ea885fa8c)

	// Return the LLM responses
	return __obf_ab781ee4d57d2f62, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_102bc8c56ccf3110 *LLM) BufferPromptModel(__obf_fd4f5fc0fc0f75cd string, __obf_6bac99e5bf8a05fc chan<- string) {
	__obf_102bc8c56ccf3110.__obf_1804c9056651edc9()

	__obf_4eca3ea6802751dd, __obf_31a1e60ea885fa8c, __obf_e4b510a2f5dd7a58, __obf_2df63f771f77e5ca := __obf_146b7a3099e1b50c(__obf_102bc8c56ccf3110)

	if __obf_2df63f771f77e5ca != nil {
		fmt.Println("Error creating pipes:", __obf_2df63f771f77e5ca)
		return
	}

	// Start the llama.cpp llm communication process
	__obf_08ca8321e99f54ef := __obf_4eca3ea6802751dd.Start()
	if __obf_08ca8321e99f54ef != nil {
		fmt.Println("Error starting command:", __obf_08ca8321e99f54ef)
		return
	}

	// Create a buffer for the stdout
	__obf_04e69d3727873762 := make([]byte, 1024)

	// Add the instruction block to the input
	__obf_bbce54933b8ea82d := __obf_102bc8c56ccf3110.InstructionBlock + __obf_fd4f5fc0fc0f75cd

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_bbce54933b8ea82d, "\n") {
		__obf_bbce54933b8ea82d += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_31a1e60ea885fa8c, __obf_bbce54933b8ea82d)

	// Create a counter to detect when the response is complete
	__obf_a9f64a3cdfc75f80 := 0
	for {
		__obf_f9ef8d129073080f, __obf_2df63f771f77e5ca := __obf_e4b510a2f5dd7a58.Read(__obf_04e69d3727873762)
		if __obf_2df63f771f77e5ca != nil {
			if __obf_2df63f771f77e5ca != io.EOF {
				fmt.Println("Error reading token:", __obf_2df63f771f77e5ca)
			}
			break
		}

		__obf_2c11f0bf65608422 := string(__obf_04e69d3727873762[:__obf_f9ef8d129073080f])

		if strings.Contains(__obf_2c11f0bf65608422, "\n>") {
			__obf_a9f64a3cdfc75f80 += 1
			if __obf_a9f64a3cdfc75f80 > 1 {
				break
			}
		} else {
			__obf_6bac99e5bf8a05fc <- __obf_2c11f0bf65608422
		}
	}

	// Close the communication with the LLM
	__obf_a5e90230dda013ab(__obf_4eca3ea6802751dd, __obf_31a1e60ea885fa8c)

	// Close the channel to signal end of data transferring
	close(__obf_6bac99e5bf8a05fc)
}
