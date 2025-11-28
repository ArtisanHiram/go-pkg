package __obf_42bbcad92b7de1a8

import (
	"fmt"
	"io"
)

// var encodeWorker = NewWorker(
// 	2, filepath.Join(".hls", "segments"),
// 	func(i any, w io.Writer) error {
// 		return nil
// 	},
// )

func (__obf_c7be2ffa97bb9914 *HlsServer) GetSegment(__obf_72ab747420e4c9ee string, __obf_de4697b1ba428e69, __obf_3b4db798ffff338d int, __obf_39100c7ef6f93ec8 io.Writer) error {
	__obf_f81998d0b0cfe63c := float32(__obf_de4697b1ba428e69) * __obf_c7be2ffa97bb9914.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_3f936d0e43e30024 := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_f81998d0b0cfe63c),

		// The source file
		"-i", __obf_72ab747420e4c9ee,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_c7be2ffa97bb9914.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_3b4db798ffff338d),

		// x264 video codec
		"-vcodec", "libx264",

		// x264 preset
		"-preset", "veryfast",

		// aac audio codec
		"-c:a", "aac",
		"-b:a", "128k",
		"-ac", "2",

		// TODO
		"-pix_fmt", "yuv420p",

		//"-r", "25", // fixed framerate

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_c7be2ffa97bb9914.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_c7be2ffa97bb9914.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_f81998d0b0cfe63c),

		"pipe:out%03d.ts",
	}

	return __obf_c7be2ffa97bb9914.__obf_97a68cd1a689bbf5.Serve(FFMpegPath, __obf_3f936d0e43e30024, __obf_39100c7ef6f93ec8)
}
