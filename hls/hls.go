package __obf_5441fcd9a319cf59

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
	__obf_73e20fe74a76e051 *Worker
	__obf_b8a98ccb6b0d29bb *Worker
	__obf_50b1d5691dcc362d *Worker
}

func NewHlsServer(__obf_521498aa92767488 *Option) *HlsServer {
	return &HlsServer{__obf_521498aa92767488, NewWorker(__obf_521498aa92767488.WorkersNum, __obf_521498aa92767488.ExeTimeout, __obf_521498aa92767488.CacheDir),
		NewWorker(__obf_521498aa92767488.WorkersNum, __obf_521498aa92767488.ExeTimeout, __obf_521498aa92767488.CacheDir),
		NewWorker(__obf_521498aa92767488.WorkersNum, __obf_521498aa92767488.ExeTimeout, __obf_521498aa92767488.CacheDir),
	}
}

func ExecJsonStdout(__obf_85b3eedfe1a05aac *exec.Cmd, __obf_f86135f935731f8f any) error {
	__obf_6408a455935ea061 := new(bytes.Buffer)
	if __obf_e8125dea727cd4c9 := ExecWriteStdout(__obf_85b3eedfe1a05aac, __obf_6408a455935ea061); __obf_e8125dea727cd4c9 != nil {
		return __obf_e8125dea727cd4c9
	}

	return json.Unmarshal(__obf_6408a455935ea061.Bytes(), __obf_f86135f935731f8f)
}

func ExecWriteStdout(__obf_85b3eedfe1a05aac *exec.Cmd, __obf_9e22338370faf798 io.Writer) error {
	__obf_ab73c502130fd530, __obf_e8125dea727cd4c9 := __obf_85b3eedfe1a05aac.StdoutPipe()
	if __obf_e8125dea727cd4c9 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_e8125dea727cd4c9)
	}
	defer __obf_ab73c502130fd530.Close()

	if __obf_e8125dea727cd4c9 = __obf_85b3eedfe1a05aac.Start(); __obf_e8125dea727cd4c9 != nil {
		return fmt.Errorf("error starting command: %v", __obf_e8125dea727cd4c9)
	}

	if _, __obf_e8125dea727cd4c9 = io.Copy(__obf_9e22338370faf798, __obf_ab73c502130fd530); __obf_e8125dea727cd4c9 != nil {
		// Ask the process to exit
		_ = __obf_85b3eedfe1a05aac.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_85b3eedfe1a05aac.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_e8125dea727cd4c9)
	}
	if __obf_e8125dea727cd4c9 = __obf_85b3eedfe1a05aac.Wait(); __obf_e8125dea727cd4c9 != nil {
		return fmt.Errorf("command failed %v", __obf_e8125dea727cd4c9)
	}

	return nil
}
