package __obf_2bcd216388b25f35

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_67f6e6f79c320e24 string, __obf_296bd3767c6fc30d ...string) (string, string, error) {
	return RunStdin(__obf_67f6e6f79c320e24, "", __obf_296bd3767c6fc30d...)
}

func RunStdin(__obf_67f6e6f79c320e24, __obf_c87de42e4e563257 string, __obf_296bd3767c6fc30d ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_6e689f14a4789ad0, __obf_7c6b8271b49e76d2 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_7c6b8271b49e76d2()

	__obf_2bcd216388b25f35 := exec.CommandContext(__obf_6e689f14a4789ad0, __obf_296bd3767c6fc30d[0], __obf_296bd3767c6fc30d[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_7f1edc5481d2a712, __obf_27875a680e75a429 bytes.Buffer
	__obf_2bcd216388b25f35.Dir = __obf_67f6e6f79c320e24
	if __obf_c87de42e4e563257 != "" {
		__obf_2bcd216388b25f35.Stdin = strings.NewReader(__obf_c87de42e4e563257)
	}
	__obf_2bcd216388b25f35.Stdout = &__obf_7f1edc5481d2a712
	__obf_2bcd216388b25f35.Stderr = &__obf_27875a680e75a429

	// 3. 启动并等待
	__obf_1707ded7881e6817 := __obf_2bcd216388b25f35.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_6e689f14a4789ad0.Err() == context.DeadlineExceeded {
		return __obf_7f1edc5481d2a712.String(), __obf_27875a680e75a429.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_7f1edc5481d2a712.String()), __obf_27875a680e75a429.String(), __obf_1707ded7881e6817
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_67f6e6f79c320e24, __obf_098ddb33c46292ab, __obf_c87de42e4e563257 string) (string, string, error) {
	return RunStdin(__obf_67f6e6f79c320e24, __obf_c87de42e4e563257, "bash", "-c", __obf_098ddb33c46292ab)
}
