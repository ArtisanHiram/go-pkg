package __obf_021b37e3798a15be

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
func (__obf_17561f12413966e0 *LLM) GetLLMProps() {
	fmt.Println("Model Path: ", __obf_17561f12413966e0.Model)
	fmt.Println("Llama.cpp Path: ", __obf_17561f12413966e0.Llamacpp)
	fmt.Println("Indexes of Cuda devices to use: ", __obf_17561f12413966e0.CudaDevices)
	fmt.Println("Size of the prompt context: ", __obf_17561f12413966e0.CtxSize)
	fmt.Println("Temperature: ", __obf_17561f12413966e0.Temp)
	fmt.Println("Top-k sampling: ", __obf_17561f12413966e0.TopK)
	fmt.Println("Penalize repeat sequence of tokens: ", __obf_17561f12413966e0.RepeatPenalty)
	fmt.Println("Number of layers to store in VRAM: ", __obf_17561f12413966e0.Ngl)
	fmt.Println("Max number of tokens for model response: ", __obf_17561f12413966e0.MaxTokens)
	fmt.Println("List of generation-stopping strings: ", __obf_17561f12413966e0.Stop)
}

// llmDefaults sets the default values to LLM struct properties in case they are missing user input.
func (__obf_17561f12413966e0 *LLM) __obf_d2731b2ba8ffcfee() {
	if __obf_17561f12413966e0.Model == "" {
		__obf_17561f12413966e0.Model = "./llama.cpp/models/ggml-vocab.bin"
	}
	if __obf_17561f12413966e0.Llamacpp == "" {
		__obf_17561f12413966e0.Llamacpp = "./llama.cpp"
	}
	if __obf_17561f12413966e0.CudaDevices == nil {
		__obf_17561f12413966e0.CudaDevices = []int{0}
	}
	if __obf_17561f12413966e0.CtxSize == 0 {
		__obf_17561f12413966e0.CtxSize = 2048
	}
	if __obf_17561f12413966e0.Temp == 0 {
		__obf_17561f12413966e0.Temp = 0.2
	}
	if __obf_17561f12413966e0.CpuCores == 0 {
		__obf_17561f12413966e0.CpuCores = 8
	}
	if __obf_17561f12413966e0.TopK == 0 {
		__obf_17561f12413966e0.TopK = 10000
	}
	if __obf_17561f12413966e0.RepeatPenalty == 0 {
		__obf_17561f12413966e0.RepeatPenalty = 1.1
	}
	if __obf_17561f12413966e0.MaxTokens == 0 {
		__obf_17561f12413966e0.MaxTokens = 1000
	}
}

// createPipes method creates the running process of llama.cpp on which the LLM will be running. It also creates the communication pipes for GoLlama.
// It returns the command to run llama.cpp, the communication pipes, and an error (in case something went wrong).
func __obf_f4c11a51dc0aa200(__obf_17561f12413966e0 *LLM) (*exec.Cmd, io.WriteCloser, io.ReadCloser, error) {
	__obf_be46af0d799641ea := __obf_17561f12413966e0.Llamacpp + "/main"

	// By default, the models are to be ran on instruction mode hence the '-ins' flag
	__obf_b2d77b745d3ca3b6 := exec.Command(__obf_be46af0d799641ea, "-m", __obf_17561f12413966e0.Model, "--color", "--ctx_size", fmt.Sprint(__obf_17561f12413966e0.CtxSize), "-n", "-1", "-ins", "-b", "128", "--top_k",
		fmt.Sprint(__obf_17561f12413966e0.TopK), "--temp", fmt.Sprint(__obf_17561f12413966e0.Temp), "--repeat_penalty", fmt.Sprint(__obf_17561f12413966e0.RepeatPenalty), "--n-gpu-layers", fmt.Sprint(__obf_17561f12413966e0.Ngl), "-t", fmt.Sprint(__obf_17561f12413966e0.CpuCores), "-ins")

	// Set the working directory if needed (for access to other directories)
	// cmd.Dir = ""

	// Create a writer for sending data to Python's stdin
	__obf_6d22ea458d57162a, __obf_f1278850744fbe33 := __obf_b2d77b745d3ca3b6.StdinPipe()
	if __obf_f1278850744fbe33 != nil {
		fmt.Println("Error creating stdin pipe:", __obf_f1278850744fbe33)
		return nil, nil, nil, __obf_f1278850744fbe33
	}

	// Create pipes for capturing Python's stdout and stderr
	__obf_ad247b98ac6d0769, __obf_f1278850744fbe33 := __obf_b2d77b745d3ca3b6.StdoutPipe()
	if __obf_f1278850744fbe33 != nil {
		fmt.Println("Error creating stdout pipe:", __obf_f1278850744fbe33)
		return nil, nil, nil, __obf_f1278850744fbe33
	}

	return __obf_b2d77b745d3ca3b6, __obf_6d22ea458d57162a, __obf_ad247b98ac6d0769, nil
}

