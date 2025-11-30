package __obf_896db4ffd2028e78

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_9b9af4f1397fafcf string, __obf_753d86ae7153c514 ...string) (string, string, error) {
	return RunStdin(__obf_9b9af4f1397fafcf, "", __obf_753d86ae7153c514...)
}

func RunStdin(__obf_9b9af4f1397fafcf, __obf_64c16dab3fb24fbf string, __obf_753d86ae7153c514 ...string) (string, string, error) {
	__obf_64c698c7faacb0cc,
		// 1. 使用带超时的 Context
		__obf_e89870b6daf12f62 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_e89870b6daf12f62()
	__obf_896db4ffd2028e78 := exec.CommandContext(__obf_64c698c7faacb0cc, __obf_753d86ae7153c514[0], __obf_753d86ae7153c514[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_14ab1ac89ef3988a, __obf_2c770edef7f7f697 bytes.Buffer
	__obf_896db4ffd2028e78.
		Dir = __obf_9b9af4f1397fafcf
	if __obf_64c16dab3fb24fbf != "" {
		__obf_896db4ffd2028e78.
			Stdin = strings.NewReader(__obf_64c16dab3fb24fbf)
	}
	__obf_896db4ffd2028e78.
		Stdout = &__obf_14ab1ac89ef3988a
	__obf_896db4ffd2028e78.
		Stderr = &__obf_2c770edef7f7f697

	// 3. 启动并等待
	__obf_8ab09a3cb35bbc63 := __obf_896db4ffd2028e78.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_64c698c7faacb0cc.Err() == context.DeadlineExceeded {
		return __obf_14ab1ac89ef3988a.String(), __obf_2c770edef7f7f697.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_14ab1ac89ef3988a.String()), __obf_2c770edef7f7f697.String(), __obf_8ab09a3cb35bbc63
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_9b9af4f1397fafcf, __obf_1bb840aa45c79762, __obf_64c16dab3fb24fbf string) (string, string, error) {
	return RunStdin(__obf_9b9af4f1397fafcf, __obf_64c16dab3fb24fbf, "bash", "-c", __obf_1bb840aa45c79762)
}
