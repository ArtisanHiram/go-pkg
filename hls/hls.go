package __obf_5028a0a829ddcdab

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
	__obf_3c1630e7a0e9c371 *Worker
	__obf_3c22b11f4aa43eed *Worker
	__obf_afca1d674bf7806a *Worker
}

func NewHlsServer(__obf_20a7bafb4e0c1704 *Option) *HlsServer {
	return &HlsServer{
		__obf_20a7bafb4e0c1704,
		NewWorker(__obf_20a7bafb4e0c1704.WorkersNum, __obf_20a7bafb4e0c1704.ExeTimeout, __obf_20a7bafb4e0c1704.CacheDir),
		NewWorker(__obf_20a7bafb4e0c1704.WorkersNum, __obf_20a7bafb4e0c1704.ExeTimeout, __obf_20a7bafb4e0c1704.CacheDir),
		NewWorker(__obf_20a7bafb4e0c1704.WorkersNum, __obf_20a7bafb4e0c1704.ExeTimeout, __obf_20a7bafb4e0c1704.CacheDir),
	}
}

func ExecJsonStdout(__obf_13daf0dd60267c30 *exec.Cmd, __obf_781de6fadb337177 any) error {

	__obf_a128e98fcbfe8702 := new(bytes.Buffer)
	if __obf_97be5ef7ba0b1a8f := ExecWriteStdout(__obf_13daf0dd60267c30, __obf_a128e98fcbfe8702); __obf_97be5ef7ba0b1a8f != nil {
		return __obf_97be5ef7ba0b1a8f
	}

	return json.Unmarshal(__obf_a128e98fcbfe8702.Bytes(), __obf_781de6fadb337177)
}

func ExecWriteStdout(__obf_13daf0dd60267c30 *exec.Cmd, __obf_f0018e81daaf1fa9 io.Writer) error {

	__obf_60beac99b7d9da01, __obf_97be5ef7ba0b1a8f := __obf_13daf0dd60267c30.StdoutPipe()
	if __obf_97be5ef7ba0b1a8f != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_97be5ef7ba0b1a8f)
	}
	defer __obf_60beac99b7d9da01.Close()

	if __obf_97be5ef7ba0b1a8f = __obf_13daf0dd60267c30.Start(); __obf_97be5ef7ba0b1a8f != nil {
		return fmt.Errorf("error starting command: %v", __obf_97be5ef7ba0b1a8f)
	}

	if _, __obf_97be5ef7ba0b1a8f = io.Copy(__obf_f0018e81daaf1fa9, __obf_60beac99b7d9da01); __obf_97be5ef7ba0b1a8f != nil {
		// Ask the process to exit
		_ = __obf_13daf0dd60267c30.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_13daf0dd60267c30.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_97be5ef7ba0b1a8f)
	}
	if __obf_97be5ef7ba0b1a8f = __obf_13daf0dd60267c30.Wait(); __obf_97be5ef7ba0b1a8f != nil {
		return fmt.Errorf("command failed %v", __obf_97be5ef7ba0b1a8f)
	}

	return nil
}
