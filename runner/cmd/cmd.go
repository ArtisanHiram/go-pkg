package __obf_e3a0b9795f1d3cf7

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_a2a31938b58896d9 string, __obf_3c6d61bd39510336 ...string) (string, string, error) {
	return RunStdin(__obf_a2a31938b58896d9, "", __obf_3c6d61bd39510336...)
}

func RunStdin(__obf_a2a31938b58896d9, __obf_2271092036b92bd8 string, __obf_3c6d61bd39510336 ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_965f71f06fa6a1dc, __obf_33e3d66e21a025ca := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_33e3d66e21a025ca()

	__obf_e3a0b9795f1d3cf7 := exec.CommandContext(__obf_965f71f06fa6a1dc, __obf_3c6d61bd39510336[0], __obf_3c6d61bd39510336[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_13ecafdedbb14f77, __obf_64c3ed1216c14745 bytes.Buffer
	__obf_e3a0b9795f1d3cf7.Dir = __obf_a2a31938b58896d9
	if __obf_2271092036b92bd8 != "" {
		__obf_e3a0b9795f1d3cf7.Stdin = strings.NewReader(__obf_2271092036b92bd8)
	}
	__obf_e3a0b9795f1d3cf7.Stdout = &__obf_13ecafdedbb14f77
	__obf_e3a0b9795f1d3cf7.Stderr = &__obf_64c3ed1216c14745

	// 3. 启动并等待
	__obf_53560dd08b3639d2 := __obf_e3a0b9795f1d3cf7.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_965f71f06fa6a1dc.Err() == context.DeadlineExceeded {
		return __obf_13ecafdedbb14f77.String(), __obf_64c3ed1216c14745.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_13ecafdedbb14f77.String()), __obf_64c3ed1216c14745.String(), __obf_53560dd08b3639d2
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_a2a31938b58896d9, __obf_89c62db4bd19d6ce, __obf_2271092036b92bd8 string) (string, string, error) {
	return RunStdin(__obf_a2a31938b58896d9, __obf_2271092036b92bd8, "bash", "-c", __obf_89c62db4bd19d6ce)
}
