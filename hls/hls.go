package __obf_08ee9322ff6adb83

// https://github.com/shimberger/gohls
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"syscall"
)

const (
	FFProbePath = "ffprobe"
	FFMpegPath  = "ffmpeg"
)

type Option struct {
	CacheDir            string  `yaml:"cache-dir"`
	SegmentLen          float32 `yaml:"segment-len"`
	WorkersNum          int     `yaml:"workers-num"`
	ExeTimeout          int     `yaml:"exe-timeout"`
	PlaylistContentType string  `yaml:"playlist-content-type"` //  "application/vnd.apple.mpegurl"
	CaptionContentType  string  `yaml:"caption-content-type"`  //  "text/vtt"
}

type HlsServer struct {
	Option                 *Option
	__obf_68bd4c2dc42e505f *Worker
	__obf_1a9d4d2fd79eae45 *Worker
	__obf_1b3c0ab03cc7a77f *Worker
}

func NewHlsServer(__obf_01b7be19c61f6241 *Option) *HlsServer {
	return &HlsServer{__obf_01b7be19c61f6241, NewWorker(__obf_01b7be19c61f6241.WorkersNum, __obf_01b7be19c61f6241.ExeTimeout, __obf_01b7be19c61f6241.CacheDir),
		NewWorker(__obf_01b7be19c61f6241.WorkersNum, __obf_01b7be19c61f6241.ExeTimeout, __obf_01b7be19c61f6241.CacheDir),
		NewWorker(__obf_01b7be19c61f6241.WorkersNum, __obf_01b7be19c61f6241.ExeTimeout, __obf_01b7be19c61f6241.CacheDir),
	}
}

func ExecJsonStdout(__obf_e13c1cb8fa57e03f *exec.Cmd, __obf_53e18a26e170d5aa any) error {
	__obf_891592dc366130ba := new(bytes.Buffer)
	if __obf_dd2d47498586b1c8 := ExecWriteStdout(__obf_e13c1cb8fa57e03f, __obf_891592dc366130ba); __obf_dd2d47498586b1c8 != nil {
		return __obf_dd2d47498586b1c8
	}

	return json.Unmarshal(__obf_891592dc366130ba.Bytes(), __obf_53e18a26e170d5aa)
}

func ExecWriteStdout(__obf_e13c1cb8fa57e03f *exec.Cmd, __obf_63e18a6991b3d10e io.Writer) error {
	__obf_8ddef433f5df5305, __obf_dd2d47498586b1c8 := __obf_e13c1cb8fa57e03f.StdoutPipe()
	if __obf_dd2d47498586b1c8 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_dd2d47498586b1c8)
	}
	defer __obf_8ddef433f5df5305.Close()

	if __obf_dd2d47498586b1c8 = __obf_e13c1cb8fa57e03f.Start(); __obf_dd2d47498586b1c8 != nil {
		return fmt.Errorf("error starting command: %v", __obf_dd2d47498586b1c8)
	}

	if _, __obf_dd2d47498586b1c8 = io.Copy(__obf_63e18a6991b3d10e, __obf_8ddef433f5df5305); __obf_dd2d47498586b1c8 != nil {
		// Ask the process to exit
		_ = __obf_e13c1cb8fa57e03f.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_e13c1cb8fa57e03f.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_dd2d47498586b1c8)
	}
	if __obf_dd2d47498586b1c8 = __obf_e13c1cb8fa57e03f.Wait(); __obf_dd2d47498586b1c8 != nil {
		return fmt.Errorf("command failed %v", __obf_dd2d47498586b1c8)
	}

	return nil
}
