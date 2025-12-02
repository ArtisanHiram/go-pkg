package __obf_1bd04ed9f422a876

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_5016851075695cca string, __obf_417cbec5e9067df4 ...string) (string, string, error) {
	return RunStdin(__obf_5016851075695cca, "", __obf_417cbec5e9067df4...)
}

func RunStdin(__obf_5016851075695cca, __obf_9e58244a7ab21eb9 string, __obf_417cbec5e9067df4 ...string) (string, string, error) {
	__obf_3e1bad6838d6baed,
		// 1. 使用带超时的 Context
		__obf_1320091081408b73 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_1320091081408b73()
	__obf_1bd04ed9f422a876 := exec.CommandContext(__obf_3e1bad6838d6baed, __obf_417cbec5e9067df4[0], __obf_417cbec5e9067df4[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_c6712f35c652d90f, __obf_3493efd5e47ad38e bytes.Buffer
	__obf_1bd04ed9f422a876.
		Dir = __obf_5016851075695cca
	if __obf_9e58244a7ab21eb9 != "" {
		__obf_1bd04ed9f422a876.
			Stdin = strings.NewReader(__obf_9e58244a7ab21eb9)
	}
	__obf_1bd04ed9f422a876.
		Stdout = &__obf_c6712f35c652d90f
	__obf_1bd04ed9f422a876.
		Stderr = &__obf_3493efd5e47ad38e

	// 3. 启动并等待
	__obf_447a758b235ca73f := __obf_1bd04ed9f422a876.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_3e1bad6838d6baed.Err() == context.DeadlineExceeded {
		return __obf_c6712f35c652d90f.String(), __obf_3493efd5e47ad38e.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_c6712f35c652d90f.String()), __obf_3493efd5e47ad38e.String(), __obf_447a758b235ca73f
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_5016851075695cca, __obf_8a2b18c14a15b110, __obf_9e58244a7ab21eb9 string) (string, string, error) {
	return RunStdin(__obf_5016851075695cca, __obf_9e58244a7ab21eb9, "bash", "-c", __obf_8a2b18c14a15b110)
}
