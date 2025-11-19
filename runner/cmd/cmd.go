package __obf_ac60111547657985

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_2a9e4dbbabca8ba9 string, __obf_bfbc291efe5d86da ...string) (string, string, error) {
	return RunStdin(__obf_2a9e4dbbabca8ba9, "", __obf_bfbc291efe5d86da...)
}

func RunStdin(__obf_2a9e4dbbabca8ba9, __obf_995f33626f4461b9 string, __obf_bfbc291efe5d86da ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_98d5610aacf2e190, __obf_da83b8ac034f2945 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_da83b8ac034f2945()

	__obf_ac60111547657985 := exec.CommandContext(__obf_98d5610aacf2e190, __obf_bfbc291efe5d86da[0], __obf_bfbc291efe5d86da[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_e1164511255a9214, __obf_7d7e9df99ce44d33 bytes.Buffer
	__obf_ac60111547657985.Dir = __obf_2a9e4dbbabca8ba9
	if __obf_995f33626f4461b9 != "" {
		__obf_ac60111547657985.Stdin = strings.NewReader(__obf_995f33626f4461b9)
	}
	__obf_ac60111547657985.Stdout = &__obf_e1164511255a9214
	__obf_ac60111547657985.Stderr = &__obf_7d7e9df99ce44d33

	// 3. 启动并等待
	__obf_2bf6186be491457f := __obf_ac60111547657985.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_98d5610aacf2e190.Err() == context.DeadlineExceeded {
		return __obf_e1164511255a9214.String(), __obf_7d7e9df99ce44d33.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_e1164511255a9214.String()), __obf_7d7e9df99ce44d33.String(), __obf_2bf6186be491457f
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_2a9e4dbbabca8ba9, __obf_cdfd6cb0e7583ba2, __obf_995f33626f4461b9 string) (string, string, error) {
	return RunStdin(__obf_2a9e4dbbabca8ba9, __obf_995f33626f4461b9, "bash", "-c", __obf_cdfd6cb0e7583ba2)
}
