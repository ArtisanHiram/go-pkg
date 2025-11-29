package __obf_409ded1d14232b1b

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_9f6113ff63320280 string, __obf_95f424d2a65a51b4 ...string) (string, string, error) {
	return RunStdin(__obf_9f6113ff63320280, "", __obf_95f424d2a65a51b4...)
}

func RunStdin(__obf_9f6113ff63320280, __obf_cad27ccd7894da45 string, __obf_95f424d2a65a51b4 ...string) (string, string, error) {
	__obf_c9cbfec19db93940,
		// 1. 使用带超时的 Context
		__obf_653132c272be39e5 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_653132c272be39e5()
	__obf_409ded1d14232b1b := exec.CommandContext(__obf_c9cbfec19db93940, __obf_95f424d2a65a51b4[0], __obf_95f424d2a65a51b4[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_cbabe2e422500c0e, __obf_85c2ebfbb0fa2da2 bytes.Buffer
	__obf_409ded1d14232b1b.
		Dir = __obf_9f6113ff63320280
	if __obf_cad27ccd7894da45 != "" {
		__obf_409ded1d14232b1b.
			Stdin = strings.NewReader(__obf_cad27ccd7894da45)
	}
	__obf_409ded1d14232b1b.
		Stdout = &__obf_cbabe2e422500c0e
	__obf_409ded1d14232b1b.
		Stderr = &__obf_85c2ebfbb0fa2da2

	// 3. 启动并等待
	__obf_e90a0d6a86d4ee27 := __obf_409ded1d14232b1b.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_c9cbfec19db93940.Err() == context.DeadlineExceeded {
		return __obf_cbabe2e422500c0e.String(), __obf_85c2ebfbb0fa2da2.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_cbabe2e422500c0e.String()), __obf_85c2ebfbb0fa2da2.String(), __obf_e90a0d6a86d4ee27
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_9f6113ff63320280, __obf_322651b697d15436, __obf_cad27ccd7894da45 string) (string, string, error) {
	return RunStdin(__obf_9f6113ff63320280, __obf_cad27ccd7894da45, "bash", "-c", __obf_322651b697d15436)
}
