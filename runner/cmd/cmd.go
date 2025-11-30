package __obf_0e189bcd686d41ba

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_8dcc7aceebd58da3 string, __obf_5a6546a7e005ca23 ...string) (string, string, error) {
	return RunStdin(__obf_8dcc7aceebd58da3, "", __obf_5a6546a7e005ca23...)
}

func RunStdin(__obf_8dcc7aceebd58da3, __obf_0c166972157cf43c string, __obf_5a6546a7e005ca23 ...string) (string, string, error) {
	__obf_48c4f4a05e03240d,
		// 1. 使用带超时的 Context
		__obf_548460cd451c254a := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_548460cd451c254a()
	__obf_0e189bcd686d41ba := exec.CommandContext(__obf_48c4f4a05e03240d, __obf_5a6546a7e005ca23[0], __obf_5a6546a7e005ca23[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_a078ccaa7c9b2230, __obf_5edde6cc93834146 bytes.Buffer
	__obf_0e189bcd686d41ba.
		Dir = __obf_8dcc7aceebd58da3
	if __obf_0c166972157cf43c != "" {
		__obf_0e189bcd686d41ba.
			Stdin = strings.NewReader(__obf_0c166972157cf43c)
	}
	__obf_0e189bcd686d41ba.
		Stdout = &__obf_a078ccaa7c9b2230
	__obf_0e189bcd686d41ba.
		Stderr = &__obf_5edde6cc93834146

	// 3. 启动并等待
	__obf_b123d2025c50bdb9 := __obf_0e189bcd686d41ba.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_48c4f4a05e03240d.Err() == context.DeadlineExceeded {
		return __obf_a078ccaa7c9b2230.String(), __obf_5edde6cc93834146.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_a078ccaa7c9b2230.String()), __obf_5edde6cc93834146.String(), __obf_b123d2025c50bdb9
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_8dcc7aceebd58da3, __obf_fe4f3da616100f19, __obf_0c166972157cf43c string) (string, string, error) {
	return RunStdin(__obf_8dcc7aceebd58da3, __obf_0c166972157cf43c, "bash", "-c", __obf_fe4f3da616100f19)
}