// closePipes function closes the provided pipes and closes the command process.
// func closePipes(cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser) {
func __obf_7b2a81f8d7cba679(__obf_b2d77b745d3ca3b6 *exec.Cmd, __obf_6d22ea458d57162a io.WriteCloser) {
	// Close the stdin pipe to signal the end of input
	__obf_420afe4b91a7e334 := __obf_6d22ea458d57162a.Close()

	if __obf_420afe4b91a7e334 != nil {
		fmt.Println("Error when closing the command:", __obf_420afe4b91a7e334)
	}

	// Close the communication with the llm
	__obf_b2d77b745d3ca3b6.Process.Kill()
}

// PromptModel method orderly prompts the LLM with the provided prompts in the array, engaging in a sort of conversation.
// It returns an array with the respones of the LLM, each response matching with the index of its prompt.
func (__obf_17561f12413966e0 *LLM) PromptModel(__obf_a8c294e396741bde []string) ([]string, error) {
	__obf_17561f12413966e0.__obf_d2731b2ba8ffcfee()
	__obf_b2d77b745d3ca3b6, __obf_6d22ea458d57162a, __obf_ad247b98ac6d0769, __obf_f1278850744fbe33 := __obf_f4c11a51dc0aa200(__obf_17561f12413966e0)

	if __obf_f1278850744fbe33 != nil {
		fmt.Println("Error creating pipes:", __obf_f1278850744fbe33)
		return []string{"Error creating pipes."}, __obf_f1278850744fbe33
	}

	// Start the llama.cpp llm communication process
	__obf_3f44c01c5da26c66 := __obf_b2d77b745d3ca3b6.Start()
	if __obf_3f44c01c5da26c66 != nil {
		fmt.Println("Error starting command:", __obf_3f44c01c5da26c66)
		return []string{"Error starting command."}, __obf_3f44c01c5da26c66
	}

	// Array for the collection of outputs
	__obf_3e6b09129ddec167 := []string{}

	// Create a buffer for the stdout
	__obf_a27e3a9d177ca09c := make([]byte, 1024)

	// Create a counter for the amount of completed inputs
	__obf_0266c2ea875a88e2 := 0

	// Prompt all the inputs
	for __obf_fb727560fe3ed043, __obf_bdef4bcb84ed8a86 := range __obf_a8c294e396741bde {
		// When a prompt is first sent, it creates a \n> character automatically, so 'i' is incremented by 1 to reflect this
		__obf_fb727560fe3ed043 = __obf_fb727560fe3ed043 + 1

		// Add the instruction block to the input
		__obf_bdef4bcb84ed8a86 = __obf_17561f12413966e0.InstructionBlock + __obf_bdef4bcb84ed8a86

		// Input must contain an EOL for the LLM to correctly interpret the propmt's end
		if !strings.Contains(__obf_bdef4bcb84ed8a86, "\n") {
			__obf_bdef4bcb84ed8a86 += "\n"
		}

		// Prompting the llm
		io.WriteString(__obf_6d22ea458d57162a, __obf_bdef4bcb84ed8a86)

		__obf_b037182233e9f5e5 := ""
		for {
			__obf_3686174cd14a452a, __obf_f1278850744fbe33 := __obf_ad247b98ac6d0769.Read(__obf_a27e3a9d177ca09c)
			if __obf_f1278850744fbe33 != nil {
				if __obf_f1278850744fbe33 != io.EOF {
					fmt.Println("Error reading token:", __obf_f1278850744fbe33)
				}
				break
			}

			__obf_cc280de71da60403 := string(__obf_a27e3a9d177ca09c[:__obf_3686174cd14a452a])
			__obf_b037182233e9f5e5 = __obf_b037182233e9f5e5 + __obf_cc280de71da60403

			if strings.Contains(__obf_cc280de71da60403, "\n>") {
				__obf_0266c2ea875a88e2 += 1
				if __obf_0266c2ea875a88e2 > __obf_fb727560fe3ed043 {
					break
				}
			}
		}
		__obf_3e6b09129ddec167 = append(__obf_3e6b09129ddec167, strings.ReplaceAll(strings.ReplaceAll(__obf_b037182233e9f5e5, "\n", ""), ">", ""))
	}

	// Close the communication with the LLM
	__obf_7b2a81f8d7cba679(__obf_b2d77b745d3ca3b6, __obf_6d22ea458d57162a)

	// Return the LLM responses
	return __obf_3e6b09129ddec167, nil
}

