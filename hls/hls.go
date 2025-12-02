package __obf_eb5e805b9fc230e3

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
	__obf_5468dd2a54dba82c *Worker
	__obf_45e21b72f0a6ee88 *Worker
	__obf_1dd73b629d0430c7 *Worker
}

func NewHlsServer(__obf_f0ccfa74ff9ff250 *Option) *HlsServer {
	return &HlsServer{__obf_f0ccfa74ff9ff250, NewWorker(__obf_f0ccfa74ff9ff250.WorkersNum, __obf_f0ccfa74ff9ff250.ExeTimeout, __obf_f0ccfa74ff9ff250.CacheDir),
		NewWorker(__obf_f0ccfa74ff9ff250.WorkersNum, __obf_f0ccfa74ff9ff250.ExeTimeout, __obf_f0ccfa74ff9ff250.CacheDir),
		NewWorker(__obf_f0ccfa74ff9ff250.WorkersNum, __obf_f0ccfa74ff9ff250.ExeTimeout, __obf_f0ccfa74ff9ff250.CacheDir),
	}
}

func ExecJsonStdout(__obf_81bd5123537bea4c *exec.Cmd, __obf_36f9f90fc6663db6 any) error {
	__obf_4872501882288e11 := new(bytes.Buffer)
	if __obf_2d43222a56faebee := ExecWriteStdout(__obf_81bd5123537bea4c, __obf_4872501882288e11); __obf_2d43222a56faebee != nil {
		return __obf_2d43222a56faebee
	}

	return json.Unmarshal(__obf_4872501882288e11.Bytes(), __obf_36f9f90fc6663db6)
}

func ExecWriteStdout(__obf_81bd5123537bea4c *exec.Cmd, __obf_1cab0e86de8129e5 io.Writer) error {
	__obf_55602775667380a3, __obf_2d43222a56faebee := __obf_81bd5123537bea4c.StdoutPipe()
	if __obf_2d43222a56faebee != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_2d43222a56faebee)
	}
	defer __obf_55602775667380a3.Close()

	if __obf_2d43222a56faebee = __obf_81bd5123537bea4c.Start(); __obf_2d43222a56faebee != nil {
		return fmt.Errorf("error starting command: %v", __obf_2d43222a56faebee)
	}

	if _, __obf_2d43222a56faebee = io.Copy(__obf_1cab0e86de8129e5, __obf_55602775667380a3); __obf_2d43222a56faebee != nil {
		// Ask the process to exit
		_ = __obf_81bd5123537bea4c.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_81bd5123537bea4c.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_2d43222a56faebee)
	}
	if __obf_2d43222a56faebee = __obf_81bd5123537bea4c.Wait(); __obf_2d43222a56faebee != nil {
		return fmt.Errorf("command failed %v", __obf_2d43222a56faebee)
	}

	return nil
}
