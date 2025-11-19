package __obf_dc577bef9ac5c6b8

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

func (__obf_fdd7931bd5a61fac *HlsServer) GetSegment(__obf_cf3312d5b91d1c43 string, __obf_8aa6a4bc911e8e61, __obf_698ddef4b980dea8 int, __obf_39c0ed11835c4488 io.Writer) error {
	__obf_10ba149b4fd3c908 := float32(__obf_8aa6a4bc911e8e61) * __obf_fdd7931bd5a61fac.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_89adfe3713e6bab9 := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_10ba149b4fd3c908),

		// The source file
		"-i", __obf_cf3312d5b91d1c43,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_fdd7931bd5a61fac.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_698ddef4b980dea8),

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

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_fdd7931bd5a61fac.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_fdd7931bd5a61fac.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_10ba149b4fd3c908),

		"pipe:out%03d.ts",
	}

	return __obf_fdd7931bd5a61fac.__obf_4def811ff2ac5d05.Serve(FFMpegPath, __obf_89adfe3713e6bab9, __obf_39c0ed11835c4488)
}