// BufferPromptModel prompts the model expecting the real time output, allowing you to use its response as it's being generated.
// It sends the LLM response tokens as strings to the provided channel.
func (__obf_17561f12413966e0 *LLM) BufferPromptModel(__obf_05b18460fed1efd9 string, __obf_af22660766dc4478 chan<- string) {
	__obf_17561f12413966e0.__obf_d2731b2ba8ffcfee()

	__obf_b2d77b745d3ca3b6, __obf_6d22ea458d57162a, __obf_ad247b98ac6d0769, __obf_f1278850744fbe33 := __obf_f4c11a51dc0aa200(__obf_17561f12413966e0)

	if __obf_f1278850744fbe33 != nil {
		fmt.Println("Error creating pipes:", __obf_f1278850744fbe33)
		return
	}

	// Start the llama.cpp llm communication process
	__obf_3f44c01c5da26c66 := __obf_b2d77b745d3ca3b6.Start()
	if __obf_3f44c01c5da26c66 != nil {
		fmt.Println("Error starting command:", __obf_3f44c01c5da26c66)
		return
	}

	// Create a buffer for the stdout
	__obf_a27e3a9d177ca09c := make([]byte, 1024)

	// Add the instruction block to the input
	__obf_bdef4bcb84ed8a86 := __obf_17561f12413966e0.InstructionBlock + __obf_05b18460fed1efd9

	// Input must contain an EOL for the LLM to correctly interpret the propmt's end
	if !strings.Contains(__obf_bdef4bcb84ed8a86, "\n") {
		__obf_bdef4bcb84ed8a86 += "\n"
	}

	// Prompting the llm
	io.WriteString(__obf_6d22ea458d57162a, __obf_bdef4bcb84ed8a86)

	// Create a counter to detect when the response is complete
	__obf_0266c2ea875a88e2 := 0
	for {
		__obf_3686174cd14a452a, __obf_f1278850744fbe33 := __obf_ad247b98ac6d0769.Read(__obf_a27e3a9d177ca09c)
		if __obf_f1278850744fbe33 != nil {
			if __obf_f1278850744fbe33 != io.EOF {
				fmt.Println("Error reading token:", __obf_f1278850744fbe33)
			}
			break
		}

		__obf_cc280de71da60403 := string(__obf_a27e3a9d177ca09c[:__obf_3686174cd14a452a])

		if strings.Contains(__obf_cc280de71da60403, "\n>") {
			__obf_0266c2ea875a88e2 += 1
			if __obf_0266c2ea875a88e2 > 1 {
				break
			}
		} else {
			__obf_af22660766dc4478 <- __obf_cc280de71da60403
		}
	}

	// Close the communication with the LLM
	__obf_7b2a81f8d7cba679(__obf_b2d77b745d3ca3b6, __obf_6d22ea458d57162a)

	// Close the channel to signal end of data transferring
	close(__obf_af22660766dc4478)
}
