package __obf_08ee9322ff6adb83

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

func (__obf_504413fedee71524 *HlsServer) GetSegment(__obf_489cc4007c169b0b string, __obf_e672fbbcc738a24a, __obf_3976094ad76d20a7 int, __obf_63e18a6991b3d10e io.Writer) error {
	__obf_cfcb06a349824846 := float32(__obf_e672fbbcc738a24a) * __obf_504413fedee71524.Option.SegmentLen
	__obf_b8bccd5d97f55f0d := // see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
		[]string{
			// Prevent encoding to run longer than 30 seonds
			"-timelimit", "45",

			// TODO: Some stuff to investigate
			// "-probesize", "524288",
			// "-fpsprobesize", "10",
			// "-analyzeduration", "2147483647",
			// "-hwaccel:0", "vda",

			// The start time
			// important: needs to be before -i to do input seeking
			"-ss", fmt.Sprintf("%.2f", __obf_cfcb06a349824846),

			// The source file
			"-i", __obf_489cc4007c169b0b, // Put all streams to output
			// "-map", "0",

			// The duration
			"-t", fmt.Sprintf("%.2f", __obf_504413fedee71524.Option.SegmentLen),

			// TODO: Find out what it does
			//"-strict", "-2",

			// Synchronize audio
			"-async", "1",

			// 720p
			"-vf", fmt.Sprintf("scale=-2:%d", __obf_3976094ad76d20a7),

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

			"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_504413fedee71524.Option.SegmentLen),

			//"-force_key_frames", "00:00:00.00",
			//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

			//"-f", "mpegts",

			"-f", "ssegment",
			"-segment_time", fmt.Sprintf("%.2f", __obf_504413fedee71524.Option.SegmentLen),
			"-initial_offset", fmt.Sprintf("%.2f", __obf_cfcb06a349824846),

			"pipe:out%03d.ts",
		}

	return __obf_504413fedee71524.__obf_1b3c0ab03cc7a77f.Serve(FFMpegPath, __obf_b8bccd5d97f55f0d, __obf_63e18a6991b3d10e)
}
