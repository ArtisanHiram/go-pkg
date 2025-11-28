package __obf_42bbcad92b7de1a8

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
	__obf_6b0feca1a08b9797 *Worker
	__obf_6b3042483ee82bf2 *Worker
	__obf_97a68cd1a689bbf5 *Worker
}

func NewHlsServer(__obf_e6c9fd00e8935168 *Option) *HlsServer {
	return &HlsServer{
		__obf_e6c9fd00e8935168,
		NewWorker(__obf_e6c9fd00e8935168.WorkersNum, __obf_e6c9fd00e8935168.ExeTimeout, __obf_e6c9fd00e8935168.CacheDir),
		NewWorker(__obf_e6c9fd00e8935168.WorkersNum, __obf_e6c9fd00e8935168.ExeTimeout, __obf_e6c9fd00e8935168.CacheDir),
		NewWorker(__obf_e6c9fd00e8935168.WorkersNum, __obf_e6c9fd00e8935168.ExeTimeout, __obf_e6c9fd00e8935168.CacheDir),
	}
}

func ExecJsonStdout(__obf_a7aad0ff1142655b *exec.Cmd, __obf_f20b361e19413a01 any) error {

	__obf_dbea6d57990eeedb := new(bytes.Buffer)
	if __obf_59b2ec060775896c := ExecWriteStdout(__obf_a7aad0ff1142655b, __obf_dbea6d57990eeedb); __obf_59b2ec060775896c != nil {
		return __obf_59b2ec060775896c
	}

	return json.Unmarshal(__obf_dbea6d57990eeedb.Bytes(), __obf_f20b361e19413a01)
}

func ExecWriteStdout(__obf_a7aad0ff1142655b *exec.Cmd, __obf_39100c7ef6f93ec8 io.Writer) error {

	__obf_c9f3c9e1a944b39e, __obf_59b2ec060775896c := __obf_a7aad0ff1142655b.StdoutPipe()
	if __obf_59b2ec060775896c != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_59b2ec060775896c)
	}
	defer __obf_c9f3c9e1a944b39e.Close()

	if __obf_59b2ec060775896c = __obf_a7aad0ff1142655b.Start(); __obf_59b2ec060775896c != nil {
		return fmt.Errorf("error starting command: %v", __obf_59b2ec060775896c)
	}

	if _, __obf_59b2ec060775896c = io.Copy(__obf_39100c7ef6f93ec8, __obf_c9f3c9e1a944b39e); __obf_59b2ec060775896c != nil {
		// Ask the process to exit
		_ = __obf_a7aad0ff1142655b.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_a7aad0ff1142655b.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_59b2ec060775896c)
	}
	if __obf_59b2ec060775896c = __obf_a7aad0ff1142655b.Wait(); __obf_59b2ec060775896c != nil {
		return fmt.Errorf("command failed %v", __obf_59b2ec060775896c)
	}

	return nil
}
